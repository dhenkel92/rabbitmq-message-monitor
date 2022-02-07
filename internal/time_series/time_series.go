package timeseries

import "sync"

const TIME_SERIES_BATCH_ALLOC = 100

type TimeSeriesLabels map[string]string

type TimeSeries struct {
	mu sync.Mutex

	metric     string
	labels     TimeSeriesLabels
	timestamps []int64
	values     []int64
}

func newTimeSeries(metric string, labels TimeSeriesLabels) *TimeSeries {
	return &TimeSeries{
		metric:     metric,
		labels:     labels,
		timestamps: make([]int64, 0, TIME_SERIES_BATCH_ALLOC),
		values:     make([]int64, 0, TIME_SERIES_BATCH_ALLOC),
	}
}

func (ts *TimeSeries) expand() {
	if len(ts.timestamps) == cap(ts.timestamps) {
		return
	}

	ts.mu.Lock()
	tmp := ts.timestamps
	ts.timestamps = make([]int64, len(ts.timestamps), cap(ts.timestamps)+TIME_SERIES_BATCH_ALLOC)
	copy(ts.timestamps, tmp)

	tmp = ts.values
	ts.values = make([]int64, len(ts.values), cap(ts.values)+TIME_SERIES_BATCH_ALLOC)
	copy(ts.values, tmp)
	ts.mu.Unlock()
}

func (ts *TimeSeries) SumValues() int64 {
	var sum int64
	for _, elem := range ts.values {
		sum += elem
	}
	return sum
}
