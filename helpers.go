package main

import (
	"fmt"
)

// byteCountSI converts a file size in bytes to a human-readable string
// representation with SI units (e.g. kB, MB, GB). The function rounds to
// one decimal place and uses base 1000 for unit conversion.
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(b)/float64(div), "kMGTPE"[exp])
}
