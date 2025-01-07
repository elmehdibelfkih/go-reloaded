package pkg

import (
	"strings"
)

func PreviousWord(line string, index int) (string, int) {
	runes := []rune(line)
	if index <= 0 || index > len(runes) {
		return "", 0
	}
	i := index - 1
	for i >= 0 && strings.ContainsRune(`.!?,:; `, runes[i]) {
		i--
	}
	if i < 0 {
		return "", 0
	}
	end := i + 1
	for i >= 0 && !strings.ContainsRune(`.!?,:; `, runes[i]) {
		i--
	}
	start := i + 1
	return string(runes[start:end]), start
}

func ReplaceAtIndex(line string, old string, new string, index int) string {
	runes := []rune(line)
	oldRunes := []rune(old)
	newRunes := []rune(new)
	if index < 0 || index+len(oldRunes) > len(runes) || string(runes[index:index+len(oldRunes)]) != old {
		return line
	}
	return string(runes[:index]) + string(newRunes) + string(runes[index+len(oldRunes):])
}

func CapWord(word string) string {
	if word == "" {
		return ""
	}
	new := []rune(strings.ToLower(word))
	new[0] = []rune(strings.ToUpper(string(new[0])))[0]
	return string(new)
}

func NextWord(index int, line string) string {
	runes := []rune(line)
	for index != len(runes)-1 && !strings.Contains(`.!?,:; `, string(runes[index])) {
		index++
	}
	for index != len(runes)-1 && strings.Contains(`.!?,:; `, string(runes[index])) {
		index++
	}
	i := index
	for i != len(runes)-1 && !strings.Contains(`.!?,:; `, string(runes[i])) {
		i++
	}
	return string(runes[index : i+1])
}

func RuneIndex(s string, target string) int {
	runes := []rune(s)
	targetRunes := []rune(target)

	if len(targetRunes) == 0 {
		return -1
	}
	for i := 0; i <= len(runes)-len(targetRunes); i++ {
		match := true
		for j := 0; j < len(targetRunes); j++ {
			if runes[i+j] != targetRunes[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}
