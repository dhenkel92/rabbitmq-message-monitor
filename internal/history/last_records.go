package history

func (recorder *HistoryRecorder) SetLatest(key RecordKey, value float64) {
	recorder.lastRecord[string(key)] = value
}

func (recorder *HistoryRecorder) GetLatest(key RecordKey) float64 {
	return recorder.lastRecord[string(key)]
}
