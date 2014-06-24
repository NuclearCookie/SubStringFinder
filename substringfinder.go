package substringfinder

import "strings"

func FindFirstOfSubString(parent, substring string) (int, int) {
	return FindFirstOfSubStringWithStartingIndex(parent, substring, 0)
}

func FindFirstOfSubStringWithStartingIndex(parent, substring string, startingIndex int) (int, int) {
	// i < len + len causes range [startingIndex: len - 1] to be evaluated.. be sure to add 1 again
	for i := startingIndex; i < len(parent)-len(substring)+1; i++ {
		foundSubString := parent[i] == substring[0]
		if foundSubString == true {
			for counter := 1; counter < len(substring); counter++ {
				if foundSubString == false {
					break
				}
				foundSubString = parent[i+counter] == substring[counter]
			}
		}
		if foundSubString == true {
			return i, i + len(substring) - 1
		}
	}
	return -1, -1
}

func FindIndicesBetweenRunes(parent string, leftRune, rightRune rune) (int, int) {
	return FindIndicesBetweenRunesWithStartingIndex(parent, leftRune, rightRune, 0)
}

func FindIndicesBetweenRunesWithStartingIndex(parent string, leftRune, rightRune rune, startingIndex int) (int, int) {
	substring := parent[startingIndex:len(parent)]
	lboundIndex := strings.IndexRune(substring, leftRune)
	if lboundIndex > -1 {
		rboundStartingPos := startingIndex + lboundIndex + 1
		substring = parent[rboundStartingPos:len(parent)]
		rboundIndex := strings.IndexRune(substring, rightRune)
		return lboundIndex + startingIndex, rboundIndex + rboundStartingPos
	}
	return -1, -1
}

func FindIndicesBetweenRunesContaining(parent string, leftRune, rightRune rune, substring string) (int, int) {
	return FindIndicesBetweenRunesContainingWithStartingIndex(parent, leftRune, rightRune, substring, 0)
}

func FindIndicesBetweenRunesContainingWithStartingIndex(parent string, leftRune, rightRune rune, substring string, startingIndex int) (int, int) {
	start, end := 0, 0
	for start != -1 && end != -1 {
		start, end = FindIndicesBetweenRunesWithStartingIndex(parent, leftRune, rightRune, end+1)
		if start != -1 && end != -1 {
			subStart, _ := FindFirstOfSubString(parent[start:end+1], substring)
			if subStart != -1 {
				return start, end
			}
		}
	}
	return -1, -1
}
