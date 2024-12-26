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
		output.WriteString(line + "\n")
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

func upHandler(line string) string {
	index := strings.Index(line, "(up")
	for index != -1 {
		err, rep, rm := parsFlag(line, "up", index)
		if err {
			fmt.Println("Error: up syntax!!")
		} else {
			for ;rep != 0; rep-- {
				word, start := pkg.PreviousWord(line, index)
				line = pkg.ReplaceAtIndex(line, word, strings.ToUpper(word) , start)
				line = strings.Replace(line, rm, "", 1)
				index = start
			}
		}
		index = strings.Index(line, "(up")
	}
	return line
}

func lowHandler(line string) string {
	index := strings.Index(line, "(low")
	for index != -1 {
		err, rep, rm := parsFlag(line, "low", index)
		if err {
			fmt.Println("Error: low syntax!!")
		} else {
			for ;rep != 0; rep-- {
				word, start := pkg.PreviousWord(line, index)
				line = pkg.ReplaceAtIndex(line, word, strings.ToLower(word) , start)
				line = strings.Replace(line, rm, "", 1)
				index = start
			}
		}
		index = strings.Index(line, "(low")
	}
	return line
}

func capHandler(line string) string { //FIXME:
	index := strings.Index(line, "(cap")
	for index != -1 {
		err, rep, rm := parsFlag(line, "cap", index)
		if err {
			fmt.Println("Error: cap syntax!!")
		} else {
			for ;rep != 0; rep-- {
				word, start := pkg.PreviousWord(line, index)
				line = pkg.ReplaceAtIndex(line, word, pkg.CapWord(word) , start)
				line = strings.Replace(line, rm, "", 1)
				index = start
			}
		}
		index = strings.Index(line, "(cap")
	}
	return line
}

func punctuationsHandler(line string) string { //FIXME:
	return line
}
