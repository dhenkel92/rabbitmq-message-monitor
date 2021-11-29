package search

import (
	"fmt"

	"github.com/gizak/termui/v3"
)

func (search *SearchBar) Render() {
	if search.isSearchActive {
		search.uiElement.BorderStyle.Fg = termui.ColorRed
	} else {
		search.uiElement.BorderStyle.Fg = termui.ColorWhite
	}

	if search.searchTerm == "" {
		search.uiElement.Text = " > [Start your search now](fg:grey)"
		return
	}

	search.uiElement.Text = fmt.Sprintf(" > %s", search.searchTerm)
}
