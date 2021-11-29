package ui

import termui "github.com/gizak/termui/v3"

func (ui *UI) handleKeyPress(e *termui.Event) bool {
	switch e.ID {
	case "q", "<C-c>":
		return true
	case "j", "<Down>":
		ui.list.ScrollDown()
	case "k", "<Up>":
		ui.list.ScrollUp()
	case "<C-d>":
		ui.list.ScrollHalfPageDown()
	case "<C-u>":
		ui.list.ScrollHalfPageUp()
	case "<C-f>":
		ui.list.ScrollPageDown()
	case "<C-b>":
		ui.list.ScrollPageUp()
	case "<Home>":
		ui.list.ScrollTop()
	case "g":
		if ui.previousKey == "g" {
			ui.list.ScrollTop()
		}
	case "G", "<End>":
		ui.list.ScrollBottom()
	}

	if ui.previousKey == "g" {
		ui.previousKey = ""
	} else {
		ui.previousKey = e.ID
	}

	ui.Render()
	return false
}
