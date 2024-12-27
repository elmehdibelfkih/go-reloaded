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
		line = hexHandler(line)   // FIXME:
		line = binHandler(line)   // FIXME:
		line = orderReplace(line) // TODO:
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

func upHandler(line string, index int) string {
	_, rep, rm := parsFlag(line, "up", index)
	for ; rep != 0; rep-- {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, strings.ToUpper(word), start)
		line = strings.Replace(line, rm, "", 1)
		index = start
	}
	return line
}

func lowHandler(line string, index int) string {
	_, rep, rm := parsFlag(line, "low", index)
	for ; rep != 0; rep-- {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, strings.ToLower(word), start)
		line = strings.Replace(line, rm, "", 1)
		index = start
	}
	return line
}

func capHandler(line string, index int) string {
	_, rep, rm := parsFlag(line, "cap", index)
	for ; rep != 0; rep-- {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, pkg.CapWord(word), start)
		line = strings.Replace(line, rm, "", 1)
		index = start
	}
	return line
}

func flagHandler(line string, index int, mode string, opp func(string) string) string {
	_, rep, rm := parsFlag(line, mode, index)
	for ; rep != 0; rep-- {
		word, start := pkg.PreviousWord(line, index)
		line = pkg.ReplaceAtIndex(line, word, opp(word), start)
		line = strings.Replace(line, rm, "", 1)
		index = start
	}
	return line
}

func punctuationsHandler(line string) string { //FIXME:
	return line
}
