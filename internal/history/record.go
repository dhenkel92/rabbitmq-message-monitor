package history

import "time"

type RecordKey string

const (
	RECORD_SUM_SIZE_KEY  RecordKey = "sum-size"
	RECORD_MSG_COUNT_KEY RecordKey = "message-count"
)

func (recorder *HistoryRecorder) RecordDiff(key RecordKey, current float64) {
	diff := current - recorder.GetLatest(key)
	recorder.Record(key, diff)
	recorder.SetLatest(key, current)
}

func (recorder *HistoryRecorder) Record(key RecordKey, value float64) {
	strKey := string(key)

	recorder.mu.Lock()
	if _, ok := recorder.history[strKey]; !ok {
		recorder.history[strKey] = make([]*HistoryEntry, 0)
	}
	recorder.history[strKey] = append(recorder.history[strKey], &HistoryEntry{
		Timestamp: time.Now(),
		Value:     value,
	})
	recorder.mu.Unlock()
}

func (recorder *HistoryRecorder) GetRecordValues(key RecordKey) []float64 {
	result := make([]float64, 0)

	recorder.mu.Lock()
	if entries, ok := recorder.history[string(key)]; ok {
		for _, entry := range entries {
			result = append(result, entry.Value)
		}
	}
	recorder.mu.Unlock()

	return result
}
