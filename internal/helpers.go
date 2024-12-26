package internal

import "strings"

func parsFlag(line string, mode string, start int) (err bool, rep int) {
	var end int

	for i, c := range line {
		if c == ')' {
			end = i
		}
	}
	if end == 0 {
		return false, 0
	}

	flag := line[start:end]
	flag = strings.TrimSpace(flag)
	flag = strings.TrimLeft(flag, "(" + mode)
	flag = strings.TrimRight(flag, ")")
	if flag == "" {
		return 1
	}
	if strings.Contains(flag, ",")
	return false, 0
}