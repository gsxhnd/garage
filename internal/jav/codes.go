package jav

import (
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
)

var videoExts = map[string]struct{}{
	".avi":  {},
	".mp4":  {},
	".mkv":  {},
	".wmv":  {},
	".iso":  {},
	".ts":   {},
	".m2ts": {},
	".mov":  {},
	".flv":  {},
	".webm": {},
}

// ExpandPrefix builds zero-padded codes from prefix+min..max (inclusive).
func ExpandPrefix(prefix string, min, max, zero uint64) []string {
	prefix = strings.TrimSpace(prefix)
	if prefix == "" || max < min {
		return nil
	}
	width := int(zero)
	codes := make([]string, 0, int(max-min)+1)
	for i := min; i <= max; i++ {
		n := strconv.FormatUint(i, 10)
		if width > len(n) {
			n = strings.Repeat("0", width-len(n)) + n
		}
		codes = append(codes, prefix+n)
	}
	return codes
}

// CodesFromVideosDir walks dir and returns video filenames without extension.
func CodesFromVideosDir(dir string) ([]string, error) {
	var codes []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if _, ok := videoExts[ext]; !ok {
			return nil
		}
		code := strings.TrimSpace(strings.TrimSuffix(d.Name(), filepath.Ext(d.Name())))
		if code != "" {
			codes = append(codes, code)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return uniquePreserve(codes), nil
}
