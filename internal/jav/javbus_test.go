package jav

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newJavbusServer(t *testing.T) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/SSIS-001", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javbusMovieHTML)
	})
	mux.HandleFunc("/SSIS-002", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javbusMovie2HTML)
	})
	mux.HandleFunc("/SSIS-003", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javbusMovie3HTML)
	})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	mux.HandleFunc("/ajax/uncledatoolsbyajax.php", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Referer") == "" {
			http.Error(w, "missing referer", http.StatusBadRequest)
			return
		}
		writeHTML(w, javbusMagnetHTML)
	})
	mux.HandleFunc("/star/abc/2", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javbusStarPage2HTML)
	})
	mux.HandleFunc("/star/abc", func(w http.ResponseWriter, r *http.Request) {
		writeHTML(w, javbusStarPage1HTML)
	})
	mux.HandleFunc("/cover.jpg", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/jpeg")
		_, _ = w.Write([]byte("cover"))
	})
	return httptest.NewServer(mux)
}

func TestJavbusCrawlByCode(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{DownloadMagnet: true})
	movies, err := crawler.Crawl(context.Background(), Query{Codes: []string{"SSIS-001"}})
	require.NoError(t, err)
	require.Len(t, movies, 1)

	m := movies[0]
	assert.Equal(t, SourceJavbus, crawler.Source())
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
	assert.Equal(t, srv.URL+"/SSIS-001", m.PageURL)
	require.Len(t, m.Magnets, 2)
	assert.Equal(t, "magnet:?xt=urn:btih:AAA", m.Magnets[0].Link)
	assert.True(t, m.Magnets[0].HD)
	assert.True(t, m.Magnets[0].Subtitle)
	assert.InDelta(t, 4966.4, m.Magnets[0].Size, 0.1)
}

func TestJavbusCrawlSkipsMissingCode(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{})
	movies, err := crawler.Crawl(context.Background(), Query{Codes: []string{"missing", "SSIS-002"}})
	require.NoError(t, err)
	require.Len(t, movies, 1)
	assert.Equal(t, "SSIS-002", movies[0].Code)
}

func TestJavbusCrawlByPrefix(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{})
	movies, err := crawler.Crawl(context.Background(), Query{
		Prefixes:   []string{"SSIS-00"},
		PrefixMin:  1,
		PrefixMax:  2,
		PrefixZero: 1,
	})
	require.NoError(t, err)
	require.Len(t, movies, 2)
	assert.Equal(t, "SSIS-001", movies[0].Code)
	assert.Equal(t, "SSIS-002", movies[1].Code)
}

func TestJavbusCrawlByStar(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{Parallelism: 2})
	movies, err := crawler.Crawl(context.Background(), Query{StarIDs: []string{"abc"}})
	require.NoError(t, err)
	require.Len(t, movies, 3)
	assert.ElementsMatch(t, []string{"SSIS-001", "SSIS-002", "SSIS-003"}, []string{
		movies[0].Code, movies[1].Code, movies[2].Code,
	})
}

func TestJavbusCrawlByVideosDir(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	dir := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(dir, "SSIS-001.mp4"), []byte("x"), 0o644))

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{})
	movies, err := crawler.Crawl(context.Background(), Query{VideosDir: dir})
	require.NoError(t, err)
	require.Len(t, movies, 1)
	assert.Equal(t, "SSIS-001", movies[0].Code)
}

func TestJavbusEmptyQuery(t *testing.T) {
	srv := newJavbusServer(t)
	t.Cleanup(srv.Close)

	crawler := newTestCrawler(t, SourceJavbus, srv, Config{})
	_, err := crawler.Crawl(context.Background(), Query{})
	require.ErrorIs(t, err, errEmptyQuery)
}
