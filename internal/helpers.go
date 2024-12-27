package internal

import (
	"go-reloaded/pkg"
	"strconv"
	"strings"
)

const maxFlagLen = 15

func parsFlag(line string, mode string, start int) (err bool, rep int, rm string) {
	var end int
	line = line[start:]
	limeter := 0
	for i, c := range line {
		limeter++
		if c == ')' {
			end = i
			break
		}
		if limeter == maxFlagLen {
			return true, 0, ""
		}
	}
	if end == 0 {
		return false, 0, ""
	}
	flag := line[0 : end+1]
	rmv := flag
	flag = strings.Replace(flag, " ", "", -1)
	flag = strings.TrimLeft(flag, "("+mode)
	flag = strings.TrimRight(flag, ")")
	if flag == "" {
		return false, 1, rmv
	}
	if strings.Contains(flag, ",") {
		flag = strings.TrimLeft(flag, ",")
		rep, err := strconv.ParseInt(flag, 10, 0)
		if flag == "" || err != nil || rep <= 0 {
			return true, 0, ""
		}
		return false, int(rep), rmv
	}
	return true, 0, ""
}

func ValideFlagIndex(line string, subster string) int {
	var ret string

	index := strings.Index(line, subster)
	for index != -1 {
		err, rep, rm := parsFlag(line, subster, index)
		_ = rep
		_ = rm
		if err {
			ret = ret + line[0:index+3]
			line = line[index+3:]
		} else {
			return index + len(ret)
		}
		index = strings.Index(line, subster)
	}
	return -1
}

func orderReplace(line string) string {
	up := ValideFlagIndex(line, "(up")
	low := ValideFlagIndex(line, "(low")
	cap := ValideFlagIndex(line, "(cap")

	for up != -1 || low != -1 || cap != -1 {
		if up != -1 && (up <= low || low == -1) && (up <= cap || cap == -1) {
			line = flagHandler(line, up, "up", strings.ToUpper)
		} else if low != -1 && (low <= up || up == -1) && (low <= cap || cap == -1) {
			line = flagHandler(line, low, "low", strings.ToLower)
		} else if cap != -1 {
			line = flagHandler(line, cap, "cap", pkg.CapWord)
		}
		line = binHexHandler(line, "(bin)", 2)
		line = binHexHandler(line, "(hex)", 16)
		up = ValideFlagIndex(line, "(up")
		low = ValideFlagIndex(line, "(low")
		cap = ValideFlagIndex(line, "(cap")

	}
	return line
}
