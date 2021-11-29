package ui

import termui "github.com/gizak/termui/v3"

func (ui *UI) handleKeyPress(e *termui.Event) bool {
	//todo: remove GetUiElement() hack
	switch e.ID {
	case "q", "<C-c>":
		return true
	case "j", "<Down>":
		ui.list.GetUiElement().ScrollDown()
	case "k", "<Up>":
		ui.list.GetUiElement().ScrollUp()
	case "<C-d>":
		ui.list.GetUiElement().ScrollHalfPageDown()
	case "<C-u>":
		ui.list.GetUiElement().ScrollHalfPageUp()
	case "<C-f>":
		ui.list.GetUiElement().ScrollPageDown()
	case "<C-b>":
		ui.list.GetUiElement().ScrollPageUp()
	case "s":
		ui.list.SortNext()
	case "<Home>":
		ui.list.GetUiElement().ScrollTop()
	case "g":
		if ui.previousKey == "g" {
			ui.list.GetUiElement().ScrollTop()
		}
	case "G", "<End>":
		ui.list.GetUiElement().ScrollBottom()
	}

	if ui.previousKey == "g" {
		ui.previousKey = ""
	} else {
		ui.previousKey = e.ID
	}

	ui.Render()
	return false
}
