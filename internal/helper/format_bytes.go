package helper

import (
	"fmt"
)

func FormatBytesToMb(bytes float64) string {
	num := bytes / 1024 / 1024
	return fmt.Sprintf("%.2f Mb", num)
}
