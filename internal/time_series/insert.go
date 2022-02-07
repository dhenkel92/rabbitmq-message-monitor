package timeseries

type InsertRow struct {
	Metic     string
	Labels    TimeSeriesLabels
	Value     int64
	Timestamp int64
}

func (db *TimeSeriesDB) InsertMany(rows []InsertRow) {
	for _, row := range rows {
		db.Insert(row)
	}
}

func (db *TimeSeriesDB) Insert(row InsertRow) {
	key := createKeyForDB(row.Metic, row.Labels)
	if _, ok := db.series[key]; !ok {
		db.series[key] = newTimeSeries(row.Metic, row.Labels)
	}

	db.series[key].insert(row.Value, row.Timestamp)
}

func (ts *TimeSeries) insert(value int64, timestamp int64) {
	ts.expand()
	ts.mu.Lock()
	smallestIdx := ts.findTimestampIdx(timestamp)

	// if the idx is -1 it means that the timestamp should be appended to the end
	if smallestIdx == -1 || smallestIdx >= len(ts.timestamps) {
		ts.timestamps = append(ts.timestamps, timestamp)
		ts.values = append(ts.values, value)
		ts.mu.Unlock()
		return
	}

	// if the smallestIdx does not match the timestamp, it means that the timestamp wasn't inserted yet
	if ts.timestamps[smallestIdx] != timestamp {
		ts.expand()

		// we are moving everything to the right, starting at the smallestIdx
		// afterwards we can just ingest it at smallestIdx
		ts.timestamps = append(ts.timestamps[:smallestIdx+1], ts.timestamps[smallestIdx:]...)
		ts.values = append(ts.values[:smallestIdx+1], ts.values[smallestIdx:]...)

		ts.timestamps[smallestIdx] = timestamp
		ts.values[smallestIdx] = value

		ts.mu.Unlock()
		return
	}

	// todo: check if we want to accumulate data
	// if the value at smallestIdx matches the exact same timestamp, well just add up the data
	ts.values[smallestIdx] += value
	ts.mu.Unlock()
}
