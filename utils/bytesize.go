package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	byteUnit  = 1
	kbUnit    = 1 << 10
	mbUnit    = 1 << 20
	gbUnit    = 1 << 30
	tbUnit    = 1 << 40
	pbUnit    = 1 << 50
	ebUnit    = 1 << 60
	mbDivisor = float64(mbUnit)
)

// Binary units with SI names (1 KB = 1024 B), matching common magnet listings.
var byteSizeUnits = map[string]float64{
	"B": byteUnit, "BYTE": byteUnit, "BYTES": byteUnit,

	"K": kbUnit, "KB": kbUnit, "KIB": kbUnit,
	"KILOBYTE": kbUnit, "KILOBYTES": kbUnit,

	"M": mbUnit, "MB": mbUnit, "MIB": mbUnit,
	"MEGABYTE": mbUnit, "MEGABYTES": mbUnit,

	"G": gbUnit, "GB": gbUnit, "GIB": gbUnit,
	"GIGABYTE": gbUnit, "GIGABYTES": gbUnit,

	"T": tbUnit, "TB": tbUnit, "TIB": tbUnit,
	"TERABYTE": tbUnit, "TERABYTES": tbUnit,

	"P": pbUnit, "PB": pbUnit, "PIB": pbUnit,
	"PETABYTE": pbUnit, "PETABYTES": pbUnit,

	"E": ebUnit, "EB": ebUnit, "EIB": ebUnit,
	"EXABYTE": ebUnit, "EXABYTES": ebUnit,
}

// ParseByteSize parses a human-readable size such as "1.5GB" or "512 MB".
// Units use binary multiples (1 KB = 1024 B).
func ParseByteSize(s string) (uint64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("empty byte size")
	}

	splitAt := -1
	for i, r := range s {
		if !unicode.IsDigit(r) && r != '.' {
			splitAt = i
			break
		}
	}
	if splitAt <= 0 {
		return 0, fmt.Errorf("unrecognized byte size %q", s)
	}

	numStr := strings.TrimSpace(s[:splitAt])
	unitStr := strings.ToUpper(strings.TrimSpace(s[splitAt:]))
	unit, ok := byteSizeUnits[unitStr]
	if !ok {
		return 0, fmt.Errorf("unrecognized size suffix %q", strings.TrimSpace(s[splitAt:]))
	}

	value, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid byte size %q: %w", s, err)
	}
	if value < 0 {
		return 0, fmt.Errorf("negative byte size %q", s)
	}

	return uint64(value * unit), nil
}

// ByteSizeToMB converts bytes to megabytes, rounded to two decimal places.
func ByteSizeToMB(n uint64) float64 {
	formatted := strconv.FormatFloat(float64(n)/mbDivisor, 'f', 2, 64)
	mb, err := strconv.ParseFloat(formatted, 64)
	if err != nil {
		return 0
	}
	return mb
}
