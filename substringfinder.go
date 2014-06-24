package substringfinder

import (
	"strings"

//"math"
//"fmt"
)

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
	substring := parent[startingIndex:]
	lboundIndex := strings.IndexRune(substring, leftRune)
	if lboundIndex > -1 {
		rboundStartingPos := startingIndex + lboundIndex + 1
		substring = parent[rboundStartingPos:]
		rboundIndex := strings.IndexRune(substring, rightRune)
		if rboundIndex > -1 {
			return lboundIndex + startingIndex, rboundIndex + rboundStartingPos
		}
	}
	return -1, -1
}

func FindIndicesBetweenMatchingRunes(parent string, leftRune, rightRune rune, ignoreQuotsAndComments bool) (int, int) {
	return FindIndicesBetweenMatchingRunesWithStartingIndex(parent, leftRune, rightRune, 0, ignoreQuotsAndComments)
}

func FindIndicesBetweenMatchingRunesWithStartingIndex(parent string, leftRune, rightRune rune, startingIndex int, ignoreQuotsAndComments bool) (int, int) {
	substring := parent[startingIndex:]
	foundLeftRune := false
	lbounds, rbounds := -1, -1
	counter := 0
	isValidCode := true
	var invalidationType string
	for i, r := range substring {
		//validate code block
		if isValidCode {
			if r == '/' && i < len(substring)-1 {
				if substring[i+1] == '/' {
					invalidationType = "//"
					isValidCode = false
					//println("single line comment found, ignoring next runes")
					continue
				} else if substring[i+1] == '*' {
					invalidationType = "/*"
					isValidCode = false
					//println("multi line comment found, ignoring next runes")
					continue
				}
			} else if r == '"' && i < len(substring)-1 {
				invalidationType = "\""
				isValidCode = false
				//println("string text found, ignoring next runes")
				continue
			}
		} else {
			if invalidationType == "//" && r == '\n' {
				isValidCode = true
				invalidationType = ""
				//println("single line comment ended, processing next runes")
				continue
			} else if invalidationType == "/*" && r == '/' {
				if substring[i-1] == '*' {
					isValidCode = true
					invalidationType = ""
					//println("multi line comment ended, processing next runes")
					continue
				}
			} else if invalidationType == "\"" && r == '"' {
				if substring[i-1] != '\\' {
					isValidCode = true
					invalidationType = ""
					//println("string text ended, processing next runes")
					continue
				}
			}
		}
		//don't process any invalid code if ignore bool has been set
		if isValidCode == false && ignoreQuotsAndComments == true {
			continue
		}
		if foundLeftRune {
			if r == leftRune {
				counter++
			} else if r == rightRune {
				counter--
			}
		}
		if foundLeftRune == false && r == leftRune && isValidCode {
			foundLeftRune = true
			lbounds = i
			counter++
		}
		if foundLeftRune && counter == 0 {
			rbounds = i - 1
			break
		}
	}
	return lbounds + startingIndex, rbounds + startingIndex
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
