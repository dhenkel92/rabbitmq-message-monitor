package uiRKList

import (
	"sort"
)

type RoutingKeySorting int

const (
	RK_SORTING_NAME_ASC         RoutingKeySorting = iota
	RK_SORTING_NAME_DESC        RoutingKeySorting = iota
	RK_SORTING_COUNT_ASC        RoutingKeySorting = iota
	RK_SORTING_COUNT_DESC       RoutingKeySorting = iota
	RK_SORTING_AVG_ASC          RoutingKeySorting = iota
	RK_SORTING_AVG_DESC         RoutingKeySorting = iota
	RK_SORTING_NINETY_FIFE_ASC  RoutingKeySorting = iota
	RK_SORTING_NINETY_FIFE_DESC RoutingKeySorting = iota
	RK_SORTING_NINETY_NINE_ASC  RoutingKeySorting = iota
	RK_SORTING_NINETY_NINE_DESC RoutingKeySorting = iota
	RK_SORTING_TOTAL_SIZE_ASC   RoutingKeySorting = iota
	RK_SORTING_TOTAL_SIZE_DESC  RoutingKeySorting = iota

	RK_DEFAULT_SORTING RoutingKeySorting = RK_SORTING_COUNT_DESC
)

func (list *RoutingKeyList) Sort(sorting RoutingKeySorting) {
	list.sorting = sorting
	list.Render()
}

func sortData(sorting RoutingKeySorting, data []*RoutingKeyData) []*RoutingKeyData {
	switch sorting {
	case RK_SORTING_NAME_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].RoutingKey < data[j].RoutingKey })
	case RK_SORTING_NAME_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].RoutingKey > data[j].RoutingKey })
	case RK_SORTING_COUNT_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].Count < data[j].Count })
	case RK_SORTING_COUNT_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].Count > data[j].Count })
	case RK_SORTING_AVG_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].Avg < data[j].Avg })
	case RK_SORTING_AVG_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].Avg > data[j].Avg })
	case RK_SORTING_NINETY_FIFE_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].NinetyFife < data[j].NinetyFife })
	case RK_SORTING_NINETY_FIFE_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].NinetyFife > data[j].NinetyFife })
	case RK_SORTING_NINETY_NINE_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].NinetyNine < data[j].NinetyNine })
	case RK_SORTING_NINETY_NINE_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].NinetyNine > data[j].NinetyNine })
	case RK_SORTING_TOTAL_SIZE_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].TotalBytes < data[j].TotalBytes })
	case RK_SORTING_TOTAL_SIZE_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].TotalBytes > data[j].TotalBytes })
	}
	return data
}

func (list *RoutingKeyList) SortNext() {
	list.Sort((list.sorting + 1) % 12)
}
