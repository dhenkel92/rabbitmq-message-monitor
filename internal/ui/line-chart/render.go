package lineChart

func (chart *LineChart) Render() {
	if len(chart.data) == 1 {
		return
	}
	chart.uiElement.Data = [][]float64{chart.data}
}
