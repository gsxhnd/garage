package jav

import (
	"bytes"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"
	"unicode"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func parseHTML(body []byte) (*html.Node, error) {
	if len(body) == 0 {
		return nil, fmt.Errorf("jav: empty html")
	}
	return htmlquery.Parse(bytes.NewReader(body))
}

func nodeText(n *html.Node) string {
	if n == nil {
		return ""
	}
	return strings.TrimSpace(htmlquery.InnerText(n))
}

func textAt(n *html.Node, xpath string) string {
	if n == nil {
		return ""
	}
	return nodeText(htmlquery.FindOne(n, xpath))
}

func attrAt(n *html.Node, xpath, key string) string {
	if n == nil {
		return ""
	}
	found := htmlquery.FindOne(n, xpath)
	if found == nil {
		return ""
	}
	return strings.TrimSpace(htmlquery.SelectAttr(found, key))
}

func resolveURL(base, ref string) string {
	ref = strings.TrimSpace(ref)
	if ref == "" {
		return ""
	}
	r, err := url.Parse(ref)
	if err != nil {
		return ref
	}
	if r.IsAbs() {
		return r.String()
	}
	b, err := url.Parse(base)
	if err != nil {
		return ref
	}
	return b.ResolveReference(r).String()
}

func joinBase(base string, elem ...string) string {
	u, err := url.Parse(trimSlash(base) + "/")
	if err != nil {
		return strings.Trim(strings.Join(append([]string{base}, elem...), "/"), "/")
	}
	u.Path = path.Join(append([]string{u.Path}, elem...)...)
	u.RawQuery = ""
	u.Fragment = ""
	return u.String()
}

func normalizeCode(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)
	s = strings.ReplaceAll(s, "_", "-")
	return s
}

func normalizeLength(s string) string {
	s = strings.TrimSpace(s)
	for _, suf := range []string{"分鐘", "分钟", "分鍾", "min", "Min", "MIN"} {
		s = strings.TrimSpace(strings.TrimSuffix(s, suf))
	}
	return s
}

func stripHeader(header string) string {
	header = strings.TrimSpace(header)
	header = strings.TrimRightFunc(header, func(r rune) bool {
		return r == ':' || r == '：' || unicode.IsSpace(r)
	})
	return header
}

func fieldValue(p *html.Node, header *html.Node) string {
	if p == nil {
		return ""
	}
	links := htmlquery.Find(p, ".//a")
	if len(links) > 0 {
		parts := make([]string, 0, len(links))
		for _, a := range links {
			if t := nodeText(a); t != "" {
				parts = append(parts, t)
			}
		}
		if len(parts) > 0 {
			return strings.Join(parts, ";")
		}
	}
	full := nodeText(p)
	if header == nil {
		return full
	}
	return strings.TrimSpace(strings.TrimPrefix(full, nodeText(header)))
}

func joinStarNames(nodes []*html.Node) string {
	parts := make([]string, 0, len(nodes))
	for _, n := range nodes {
		name := strings.TrimSpace(htmlquery.SelectAttr(n, "title"))
		if name == "" {
			name = nodeText(n)
		}
		if name != "" {
			parts = append(parts, name)
		}
	}
	return strings.Join(parts, ";")
}

func firstMagnetSize(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	if i := strings.IndexAny(raw, ",，"); i >= 0 {
		raw = raw[:i]
	}
	return strings.TrimSpace(raw)
}

func looksLikeDate(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 10 {
		return false
	}
	_, err := time.Parse("2006-01-02", s[:10])
	return err == nil
}
