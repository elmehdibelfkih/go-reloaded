package pkg

import (
	"strings"
)

func PreviousWord(line string, index int) string {
	i := index
	runes := []rune(line)
	if 
	for i != 0 && !strings.Contains(`.!?,:; `, string(runes[i] )) {
		i--
	}

	for i != 0 && !strings.Contains(`.!?,:; `, string(runes[i] )) {
		i--
	}
	return line[i:index]
}