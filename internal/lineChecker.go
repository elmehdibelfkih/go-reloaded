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
		line = binHexHandler(line, "(bin)", 2)
		line = binHexHandler(line, "(hex)", 16)
		line = orderReplace(line)
		line = punctuationsHandler(line)
		output.WriteString(line + "\n")
	}
}

func binHexHandler(line string, mode string, base int) string {
	index := strings.Index(line, mode)
	for index != -1 {
		word, start := pkg.PreviousWord(line, index)
		intValue, err := strconv.ParseInt(word, base, 0)
		if err == nil {
			line = pkg.ReplaceAtIndex(line, word, strconv.Itoa(int(intValue)), start)
			line = strings.Replace(line, mode, "", 1)

		} else {
			fmt.Println("Error:", err)
			return line
		}
		index = strings.Index(line, mode)
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
