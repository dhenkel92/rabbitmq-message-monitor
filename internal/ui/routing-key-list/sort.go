package uiRKList

import (
	"sort"
)

func (list *RoutingKeyList) Sort(sorting RoutingKeySorting) {
	list.sorting = sorting
	list.data = sortData(list.sorting, list.data)
	list.renderList()
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
	case RK_SORTING_TOTAL_SIZE_ASC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].TotalBytes < data[j].TotalBytes })
	case RK_SORTING_TOTAL_SIZE_DESC:
		sort.SliceStable(data, func(i, j int) bool { return data[i].TotalBytes > data[j].TotalBytes })
	}
	return data
}

func (list *RoutingKeyList) SortNext() {
	list.Sort((list.sorting + 1) % 6)
}
