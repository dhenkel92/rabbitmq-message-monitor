package lineChart

import "math"

func (chart *LineChart) SetBytes(data []float64) {
	for i, size := range data {
		data[i] = math.Round((size / 10240 / 1024 * 100)) / 100
	}

	chart.data = data
	chart.Render()
}

func (chart *LineChart) SetRPS(data []float64) {
	for i, size := range data {
		data[i] = math.Round((size / 1000 * 100)) / 100
	}

	chart.data = data
	chart.Render()
}
