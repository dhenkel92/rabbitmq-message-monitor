package uiRKList

import "regexp"

func (list *RoutingKeyList) FilterName(expression string) {
	list.nameFilterExpression = expression
	list.Render()
}

func filterData(expression string, data []*RoutingKeyData) []*RoutingKeyData {
	filtered := make([]*RoutingKeyData, 0)
	re, err := regexp.Compile(expression)

	if err != nil {
		// todo: show error to user
		return filtered
	}

	for _, entry := range data {
		if re.Match([]byte(entry.RoutingKey)) {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}
