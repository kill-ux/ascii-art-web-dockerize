package ascii

func LastLine(strTab []string) int {
	index := -1
	for i, word := range strTab {
		if word != "\n" {
			index = i
		}
	}
	if index != -1 {
		return index
	}
	return 0
}
