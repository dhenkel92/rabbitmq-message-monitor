package uiRKList

import (
	"fmt"
	"math"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
)

func (list *RoutingKeyList) renderList() {
	result := make([]string, 0)

	width := calculateCellWidths(list.list.Inner.Max.X)

	result = append(result, list.renderHeader(width))
	for _, entry := range list.data {
		result = append(result, fmt.Sprintf(
			"%-*s %-*d %-*s %-*s %-*s %-*s",
			width.routingKey, entry.RoutingKey,
			width.count, entry.Count,
			width.avg, helper.FormatBytesToKb(entry.Avg),
			width.ninetyFife, helper.FormatBytesToKb(entry.NinetyFife),
			width.ninetyNine, helper.FormatBytesToKb(entry.NinetyNine),
			width.totalSize, helper.FormatBytesToKb(entry.TotalBytes),
		))
	}

	list.list.Rows = result
}

func (list *RoutingKeyList) renderHeader(width ListColumnWidhts) string {
	rkH := "Routing Key"
	countH := "Count"
	sizeH := "Total Size"
	switch list.sorting {
	case RK_SORTING_NAME_ASC:
		rkH += " ▲"
	case RK_SORTING_NAME_DESC:
		rkH += " ▼"
	case RK_SORTING_COUNT_ASC:
		countH += " ▲"
	case RK_SORTING_COUNT_DESC:
		countH += " ▼"
	case RK_SORTING_TOTAL_SIZE_ASC:
		sizeH += " ▲"
	case RK_SORTING_TOTAL_SIZE_DESC:
		sizeH += " ▼"
	}

	return fmt.Sprintf(
		"[%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white)",
		width.routingKey, rkH,
		width.count, countH,
		width.avg, "avg",
		width.ninetyFife, "95%",
		width.ninetyNine, "99%",
		width.totalSize, sizeH,
	)
}

func calculateCellWidths(rowLenght int) ListColumnWidhts {
	return ListColumnWidhts{
		routingKey: int(math.Round(float64(rowLenght) / 100.0 * 25.0)),
		count:      int(math.Round(float64(rowLenght) / 100.0 * 14.0)),
		avg:        int(math.Round(float64(rowLenght) / 100.0 * 14.0)),
		ninetyFife: int(math.Round(float64(rowLenght) / 100.0 * 14.0)),
		ninetyNine: int(math.Round(float64(rowLenght) / 100.0 * 14.0)),
		totalSize:  int(math.Round(float64(rowLenght) / 100.0 * 14.0)),
	}
}
