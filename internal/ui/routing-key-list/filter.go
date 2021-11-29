package uiRKList

import "regexp"

func (list *RoutingKeyList) FilterName(expression string) {
	list.nameFilterExpression = expression
	list.Render()
}

func filterData(expression string, data []*RoutingKeyData) []*RoutingKeyData {
	re := regexp.MustCompile(expression)
	filtered := make([]*RoutingKeyData, 0)

	for _, entry := range data {
		if re.Match([]byte(entry.RoutingKey)) {
			filtered = append(filtered, entry)
		}
	}

	return filtered
}
