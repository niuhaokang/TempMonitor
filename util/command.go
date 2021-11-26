package util

import (
	"log"
	"strconv"
	"strings"
)

type TempInfo struct {
	Label string
	Temp int
}

func ParseCMD(response string) []*TempInfo {
	tempInfo := make([]*TempInfo, 0)
	lines := strings.Split(response, "\n")
	for _, line := range lines[1 : len(lines)-1] {
		info := strings.Split(line, "|")
		label := info[0]
		temp := func(t string) int {
			t = strings.Split(t, ",")[0]
			t = strings.Split(t, "'")[0]
			t = strings.TrimSpace(t)
			tempInt, err :=strconv.Atoi(t)
			if err != nil {
				log.Fatalf("err : %v\n", err)
			}
			return tempInt
		}(info[1])
		tempInfo = append(tempInfo, &TempInfo{
			Label: label,
			Temp: temp,
		})
	}
	return tempInfo
}
