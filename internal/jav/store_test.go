package jav

import (
	"context"
	"encoding/csv"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCrawlerUnsupportedSource(t *testing.T) {
	_, err := NewCrawler(Source("nope"), nopLogger{}, Config{})
	require.EqualError(t, err, `jav: unsupported source "nope"`)
}

func TestNewCrawlerNilLogger(t *testing.T) {
	_, err := NewCrawler(SourceJavbus, nil, Config{})
	require.EqualError(t, err, "jav: logger is nil")
}

func TestNewCrawlerInvalidProxy(t *testing.T) {
	_, err := NewCrawler(SourceJavbus, nopLogger{}, Config{Proxy: "not-a-url"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid proxy")
}

func TestLocalStoreSave(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("img"))
	}))
	t.Cleanup(srv.Close)

	out := t.TempDir()
	store, err := NewLocalStore(nopLogger{}, Config{
		OutDir:        out,
		DownloadCover: true,
		Timeout:       0, // apply default
	})
	require.NoError(t, err)

	err = store.Save(context.Background(), []Movie{{
		Code:    "SSIS-001",
		Title:   "Sample",
		Cover:   srv.URL + "/cover.jpg",
		PageURL: "http://example.test/SSIS-001",
		Magnets: []Magnet{{Link: "magnet:?xt=urn:btih:AAA"}},
	}})
	require.NoError(t, err)

	csvFiles, err := filepath.Glob(filepath.Join(out, "*-jav_info.csv"))
	require.NoError(t, err)
	require.Len(t, csvFiles, 1)

	f, err := os.Open(csvFiles[0])
	require.NoError(t, err)
	defer f.Close()
	rows, err := csv.NewReader(f).ReadAll()
	require.NoError(t, err)
	require.Len(t, rows, 2)
	assert.Equal(t, "SSIS-001", rows[1][0])
	assert.Equal(t, "Sample", rows[1][1])

	magnetFiles, err := filepath.Glob(filepath.Join(out, "*-jav_magnet.txt"))
	require.NoError(t, err)
	require.Len(t, magnetFiles, 1)
	raw, err := os.ReadFile(magnetFiles[0])
	require.NoError(t, err)
	assert.True(t, strings.Contains(string(raw), "SSIS-001\tmagnet:?xt=urn:btih:AAA"))

	cover, err := os.ReadFile(filepath.Join(out, "cover", "SSIS-001.jpg"))
	require.NoError(t, err)
	assert.Equal(t, []byte("img"), cover)
}

func TestLocalStoreSaveEmpty(t *testing.T) {
	store, err := NewLocalStore(nopLogger{}, Config{OutDir: t.TempDir()})
	require.NoError(t, err)
	require.NoError(t, store.Save(context.Background(), nil))
}
