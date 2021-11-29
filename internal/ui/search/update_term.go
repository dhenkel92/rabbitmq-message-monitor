package search

func (search *SearchBar) SetSeachActive(isActive bool) {
	search.isSearchActive = isActive
	search.Render()
}

func (search *SearchBar) AddKey(key string) {
	search.searchTerm += key
	search.Render()
}

func (search *SearchBar) RemoveLast() {
	if search.searchTerm == "" {
		return
	}
	search.searchTerm = search.searchTerm[:(len(search.searchTerm) - 1)]
	search.Render()
}

func (search *SearchBar) GetExpression() string {
	return search.searchTerm
}
