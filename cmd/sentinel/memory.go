package main

import (
	"encoding/json"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

var memoryCsvHeader = "time,used,free,usedPercent\n"

type MemUsage struct {
	Time        string  `json:"time"`
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        uint64  `json:"free"`
}

func getMemUsage(csv bool) (string, error) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Failed to get memory usage:", err)
		return "", err
	}
	queryTimeInUnixString := getUnixTimeInNanoUTC()
	usages := MemUsage{
		Time:        queryTimeInUnixString,
		Total:       memory.Total,
		Available:   memory.Available,
		Used:        memory.Used,
		UsedPercent: memory.UsedPercent,
		Free:        memory.Free,
	}
	jsonData, err := json.MarshalIndent(usages, "", "    ")
	if err != nil {
		return "", err
	}
	if csv {
		return fmt.Sprintf("%s,%d,%d,%.2f\n", queryTimeInUnixString, memory.Used, memory.Free, memory.UsedPercent), nil
	}

	return string(jsonData), nil

}
