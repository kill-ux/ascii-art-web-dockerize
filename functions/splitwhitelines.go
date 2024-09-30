package ascii

func SplitWhitelines(s string) []string {
	var result []string
	var currentLine string

	for i := 0; i < len(s); i++ {
		if s[i] == 13 {
			continue
		}
		/* if s[i] == '\\' && i < len(s)-1 && s[i+1] == 'n' {
			if len(currentLine) > 0 {
				result = append(result, currentLine)
			}
			result = append(result, "\n")
			currentLine = ""
			i++ // Skip the next character . it's part of the "\n"
		} else  */
		if s[i] == 10 {
			if len(currentLine) > 0 {
				result = append(result, currentLine)
			}
			result = append(result, "\n")
			currentLine = ""
		} else {
			currentLine += string(s[i])
		}
	}

	// Append the last line if there is any
	if len(currentLine) > 0 {
		result = append(result, currentLine)
	}
	return result
}
