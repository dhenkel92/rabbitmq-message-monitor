package helper

func SplitStringArray(data []string, chunkSize int) [][]string {
	var divided [][]string

	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize

		if end > len(data) {
			end = len(data)
		}

		divided = append(divided, data[i:end])
	}
	return divided
}
