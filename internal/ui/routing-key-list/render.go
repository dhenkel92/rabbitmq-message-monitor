package uiRKList

import (
	"fmt"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
)

func (list *RoutingKeyList) renderList() {
	result := make([]string, 0)

	rtL, countL, sizeL := calculateCellWidths(list.list.Inner.Max.X)

	result = append(result, list.renderHeader())
	for _, entry := range list.data {
		result = append(result, fmt.Sprintf("%-*s %-*d %-*s", rtL, entry.RoutingKey, countL, entry.Count, sizeL, helper.FormatBytesToMb(entry.TotalBytes)))
	}

	list.list.Rows = result
}

func (list *RoutingKeyList) renderHeader() string {
	rtL, countL, sizeL := calculateCellWidths(list.list.Inner.Max.X)

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

	return fmt.Sprintf("[%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white)", rtL, rkH, countL, countH, sizeL, sizeH)
}

func calculateCellWidths(rowLenght int) (int, int, int) {
	routingKey := rowLenght / 100 * 60
	count := rowLenght / 100 * 20
	totalSize := rowLenght / 100 * 20
	return routingKey, count, totalSize
}
