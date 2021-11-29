package search

import "github.com/gizak/termui/v3/widgets"

type SearchBar struct {
	uiElement      *widgets.Paragraph
	searchTerm     string
	isSearchActive bool
}

func (bar *SearchBar) GetUiElement() *widgets.Paragraph {
	return bar.uiElement
}
