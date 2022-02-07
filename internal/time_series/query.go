package timeseries

type Query struct {
	Metric string
	Labels TimeSeriesLabels
	// How many entries should be returned
	// -1 == all
	Limit int

	From  int64
	Until int64
}

func (db *TimeSeriesDB) Query(query Query) []*TimeSeries {
	result := make([]*TimeSeries, 0)
	for _, ts := range db.series {
		if ts.metric == query.Metric && includesLabels(ts.labels, query.Labels) {
			result = append(result, ts.getTimeRange(query.From, query.Until))
		}
		if query.Limit >= 0 && len(result) == query.Limit {
			break
		}
	}
	return result
}

func (ts *TimeSeries) getTimeRange(from, until int64) *TimeSeries {
	fromIdx := ts.findTimestampIdx(from)
	untilIdx := ts.findTimestampIdx(until)
	return &TimeSeries{
		metric:     ts.metric,
		labels:     ts.labels,
		timestamps: ts.timestamps[fromIdx:untilIdx],
		values:     ts.values[fromIdx:untilIdx],
	}
}
