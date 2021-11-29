package lineChart

import (
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func New(title string, size types.WidgetSize) LineChart {
	return LineChart{
		uiElement: newPlot(title, size),
	}
}

func newPlot(title string, size types.WidgetSize) *widgets.Plot {
	plot := widgets.NewPlot()
	plot.SetRect(size.StartX, size.StartY, size.EndX, size.EndY)
	plot.AxesColor = termui.ColorWhite
	plot.LineColors[0] = termui.ColorCyan
	plot.HorizontalScale = 5
	plot.Title = title
	plot.DrawDirection = widgets.DrawRight
	return plot
}
