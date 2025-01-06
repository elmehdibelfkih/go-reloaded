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
		line = orderReplace(line)
		line = anHandler(line)
		line = punctuationsHandler(line)
		line = quoteHandler(line)
		output.WriteString(line + "\n")
	}
}

func binHexHandler(line string, mode string, base int, index int) string {

	word, start := pkg.PreviousWord(line, index)
	intValue, err := strconv.ParseInt(word, base, 0)
	if err == nil {
		line = pkg.ReplaceAtIndex(line, word, strconv.Itoa(int(intValue)), start)
		line = strings.Replace(line, mode, "", 1)
	} else {
		fmt.Println("Error:", err)
		line = strings.Replace(line, mode, "", 1)
		return line
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

func punctuationsHandler(line string) string {
	var ret string
	var isPunctuation bool

	for _, c := range line {
		if strings.Contains(`.!?,:;`, string(c)) {
			ret = strings.TrimRight(ret, " ")
			ret += string(c)
			isPunctuation = true
		} else {
			if c == ' ' && isPunctuation {
				continue
			} else {
				if isPunctuation {
					isPunctuation = false
					ret += " "
				}
				ret += string(c)

			}
		}
	}
	return ret
}

func quoteHandler(line string) string {
	var ret string
	var firstQuote bool

	quoteCounter := strings.Count(line, `'`)
	if quoteCounter < 2 {
		return line
	}
	if quoteCounter%2 != 0 {
		quoteCounter--
	}

	for _, c := range line {
		if c == '\'' && quoteCounter%2 == 0 {
			firstQuote = true
			quoteCounter--
			ret += string(c)
		} else if c == ' ' && firstQuote {
			continue
		} else if c == '\'' && quoteCounter%2 != 0 {
			ret = strings.TrimRight(ret, " ")
			ret += "'"
			quoteCounter--
			firstQuote = false
		} else {
			ret += string(c)
			firstQuote = false
		}
	}
	return ret
}

func anHandler(line string) string {
	for i := 0; i < len(line); i++ {
		if line[i] == 'a' || line[i] == 'A' {
			if i != len(line)-1 && strings.Contains(`.!?,:; `, string(line[i+1])) {
				if i == 0 || (i != 0 && strings.Contains(`.!?,:; `, string(line[i-1]))) {
					next := pkg.NextWord(i, line)
					if next != "" && strings.Contains(`aeiouAEIOUhH`, string(next[0])) {
						if line[i] == 'a' {
							line = pkg.ReplaceAtIndex(line, "a", "an", i)
						} else {
							line = pkg.ReplaceAtIndex(line, "A", "An", i)
						}
					}
				}
			}
		}
	}
	return line
}
