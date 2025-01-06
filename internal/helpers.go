package internal

import (
	"go-reloaded/pkg"
	"strconv"
	"strings"
)

const maxFlagLen = 15

func parsFlag(line string, mode string, start int) (err bool, rep int, rm string) {
	var end int
	runes := []rune(line)
	runes = runes[start:]
	limeter := 0
	for i, c := range runes {
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
		return true, 0, ""
	}

	flag := runes[0 : end+1]
	rmv := flag
	flag = flag[len(mode)+2:]
	flag = flag[:len(flag)-1]
	if string(flag) == "" {
		return false, 1, string(rmv)
	}
	if strings.Contains(string(flag), ",") {
		if strings.Count(string(flag), ",") != 1 {
			return true, 0, ""
		}
		flag = []rune(strings.TrimLeft(string(flag), ","))
		flag = []rune(strings.Replace(string(flag), " ", "", -1))
		rep, err := strconv.ParseInt(string(flag), 10, 0)
		if string(flag) == "" || err != nil || rep <= 0 {
			return true, 0, ""
		}
		return false, int(rep), string(rmv)
	}
	return true, 0, ""
}

func ValideFlagIndex(line string, subster string) int {
	var ret string

	index := pkg.RuneIndex(line, subster)
	for index != -1 {
		err, _, _ := parsFlag(line, subster[2:], index)
		if err {
			ret = ret + line[0:index+3]
			line = line[index+3:]
		} else {
			return index + len(ret)
		}
		index = pkg.RuneIndex(line, subster)
	}
	return -1
}

func orderReplace(line string) string {
	for {
		up := ValideFlagIndex(line, " (up")
		low := ValideFlagIndex(line, " (low")
		cap := ValideFlagIndex(line, " (cap")
		bin := pkg.RuneIndex(line, " (bin)")
		hex := pkg.RuneIndex(line, " (hex)")
		smallest := -1
		flag := ""

		if up != -1 && (smallest == -1 || up < smallest) {
			smallest = up
			flag = "up"
		}
		if low != -1 && (smallest == -1 || low < smallest) {
			smallest = low
			flag = "low"
		}
		if cap != -1 && (smallest == -1 || cap < smallest) {
			smallest = cap
			flag = "cap"
		}
		if bin != -1 && (smallest == -1 || bin < smallest) {
			smallest = bin
			flag = "bin"
		}
		if hex != -1 && (smallest == -1 || hex < smallest) {
			smallest = hex
			flag = "hex"
		}
		if smallest == -1 {
			break
		}
		switch flag {
		case "up":
			line = flagHandler(line, smallest, "up", strings.ToUpper)
		case "low":
			line = flagHandler(line, smallest, "low", strings.ToLower)
		case "cap":
			line = flagHandler(line, smallest, "cap", pkg.CapWord)
		case "bin":
			line = binHexHandler(line, " (bin)", 2, bin)
		case "hex":
			line = binHexHandler(line, " (hex)", 16, hex)
		}
	}
	return line
}
