package ui

import (
	"fmt"
	"strings"

	"github.com/dhenkel92/rabbitmq-message-monitor/internal/helper"
	"github.com/dhenkel92/rabbitmq-message-monitor/internal/types"
	"github.com/gizak/termui/v3/widgets"
)

func newSettingsHeader(size types.WidgetSize) *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Title = "RabbitMQ Message Monitor"
	header.Text = formatInformation(1, [][]string{
		{"Refresh rate:", "10s"},
	})
	header.SetRect(size.StartX, size.StartY, size.EndX, size.EndY)
	header.Border = true
	return header
}

func newShortcutHeader(size types.WidgetSize) *widgets.Paragraph {
	header := widgets.NewParagraph()
	header.Title = "Shortcurts"
	header.Text = formatInformation(2, [][]string{
		{"[<k>](fg:blue)", "up"},
		{"[<gg>](fg:blue)", "top"},
		{"[<j>](fg:blue)", "down"},
		{"[<G>](fg:blue)", "bottom"},
		{"[](fg:blue)", ""},
		{"[](fg:blue)", ""},
		{"[</>](fg:blue)", "search"},
		{"[<s>](fg:blue)", "change sorting"},
		{"[<q>](fg:blue)", "quit"},
	})
	header.SetRect(size.StartX, size.StartY, size.EndX, size.EndY)
	header.Border = true
	return header
}

func formatInformation(columns int, data [][]string) string {
	var column1Wdth, column2Wdt int
	for _, row := range data {
		if len(row[0]) > column1Wdth {
			column1Wdth = len(row[0])
		}
		if len(row[1]) > column2Wdt {
			column2Wdt = len(row[1])
		}
	}

	res := make([]string, 0)

	for _, entry := range data {
		res = append(res, fmt.Sprintf("%*s %-*s", column1Wdth, entry[0], column2Wdt, entry[1]))
	}

	var tmpRes []string
	chunks := helper.SplitStringArray(res, columns)
	for _, chunk := range chunks {
		tmpRes = append(tmpRes, strings.Join(chunk, "     "))
	}
	return strings.Join(tmpRes, "\n")
}
