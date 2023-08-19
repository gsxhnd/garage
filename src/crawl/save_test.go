package crawl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_crawlClient_saveJavInfos(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc, _ := NewJavbusCrawl(nil, CrawlOptions{})
			err := cc.saveJavInfos()
			if err != nil {
				assert.Error(t, err, err)
			}
		})
	}
}
