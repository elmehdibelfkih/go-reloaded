package internal

import (
	"bufio"
	"fmt"
	"go-reloaded/pkg"
	"os"
	"strconv"
	"strings"
)

func HandelLine(input *os.File, output *os.File) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		line = hexHandler(line)
		line = binHandler(line)
		line = upHandler(line)
		line = lowHandler(line)
		line = capHandler(line)
		line = punctuationsHandler(line)
		output.WriteString(line)
	}
}

func hexHandler(line string) string {
	index := strings.Index(line, "(hex)")
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		intValue, err := strconv.ParseInt(word, 16, 0)
		if err == nil {
			line = pkg.ReplaceAtIndex(line, word, strconv.Itoa(int(intValue)), start)
			line = strings.Replace(line, "(hex)", "", 1)

		} else {
			fmt.Println("Error:", err)
			return line
		}
		index = strings.Index(line, "(hex)")
	}
	return line
}

func binHandler(line string) string {
	index := strings.Index(line, "(bin)")
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		intValue, err := strconv.ParseInt(word, 2, 0)
		if err == nil {
			line = pkg.ReplaceAtIndex(line, word, strconv.Itoa(int(intValue)), start)
			line = strings.Replace(line, "(bin)", "", 1)

		} else {
			fmt.Println("Error:", err)
			return line
		}
		index = strings.Index(line, "(bin)")
	}
	return line
}

func upHandler(line string) string { //FIXME: EX: (up, 9)
	index := strings.Index(line, "(up")
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, strings.ToUpper(word) , start)
		line = strings.Replace(line, "(up)", "", 1)
		index = strings.Index(line, "(up)")
	}
	return line
}

func lowHandler(line string) string { //FIXME:
	index := strings.Index(line, "(low)")
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, strings.ToLower(word) , start)
		line = strings.Replace(line, "(low)", "", 1)
		index = strings.Index(line, "(low)")
	}
	return line
}

func capHandler(line string) string { //FIXME:
	index := strings.Index(line, "(cap)")
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		new := []rune(strings.ToLower(word))
		new[0] = rune(strings.ToUpper(string(new[0]))[0])
		line = pkg.ReplaceAtIndex(line, word, string(new) , start)
		line = strings.Replace(line, "(cap)", "", 1)
		index = strings.Index(line, "(cap)")
	}
	return line
}

func punctuationsHandler(line string) string { //FIXME:
	return line
}
