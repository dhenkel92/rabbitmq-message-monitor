package history

import (
	"sync"
	"time"
)

type HistoryEntry struct {
	Timestamp time.Time
	Value     float64
}

type HistoryRecorder struct {
	mu         sync.Mutex
	history    map[string][]*HistoryEntry
	lastRecord map[string]float64
}
