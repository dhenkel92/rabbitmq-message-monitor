package ui

import (
	"fmt"

	termui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const headerHeight = 6

type UI struct {
	settingsHeader *widgets.Paragraph
	shortcutHeader *widgets.Paragraph
	list           *widgets.List

	previousKey string
}

func NewUI() (*UI, error) {
	if err := termui.Init(); err != nil {
		return nil, fmt.Errorf("failed to initialize termui: %v", err)
	}

	termWidth, termHeight := termui.TerminalDimensions()

	settingsHeader := newSettingsHeader(WidgetSize{startX: 0, startY: 0, width: termWidth / 100 * 30, height: headerHeight})
	shortcutHeader := newShortcutHeader(WidgetSize{startX: termWidth / 100 * 30, startY: 0, width: termWidth / 100 * 70, height: headerHeight})

	return &UI{
		settingsHeader: settingsHeader,
		shortcutHeader: shortcutHeader,
		list:           newList("Routing Keys", termWidth, termHeight, headerHeight),
	}, nil
}

func (ui *UI) Render() {
	termui.Render(ui.settingsHeader, ui.shortcutHeader, ui.list)
}

func (ui *UI) Start() error {
	defer termui.Close()

	ui.Render()
	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		if ui.handleKeyPress(&e) {
			return nil
		}
	}
}

func (ui *UI) LineLength() int {
	return ui.list.Inner.Max.X
}

func (ui *UI) UpdateListEntries(entries []string) {
	ui.list.Rows = entries
	ui.Render()
}
