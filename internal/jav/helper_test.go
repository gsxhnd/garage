package jav

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type nopLogger struct{}

func (nopLogger) Infof(string, ...interface{})  {}
func (nopLogger) Infow(string, ...interface{})  {}
func (nopLogger) Warnw(string, ...interface{})  {}
func (nopLogger) Errorw(string, ...interface{}) {}

func newTestCrawler(t *testing.T, source Source, srv *httptest.Server, cfg Config) Crawler {
	t.Helper()
	cfg.BaseURL = srv.URL
	cfg.Timeout = time.Second
	crawler, err := NewCrawler(source, nopLogger{}, cfg)
	require.NoError(t, err)
	return crawler
}

func writeHTML(w http.ResponseWriter, body string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(body))
}

const javbusMovieHTML = `<!DOCTYPE html>
<html><body>
<div class="container">
  <h3>SSIS-001 Sample Title</h3>
  <div class="screencap">
    <a class="bigImage" href="/cover.jpg"><img src="/cover.jpg"></a>
  </div>
  <div class="info">
    <p><span class="header">識別碼:</span> <span>SSIS-001</span></p>
    <p><span class="header">發行日期:</span> 2021-03-09</p>
    <p><span class="header">長度:</span> 120分鐘</p>
    <p><span class="header">導演:</span> <a href="/director/x">Some Director</a></p>
    <p><span class="header">製作商:</span> <a>S1</a></p>
    <p><span class="header">發行商:</span> <a>S1 Publisher</a></p>
    <p><span class="header">系列:</span> <a>Series A</a></p>
  </div>
  <ul>
    <li><div class="star-name"><a title="Star One">Star One</a></div></li>
    <li><div class="star-name"><a title="Star Two">Star Two</a></div></li>
  </ul>
</div>
<script>
var gid = 42785257471;
var uc = 0;
var img = '/cover.jpg';
</script>
</body></html>`

const javbusMovie2HTML = `<!DOCTYPE html>
<html><body>
<div class="container">
  <h3>SSIS-002 Another Title</h3>
  <div class="screencap"><a class="bigImage" href="/cover.jpg"><img src="/cover.jpg"></a></div>
  <div class="info">
    <p><span class="header">識別碼:</span> <span>SSIS-002</span></p>
    <p><span class="header">發行日期:</span> 2021-04-01</p>
    <p><span class="header">長度:</span> 90分鐘</p>
  </div>
</div>
</body></html>`

const javbusMovie3HTML = `<!DOCTYPE html>
<html><body>
<div class="container">
  <h3>SSIS-003 Third Title</h3>
  <div class="info"><p><span class="header">識別碼:</span> <span>SSIS-003</span></p></div>
</div>
</body></html>`

const javbusMagnetHTML = `
<tr height="35px">
  <td><a href="magnet:?xt=urn:btih:AAA">SSIS-001-HD</a></td>
  <td><a href="magnet:?xt=urn:btih:AAA">4.85GB</a></td>
  <td><a href="magnet:?xt=urn:btih:AAA">2021-03-10</a></td>
  <td><a>高清</a><a>字幕</a></td>
</tr>
<tr height="35px">
  <td><a href="magnet:?xt=urn:btih:BBB">SSIS-001</a></td>
  <td><a>1.2GB</a></td>
  <td><a>2021-03-09</a></td>
</tr>`

const javbusStarPage1HTML = `<!DOCTYPE html>
<html><body>
<div id="waterfall">
  <div><a class="movie-box" href="/SSIS-001"></a></div>
  <div><a class="movie-box" href="/SSIS-002"></a></div>
</div>
<div class="text-center hidden-xs">
  <ul><a id="next" href="/star/abc/2">next</a></ul>
</div>
</body></html>`

const javbusStarPage2HTML = `<!DOCTYPE html>
<html><body>
<div id="waterfall">
  <div><a class="movie-box" href="/SSIS-003"></a></div>
</div>
</body></html>`

const javdbSearchHTML = `<!DOCTYPE html>
<html><body>
<div id="videos">
  <div class="grid columns">
    <div class="grid-item column">
      <a class="box" href="/v/other"><div class="uid">SSIS-888</div></a>
    </div>
    <div class="grid-item column">
      <a class="box" href="/v/Match"><div class="uid">SSIS-001</div></a>
    </div>
  </div>
</div>
</body></html>`

const javdbMovieHTML = `<!DOCTYPE html>
<html><body>
<section>
  <h2 class="title is-4"><strong>SSIS-001</strong> Sample Title</h2>
  <div class="column column-video-cover">
    <a href="/cover.jpg"><img src="/cover.jpg"></a>
  </div>
  <div class="movie-panel-info">
    <div class="panel-block"><strong>番號:</strong><span class="value">SSIS-001</span></div>
    <div class="panel-block"><strong>日期:</strong><span class="value">2021-03-09</span></div>
    <div class="panel-block"><strong>時長:</strong><span class="value">120 分鍾</span></div>
    <div class="panel-block"><strong>導演:</strong><span class="value"><a>Some Director</a></span></div>
    <div class="panel-block"><strong>片商:</strong><span class="value"><a>S1</a></span></div>
    <div class="panel-block"><strong>發行:</strong><span class="value"><a>S1 Publisher</a></span></div>
    <div class="panel-block"><strong>系列:</strong><span class="value"><a>Series A</a></span></div>
    <div class="panel-block"><strong>演員:</strong><span class="value"><a>Star One</a><a>Star Two</a></span></div>
  </div>
  <div id="magnets-content">
    <div class="item">
      <div class="magnet-name">
        <a href="magnet:?xt=urn:btih:AAA">
          <span class="name">SSIS-001-HD</span>
          <span class="meta">4.85GB, 3個文件</span>
          <span class="tag">高清</span>
          <span class="tag">字幕</span>
        </a>
      </div>
    </div>
  </div>
</section>
</body></html>`

const javdbMovie2HTML = `<!DOCTYPE html>
<html><body>
<section>
  <h2 class="title is-4"><strong>SSIS-002</strong> Another</h2>
  <div class="movie-panel-info">
    <div class="panel-block"><strong>番號:</strong><span class="value">SSIS-002</span></div>
  </div>
</section>
</body></html>`

const javdbActorPage1HTML = `<!DOCTYPE html>
<html><body>
<div id="videos">
  <div class="grid-item column"><a class="box" href="/v/Match"><div class="uid">SSIS-001</div></a></div>
</div>
<a class="pagination-next" href="/actors/xyz?page=2">Next</a>
</body></html>`

const javdbActorPage2HTML = `<!DOCTYPE html>
<html><body>
<div id="videos">
  <div class="grid-item column"><a class="box" href="/v/Second"><div class="uid">SSIS-002</div></a></div>
</div>
</body></html>`
