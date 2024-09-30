package ascii

import (
	"errors"
)

func Banner(str string) bool {
	switch str {
	case "shadow":
		return true
	case "standard":
		return true
	case "thinkertoy":
		return true
	case "chap":
		return true
	default:
		return false
	}
}

func Art(str string, banner string) ([]byte, error) {
	// string not printable so it exit
	for _, char := range str {
		if (char < 32 || char > 126) && char != 10 && char != 13 {
			err := errors.New("not printable")
			return []byte{}, err
		}
	}

	str2 := []byte{}
	for _, char := range []byte(str) {
		if char != '\r' {
			str2 = append(str2, char)
		}
	}

	// split any words with '\n' and add it to []string
	strTab := SplitWhitelines(string(str2))

	if !Banner(banner) {
		return []byte{}, errors.New("this banner is not exists")
	}

	// reade file
	data, err := ReadF(banner)
	if err != nil {
		return []byte{}, err
	}

	// return the index of last line
	index := LastLine(strTab)
	// remove '\n' after the word if it's not the last word
	for i := 0; i < index; i++ {
		if i != 0 && i < len(strTab)-1 && strTab[i] == "\n" && strTab[i-1] != "\n" {
			strTab = append(strTab[:i], strTab[i+1:]...)
		}
	}

	dataOut := []byte{}
	// range the table
	for _, word := range strTab {
		Tab := [][]byte{}

		// if the word equal '\n'  print line
		if word == "\n" {
			dataOut = append(dataOut, '\n')
		} else {
			// the ascii of the characteres change to the index in the data 'Tab'
			for _, char := range word {
				Tab = append(Tab, data[char-32])
			}
			end := false
			// range the table
			for i := 0; i < len(Tab); i++ {
				// line
				stdChar := Tab[i]
				for j := 0; j < len(stdChar); j++ {
					if stdChar[j] != 10 {
						dataOut = append(dataOut, stdChar[j])
					} else {
						// last char(article) and last byte in the char
						if i == len(Tab)-1 && j == len(stdChar)-1 {
							end = true
						}
						// delete that printed
						if len(stdChar) > 0 {
							Tab[i] = stdChar[j+1:]
						}
						break
					}
				}
				// last char so \n
				if i == len(Tab)-1 {
					dataOut = append(dataOut, '\n')
					// return the first char(article)
					i = -1
				}
				if end {
					break
				}
			}
		}

	}
	return dataOut, nil
}
