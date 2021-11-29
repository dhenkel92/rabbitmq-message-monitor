package search

import "fmt"

func (search *SearchBar) Render() {
	if search.searchTerm == "" {
		search.uiElement.Text = " > [Start your search now](fg:grey)"
		return
	}

	search.uiElement.Text = fmt.Sprintf(" > %s", search.searchTerm)
}
