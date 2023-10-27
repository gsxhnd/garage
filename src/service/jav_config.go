package service

import "time"

type JavCrawlConfig struct {
	Proxy          string        `json:"proxy"`
	DestPath       string        `json:"dest_path"`
	DownloadMagent bool          `json:"download_magent"`
	DownloadCover  bool          `json:"download_cover"`
	Code           string        `json:"code"`
	StarCode       string        `json:"star_code"`
	PrefixCode     string        `json:"prefix_code"`
	PrefixMinNo    uint          `json:"prefix_min_no"`
	PrefixMaxNo    uint          `json:"prefix_max_no"`
	PageStartNo    uint          `json:"page_start_no"`
	RandomDelay    time.Duration `json:"random_delay"`
	Parallelism    int           `json:"parallelism"`
}
