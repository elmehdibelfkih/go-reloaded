package pkg

import (
	// "os"
	"strings"
)

func PreviousWord(line string, index int) (string, int) {
	runes := []rune(line)
	if index == 0 {
		return "", 0
	}
	for index != 0 && strings.Contains(`.!?,:; `, string(runes[index-1])) {
		index--
	}
	if index == 0 {
		return "", 0
	}
	i := index
	for i != 0 && !strings.Contains(`.!?,:; `, string(runes[i-1])) {
		i--
	}
	return string(runes[i:index]), i
}


func ReplaceAtIndex(line string, old string, new string, index int) string {
	return line[0:index] + new + line[index+len(old):]
}

func CapWord(word string) string {
	if word == "" {
		return ""
	}
	new := []rune(strings.ToLower(word))
	new[0] = rune(strings.ToUpper(string(new[0]))[0])
	return string(new)
}

func NextWord(index int, line string) string {
	runes := []rune(line)

	for index != len(runes)-1 && strings.Contains(`.!?,:; `, string(runes[index])) {
		index++
	}
	i := index
	for i != len(runes)-1 && !strings.Contains(`.!?,:; `, string(runes[i])) {
		i++
	}
	return line[index:i]
}
