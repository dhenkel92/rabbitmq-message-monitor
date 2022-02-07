package timeseries

type TimeSeriesDB struct {
	series map[string]*TimeSeries
}

func New() *TimeSeriesDB {
	return &TimeSeriesDB{
		series: make(map[string]*TimeSeries),
	}
}
