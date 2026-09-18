package jav

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gsxhnd/garage/utils"
)

var (
	_ Crawler = (*javbusCrawler)(nil)

	javbusGIDRe = regexp.MustCompile(`var\s+gid\s*=\s*(\d+)`)
	javbusUCRe  = regexp.MustCompile(`var\s+uc\s*=\s*(\d+)`)
	javbusImgRe = regexp.MustCompile(`var\s+img\s*=\s*'([^']+)'`)
)

type javbusCrawler struct {
	*engine
}

func (c *javbusCrawler) Source() Source { return SourceJavbus }

func (c *javbusCrawler) Crawl(ctx context.Context, query Query) ([]Movie, error) {
	codes, err := query.resolvedCodes()
	if err != nil {
		return nil, err
	}
	if len(codes) == 0 && len(query.StarIDs) == 0 {
		return nil, errEmptyQuery
	}

	var targets []string
	for _, code := range codes {
		targets = append(targets, c.movieURL(code))
	}
	for _, starID := range query.StarIDs {
		urls, err := c.listStarMovies(ctx, starID)
		if err != nil {
			return nil, err
		}
		targets = append(targets, urls...)
	}
	targets = uniquePreserve(targets)
	c.logger.Infow("javbus crawl start", "targets", len(targets))
	return c.crawlTargets(ctx, targets, c.fetchMovie)
}

func (c *javbusCrawler) movieURL(code string) string {
	return joinBase(c.cfg.BaseURL, strings.TrimSpace(code))
}

func (c *javbusCrawler) listStarMovies(ctx context.Context, starID string) ([]string, error) {
	page := joinBase(c.cfg.BaseURL, "star", strings.TrimSpace(starID))
	visited := make(map[string]struct{})
	var movies []string

	for i := 0; i < maxListPages; i++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		if _, ok := visited[page]; ok {
			break
		}
		visited[page] = struct{}{}

		c.logger.Infow("javbus star page", "url", page)
		body, err := c.http.Get(ctx, page, nil)
		if err != nil {
			return nil, fmt.Errorf("javbus star %s: %w", starID, err)
		}
		doc, err := parseHTML(body)
		if err != nil {
			return nil, err
		}

		for _, a := range htmlquery.Find(doc, "//a[contains(@class,'movie-box')]") {
			href := strings.TrimSpace(htmlquery.SelectAttr(a, "href"))
			if href == "" {
				continue
			}
			movies = append(movies, resolveURL(page, href))
		}

		next := attrAt(doc, "//a[@id='next']", "href")
		if next == "" {
			break
		}
		page = resolveURL(page, next)
	}
	return uniquePreserve(movies), nil
}

func (c *javbusCrawler) fetchMovie(ctx context.Context, pageURL string) (Movie, error) {
	c.logger.Infow("javbus movie", "url", pageURL)
	body, err := c.http.Get(ctx, pageURL, nil)
	if err != nil {
		return Movie{}, err
	}
	movie, err := parseJavbusMovie(body, pageURL)
	if err != nil {
		return Movie{}, err
	}
	if c.cfg.DownloadMagnet {
		magnets, err := c.fetchMagnets(ctx, body, pageURL)
		if err != nil {
			c.logger.Errorw("javbus magnets failed", "url", pageURL, "error", err)
		} else {
			movie.Magnets = magnets
		}
	}
	return movie, nil
}

func (c *javbusCrawler) fetchMagnets(ctx context.Context, page []byte, pageURL string) ([]Magnet, error) {
	gid, uc, img, ok := parseJavbusGID(page)
	if !ok {
		return nil, nil
	}
	ajaxURL, err := javbusMagnetURL(c.cfg.BaseURL, gid, uc, img)
	if err != nil {
		return nil, err
	}
	body, err := c.http.Get(ctx, ajaxURL, map[string]string{
		"Referer": pageURL,
	})
	if err != nil {
		return nil, err
	}
	return parseJavbusMagnets(body)
}

func parseJavbusMovie(body []byte, pageURL string) (Movie, error) {
	doc, err := parseHTML(body)
	if err != nil {
		return Movie{}, err
	}

	var movie Movie
	movie.PageURL = pageURL
	movie.Title = textAt(doc, "//h3")
	cover := attrAt(doc, "//a[contains(@class,'bigImage')]", "href")
	if cover == "" {
		cover = attrAt(doc, "//div[contains(@class,'screencap')]//img", "src")
	}
	movie.Cover = resolveURL(pageURL, cover)

	for _, p := range htmlquery.Find(doc, "//div[contains(@class,'info')]/p") {
		headerNode := htmlquery.FindOne(p, ".//span[contains(@class,'header')]")
		header := stripHeader(nodeText(headerNode))
		value := fieldValue(p, headerNode)

		switch {
		case strings.Contains(header, "識別碼"), strings.Contains(header, "识别码"):
			spans := htmlquery.Find(p, "./span")
			if len(spans) >= 2 {
				movie.Code = nodeText(spans[len(spans)-1])
			} else {
				movie.Code = value
			}
		case strings.Contains(header, "發行日期"), strings.Contains(header, "发行日期"):
			movie.PublishDate = value
		case strings.Contains(header, "長度"), strings.Contains(header, "长度"):
			movie.Length = normalizeLength(value)
		case strings.Contains(header, "導演"), strings.Contains(header, "导演"):
			movie.Director = value
		case strings.Contains(header, "製作商"), strings.Contains(header, "制作商"):
			movie.ProduceCompany = value
		case strings.Contains(header, "發行商"), strings.Contains(header, "发行商"):
			movie.PublishCompany = value
		case strings.Contains(header, "系列"):
			movie.Series = value
		}
	}

	movie.Stars = joinStarNames(htmlquery.Find(doc, "//div[contains(@class,'star-name')]/a"))
	if movie.Code == "" {
		return Movie{}, errNotFound
	}
	return movie, nil
}

func parseJavbusGID(page []byte) (gid, uc, img string, ok bool) {
	s := string(page)
	if m := javbusGIDRe.FindStringSubmatch(s); len(m) == 2 {
		gid = m[1]
	}
	if m := javbusUCRe.FindStringSubmatch(s); len(m) == 2 {
		uc = m[1]
	}
	if m := javbusImgRe.FindStringSubmatch(s); len(m) == 2 {
		img = m[1]
	}
	return gid, uc, img, gid != "" && uc != ""
}

func javbusMagnetURL(base, gid, uc, img string) (string, error) {
	u, err := url.Parse(trimSlash(base) + "/ajax/uncledatoolsbyajax.php")
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("gid", gid)
	q.Set("lang", "zh")
	q.Set("img", img)
	q.Set("uc", uc)
	q.Set("floor", strconv.Itoa(rand.IntN(1000)+1))
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func parseJavbusMagnets(body []byte) ([]Magnet, error) {
	wrapped := "<table><tbody>" + string(body) + "</tbody></table>"
	doc, err := parseHTML([]byte(wrapped))
	if err != nil {
		return nil, err
	}

	var magnets []Magnet
	for _, tr := range htmlquery.Find(doc, "//tr") {
		anchors := htmlquery.Find(tr, ".//td/a")
		if len(anchors) == 0 {
			continue
		}
		m := Magnet{
			Link: strings.TrimSpace(htmlquery.SelectAttr(anchors[0], "href")),
			Name: strings.TrimSpace(nodeText(anchors[0])),
		}
		if m.Link == "" || !strings.HasPrefix(m.Link, "magnet:") {
			continue
		}
		for i, a := range anchors {
			if i == 0 {
				continue
			}
			text := strings.TrimSpace(nodeText(a))
			text = strings.ReplaceAll(text, "\n", "")
			text = strings.ReplaceAll(text, "\t", "")
			text = strings.TrimSpace(text)
			switch {
			case text == "高清" || strings.EqualFold(text, "HD"):
				m.HD = true
			case text == "字幕" || strings.EqualFold(text, "SUB"):
				m.Subtitle = true
			case looksLikeDate(text):
				continue
			default:
				if size, err := utils.ParseByteSize(firstMagnetSize(text)); err == nil {
					m.Size = utils.ByteSizeToMB(size)
				}
			}
		}
		magnets = append(magnets, m)
	}
	return magnets, nil
}
