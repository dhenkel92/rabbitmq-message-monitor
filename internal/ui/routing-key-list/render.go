package uiRKList

import (
	"fmt"
	"math"

	"github.com/dustin/go-humanize"
)

func (list *RoutingKeyList) Render() {
	filtered := list.data
	if list.nameFilterExpression != "" {
		filtered = filterData(list.nameFilterExpression, filtered)
	}
	filtered = sortData(list.sorting, filtered)
	list.renderList(filtered)
}

func (list *RoutingKeyList) renderList(data []*RoutingKeyData) {
	result := make([]string, 0)

	width := calculateCellWidths(list.list.Inner.Max.X)

	result = append(result, list.renderHeader(width))
	for _, entry := range data {
		result = append(result, fmt.Sprintf(
			"%-*s %-*d %-*s %-*s %-*s %-*s",
			width.routingKey, entry.RoutingKey,
			width.count, entry.Count,
			width.avg, humanize.Bytes(uint64(entry.Avg)),
			width.ninetyFife, humanize.Bytes(uint64(entry.NinetyFife)),
			width.ninetyNine, humanize.Bytes(uint64(entry.NinetyNine)),
			width.totalSize, humanize.Bytes(uint64(entry.TotalBytes)),
		))
	}

	list.list.Rows = result
}

func (list *RoutingKeyList) renderHeader(width ListColumnWidhts) string {
	rkH := "Routing Key"
	countH := "Count"
	avgH := "AVG"
	nfH := "95%"
	nnH := "99%"
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
	case RK_SORTING_AVG_ASC:
		avgH += " ▲"
	case RK_SORTING_AVG_DESC:
		avgH += " ▼"
	case RK_SORTING_NINETY_FIFE_ASC:
		nfH += " ▲"
	case RK_SORTING_NINETY_FIFE_DESC:
		nfH += " ▼"
	case RK_SORTING_NINETY_NINE_ASC:
		nnH += " ▲"
	case RK_SORTING_NINETY_NINE_DESC:
		nnH += " ▼"
	case RK_SORTING_TOTAL_SIZE_ASC:
		sizeH += " ▲"
	case RK_SORTING_TOTAL_SIZE_DESC:
		sizeH += " ▼"
	}

	return fmt.Sprintf(
		"[%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white) [%-*s](fg:white)",
		width.routingKey, rkH,
		width.count, countH,
		width.avg, avgH,
		width.ninetyFife, nfH,
		width.ninetyNine, nnH,
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
