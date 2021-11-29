package history

func New() HistoryRecorder {
	return HistoryRecorder{
		history:    make(map[string][]*HistoryEntry),
		lastRecord: make(map[string]float64),
	}
}
