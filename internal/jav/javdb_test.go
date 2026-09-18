package jav

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newJavdbServer(t *testing.T) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("q") == "MISSING-001" {
			writeHTML(w, `<div id="videos"></div>`)
			return
		}
		writeHTML(w, javdbSearchHTML)
	})
	mux.HandleFunc("/v/Match", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javdbMovieHTML)
	})
	mux.HandleFunc("/v/Second", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javdbMovie2HTML)
	})
	mux.HandleFunc("/actors/xyz", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("page") == "2" {
			writeHTML(w, javdbActorPage2HTML)
			return
		}
		writeHTML(w, javdbActorPage1HTML)
	})
	mux.HandleFunc("/cover.jpg", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("cover"))
	})
	return httptest.NewServer(mux)
}

func TestJavdbCrawlByCode(t *testing.T) {
	srv := newJavdbServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavdb, srv, Config{DownloadMagnet: true})
	movies, err := crawler.Crawl(context.Background(), Query{Codes: []string{"ssis-001"}})
	require.NoError(t, err)
	require.Len(t, movies, 1)

	m := movies[0]
	assert.Equal(t, SourceJavdb, crawler.Source())
	assert.Equal(t, "SSIS-001", m.Code)
	assert.Equal(t, "SSIS-001 Sample Title", m.Title)
	assert.Equal(t, srv.URL+"/cover.jpg", m.Cover)
	assert.Equal(t, "2021-03-09", m.PublishDate)
	assert.Equal(t, "120", m.Length)
	assert.Equal(t, "Some Director", m.Director)
	assert.Equal(t, "S1", m.ProduceCompany)
	assert.Equal(t, "S1 Publisher", m.PublishCompany)
	assert.Equal(t, "Series A", m.Series)
	assert.Equal(t, "Star One;Star Two", m.Stars)
	require.Len(t, m.Magnets, 1)
	assert.Equal(t, "magnet:?xt=urn:btih:AAA", m.Magnets[0].Link)
	assert.Equal(t, "SSIS-001-HD", m.Magnets[0].Name)
	assert.True(t, m.Magnets[0].HD)
	assert.True(t, m.Magnets[0].Subtitle)
	assert.InDelta(t, 4966.4, m.Magnets[0].Size, 0.1)
}

func TestJavdbCrawlSkipsUnknownCode(t *testing.T) {
	srv := newJavdbServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavdb, srv, Config{})
	movies, err := crawler.Crawl(context.Background(), Query{Codes: []string{"MISSING-001"}})
	require.NoError(t, err)
	assert.Empty(t, movies)
}

func TestJavdbCrawlByStar(t *testing.T) {
	srv := newJavdbServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavdb, srv, Config{})
	movies, err := crawler.Crawl(context.Background(), Query{StarIDs: []string{"xyz"}})
	require.NoError(t, err)
	require.Len(t, movies, 2)
	assert.ElementsMatch(t, []string{"SSIS-001", "SSIS-002"}, []string{
		movies[0].Code, movies[1].Code,
	})
}

func TestJavdbEmptyQuery(t *testing.T) {
	srv := newJavdbServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavdb, srv, Config{})
	_, err := crawler.Crawl(context.Background(), Query{})
	require.ErrorIs(t, err, errEmptyQuery)
}
