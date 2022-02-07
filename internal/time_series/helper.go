package timeseries

import (
	"fmt"
	"sort"
	"strings"
)

func createKeyForDB(metric string, labels TimeSeriesLabels) string {
	accLabels := make([]string, len(labels))
	for key, val := range labels {
		accLabels = append(accLabels, fmt.Sprintf("%s=%s", key, val))
	}
	sort.Strings(accLabels)
	return fmt.Sprintf("%s_%s", metric, strings.Join(accLabels, "_"))
}

func (ts *TimeSeries) findTimestampIdx(timestamp int64) int {
	return sort.Search(len(ts.timestamps), func(i int) bool {
		return ts.timestamps[i] >= timestamp
	})
}

func includesLabels(labels, search TimeSeriesLabels) bool {
	var includeCounter int
	for key, val := range search {
		if val2, ok := labels[key]; ok && val == val2 {
			includeCounter += 1
		}
	}
	return includeCounter == len(search)
}
