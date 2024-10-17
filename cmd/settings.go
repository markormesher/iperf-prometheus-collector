package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	TargetList       []string
	UpdateIntervalMs int
}

func getSettings() (*Settings, error) {
	targetListRaw := os.Getenv("TARGET_LIST")
	targetList := strings.Split(targetListRaw, ",")

	updateIntervalMsStr := os.Getenv("UPDATE_INTERVAL_MS")
	if len(updateIntervalMsStr) == 0 {
		updateIntervalMsStr = "600000"
	}
	updateIntervalMs, err := strconv.Atoi(updateIntervalMsStr)
	if err != nil {
		return nil, fmt.Errorf("Could not parse update interval as an integer: %w", err)
	}

	return &Settings{
		TargetList:       targetList,
		UpdateIntervalMs: updateIntervalMs,
	}, nil
}
