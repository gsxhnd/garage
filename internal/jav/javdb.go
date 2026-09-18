package jav

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gsxhnd/garage/utils"
)

var _ Crawler = (*javdbCrawler)(nil)

type javdbCrawler struct {
	*engine
}

func (c *javdbCrawler) Source() Source { return SourceJavdb }

func (c *javdbCrawler) Crawl(ctx context.Context, query Query) ([]Movie, error) {
	codes, err := query.resolvedCodes()
	if err != nil {
		return nil, err
	}
	if len(codes) == 0 && len(query.StarIDs) == 0 {
		return nil, errEmptyQuery
	}

	var targets []string
	var searchFail int
	for _, code := range codes {
		pageURL, err := c.searchMovie(ctx, code)
		if err != nil {
			if errors.Is(err, errNotFound) {
				c.logger.Warnw("javdb code not found", "code", code)
				continue
			}
			searchFail++
			c.logger.Errorw("javdb search failed", "code", code, "error", err)
			continue
		}
		targets = append(targets, pageURL)
	}
	for _, starID := range query.StarIDs {
		urls, err := c.listActorMovies(ctx, starID)
		if err != nil {
			return nil, err
		}
		targets = append(targets, urls...)
	}
	targets = uniquePreserve(targets)
	if len(targets) == 0 && searchFail > 0 && len(query.StarIDs) == 0 {
		return nil, fmt.Errorf("javdb: failed to resolve %d codes", searchFail)
	}
	c.logger.Infow("javdb crawl start", "targets", len(targets))
	return c.crawlTargets(ctx, targets, c.fetchMovie)
}

func (c *javdbCrawler) searchMovie(ctx context.Context, code string) (string, error) {
	searchURL, err := javdbSearchURL(c.cfg.BaseURL, code)
	if err != nil {
		return "", err
	}
	c.logger.Infow("javdb search", "url", searchURL, "code", code)
	body, err := c.http.Get(ctx, searchURL, nil)
	if err != nil {
		return "", err
	}
	pageURL, err := parseJavdbSearch(body, c.cfg.BaseURL, code)
	if err != nil {
		return "", err
	}
	return pageURL, nil
}

func (c *javdbCrawler) listActorMovies(ctx context.Context, actorID string) ([]string, error) {
	page := joinBase(c.cfg.BaseURL, "actors", strings.TrimSpace(actorID))
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

		c.logger.Infow("javdb actor page", "url", page)
		body, err := c.http.Get(ctx, page, nil)
		if err != nil {
			return nil, fmt.Errorf("javdb actor %s: %w", actorID, err)
		}
		doc, err := parseHTML(body)
		if err != nil {
			return nil, err
		}

		for _, a := range htmlquery.Find(doc, "//div[@id='videos']//a[contains(@class,'box')]") {
			href := strings.TrimSpace(htmlquery.SelectAttr(a, "href"))
			if href == "" {
				continue
			}
			movies = append(movies, resolveURL(page, href))
		}

		next := attrAt(doc, "//a[contains(@class,'pagination-next')]", "href")
		if next == "" || next == "#" {
			break
		}
		page = resolveURL(page, next)
	}
	return uniquePreserve(movies), nil
}

func (c *javdbCrawler) fetchMovie(ctx context.Context, pageURL string) (Movie, error) {
	c.logger.Infow("javdb movie", "url", pageURL)
	body, err := c.http.Get(ctx, pageURL, nil)
	if err != nil {
		return Movie{}, err
	}
	movie, err := parseJavdbMovie(body, pageURL)
	if err != nil {
		return Movie{}, err
	}
	if c.cfg.DownloadMagnet {
		movie.Magnets = parseJavdbMagnets(body)
	}
	return movie, nil
}

func javdbSearchURL(base, code string) (string, error) {
	u, err := url.Parse(trimSlash(base) + "/search")
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("q", strings.TrimSpace(code))
	q.Set("f", "all")
	u.RawQuery = q.Encode()
	return u.String(), nil
}

func parseJavdbSearch(body []byte, base, code string) (string, error) {
	doc, err := parseHTML(body)
	if err != nil {
		return "", err
	}
	want := normalizeCode(code)
	for _, a := range htmlquery.Find(doc, "//div[@id='videos']//a") {
		uid := normalizeCode(textAt(a, ".//div[contains(@class,'uid')]"))
		if uid == "" || uid != want {
			continue
		}
		href := strings.TrimSpace(htmlquery.SelectAttr(a, "href"))
		if href == "" {
			continue
		}
		return resolveURL(base+"/", href), nil
	}
	return "", fmt.Errorf("javdb: code %s not found: %w", code, errNotFound)
}

func parseJavdbMovie(body []byte, pageURL string) (Movie, error) {
	doc, err := parseHTML(body)
	if err != nil {
		return Movie{}, err
	}

	var movie Movie
	movie.PageURL = pageURL
	movie.Title = textAt(doc, "//h2[contains(@class,'title')]")
	cover := attrAt(doc, "//div[contains(@class,'column-video-cover')]//img", "src")
	if cover == "" {
		cover = attrAt(doc, "//img[contains(@class,'video-cover')]", "src")
	}
	movie.Cover = resolveURL(pageURL, cover)

	for _, block := range htmlquery.Find(doc, "//div[contains(@class,'movie-panel-info')]//div[contains(@class,'panel-block')]") {
		header := stripHeader(textAt(block, ".//strong"))
		value := textAt(block, ".//span[contains(@class,'value')]")
		if value == "" {
			value = fieldValue(block, htmlquery.FindOne(block, ".//strong"))
		}
		switch {
		case strings.Contains(header, "番號"), strings.EqualFold(header, "ID"):
			movie.Code = strings.TrimSpace(value)
		case strings.Contains(header, "日期"), strings.Contains(header, "時間"), strings.Contains(header, "时间"), strings.EqualFold(header, "Released Date"):
			movie.PublishDate = value
		case strings.Contains(header, "時長"), strings.Contains(header, "时长"), strings.EqualFold(header, "Duration"):
			movie.Length = normalizeLength(value)
		case strings.Contains(header, "導演"), strings.Contains(header, "导演"), strings.EqualFold(header, "Director"):
			movie.Director = value
		case strings.Contains(header, "片商"), strings.EqualFold(header, "Maker"):
			movie.ProduceCompany = value
		case strings.Contains(header, "發行"), strings.Contains(header, "发行"), strings.EqualFold(header, "Publisher"):
			movie.PublishCompany = value
		case strings.Contains(header, "系列"), strings.EqualFold(header, "Series"):
			movie.Series = value
		case strings.Contains(header, "演員"), strings.Contains(header, "演员"), strings.Contains(header, "Actor"):
			movie.Stars = joinStarNames(htmlquery.Find(block, ".//a"))
			if movie.Stars == "" {
				movie.Stars = value
			}
		}
	}

	if movie.Code == "" {
		strong := textAt(doc, "//h2[contains(@class,'title')]//strong")
		movie.Code = strings.TrimSpace(strong)
	}
	if movie.Code == "" {
		return Movie{}, errNotFound
	}
	return movie, nil
}

func parseJavdbMagnets(body []byte) []Magnet {
	doc, err := parseHTML(body)
	if err != nil {
		return nil
	}

	var magnets []Magnet
	for _, item := range htmlquery.Find(doc, "//div[@id='magnets-content']//div[contains(@class,'item')]") {
		a := htmlquery.FindOne(item, ".//a[starts-with(@href,'magnet:')]")
		if a == nil {
			a = htmlquery.FindOne(item, ".//div[contains(@class,'magnet-name')]//a")
		}
		if a == nil {
			continue
		}
		link := strings.TrimSpace(htmlquery.SelectAttr(a, "href"))
		if link == "" || !strings.HasPrefix(link, "magnet:") {
			continue
		}
		m := Magnet{
			Link: link,
			Name: textAt(a, ".//span[contains(@class,'name')]"),
		}
		if m.Name == "" {
			m.Name = strings.TrimSpace(nodeText(a))
		}
		if sizeRaw := firstMagnetSize(textAt(a, ".//span[contains(@class,'meta')]")); sizeRaw != "" {
			if size, err := utils.ParseByteSize(sizeRaw); err == nil {
				m.Size = utils.ByteSizeToMB(size)
			}
		}
		for _, tag := range htmlquery.Find(item, ".//span[contains(@class,'tag')]") {
			text := nodeText(tag)
			switch {
			case strings.Contains(text, "高清") || strings.EqualFold(text, "HD"):
				m.HD = true
			case strings.Contains(text, "字幕") || strings.Contains(strings.ToLower(text), "sub"):
				m.Subtitle = true
			}
		}
		magnets = append(magnets, m)
	}
	return magnets
}
