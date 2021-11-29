package lineChart

import "github.com/gizak/termui/v3/widgets"

type LineChart struct {
	uiElement *widgets.Plot
	data      []float64
}

func (chart *LineChart) GetUiElement() *widgets.Plot {
	return chart.uiElement
}
