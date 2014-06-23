package substringfinder

import "strings"

func FindFirstOfSubString(parent, substring string) (int, int) {
	return FindFirstOfSubStringWithStartingIndex(parent, substring, 0)
}

func FindFirstOfSubStringWithStartingIndex(parent, substring string, startingIndex int) (int, int) {
	for i := startingIndex; i < len(parent)-len(substring); i++ {
		foundSubString := parent[i] == substring[0]
		for counter := 1; counter < len(substring); counter++ {
			if foundSubString == false {
				break
			}
			foundSubString = parent[i+counter] == substring[counter]
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
	rboundStartingPos := startingIndex + lboundIndex + 1
	substring = parent[rboundStartingPos:len(parent)]
	rboundIndex := strings.IndexRune(substring, rightRune)
	return lboundIndex + startingIndex, rboundIndex + rboundStartingPos
}
