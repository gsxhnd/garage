package jav

import (
	"testing"
)

func Test_crawlClient_saveJavInfos(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cc := NewCrawlClient(nil)
			err := cc.saveJavInfos()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
