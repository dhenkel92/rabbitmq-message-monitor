package ui

import (
	"regexp"

	termui "github.com/gizak/termui/v3"
)

var SEARCH_KEY_REGEX = regexp.MustCompile(`^<.*>$`)

func (ui *UI) handleKeyPress(e *termui.Event) bool {
	if ui.isSearchActive {
		ui.search(e)
		ui.list.FilterName(ui.searchBar.GetExpression())
		ui.Render()
		return false
	}
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
	case "/":
		ui.isSearchActive = true
		ui.searchBar.SetSeachActive(true)
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

func (ui *UI) search(e *termui.Event) {
	if SEARCH_KEY_REGEX.Match([]byte(e.ID)) {
		switch e.ID {
		case "<Backspace>":
			ui.searchBar.RemoveLast()
		case "<Escape>":
			fallthrough
		case "<Enter>":
			ui.isSearchActive = false
			ui.searchBar.SetSeachActive(false)
		}
		return
	}

	ui.searchBar.AddKey(e.ID)
}
