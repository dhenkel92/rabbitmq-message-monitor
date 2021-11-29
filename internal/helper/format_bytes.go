package helper

import (
	"fmt"
)

func FormatBytesToKb(bytes float64) string {
	num := bytes / 1024.0
	return fmt.Sprintf("%.2f Kb", num)
}
