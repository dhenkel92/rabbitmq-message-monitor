package lineChart

func (chart *LineChart) Render() {
	// the termui plotter needs to have at least two datapoints
	if len(chart.data) < 2 {
		return
	}

	data := chart.data
	if len(chart.data) >= 16 {
		// the charts are only showing 16 datapoints, so it doesn't make sense to provide more
		data = chart.data[(len(data) - 16):]
	}

	chart.uiElement.Data = [][]float64{data}
}
