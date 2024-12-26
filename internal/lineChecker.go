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
		word := pkg.PreviousWord(line, index)
		intValue, err := strconv.ParseInt(word, 16, 0)
		if err == nil {
			line = strings.Replace(line, word, strconv.Itoa(int(intValue)), 1)
			line = strings.Replace(line, "(hex)", "", 1)

		} else {
			fmt.Println("Error:", err)
		}
		index = strings.Index(line, "(hex)")
	}
	return line
}

func binHandler(line string) string {
	return line
}

func upHandler(line string) string {
	return line
}

func lowHandler(line string) string {
	return line
}

func capHandler(line string) string {
	return line
}

func punctuationsHandler(line string) string {
	return line
}