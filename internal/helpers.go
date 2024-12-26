package internal

import (
	"strconv"
	"strings"
)

func parsFlag(line string, mode string, start int) (err bool, rep int, rm string) {
	var end int
	line = line[start:]
	for i, c := range line {
		if c == ')' {
			end = i
			break
		}
	}
	if end == 0 {
		return false, 0, ""
	}
	
	flag := line[0:end+1]
	rmv := flag
	flag = strings.Replace(flag, " ", "", -1)
	flag = strings.TrimLeft(flag, "(" + mode)
	flag = strings.TrimRight(flag, ")")
	if flag == "" {
		return false, 1, rmv
	}
	if strings.Contains(flag, ",") {
		flag = strings.TrimLeft(flag, ",")
		rep, err := strconv.ParseInt(flag, 10, 0)
		if flag == "" || err != nil {
			return true, 0, ""
		}
		return false, int(rep), rmv
	}
	return true, 0, ""
}