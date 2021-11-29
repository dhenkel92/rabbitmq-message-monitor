package search

import "github.com/gizak/termui/v3/widgets"

type SearchBar struct {
	uiElement  *widgets.Paragraph
	searchTerm string
}

func (bar *SearchBar) GetUiElement() *widgets.Paragraph {
	return bar.uiElement
}
