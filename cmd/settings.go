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
	Protocol         string
}

func getSettings() (*Settings, error) {
	targetListRaw := os.Getenv("TARGET_LIST")
	targetList := strings.Split(targetListRaw, ",")

	updateIntervalMsStr := os.Getenv("UPDATE_INTERVAL_MS")
	if updateIntervalMsStr == "" {
		updateIntervalMsStr = "600000"
	}
	updateIntervalMs, err := strconv.Atoi(updateIntervalMsStr)
	if err != nil {
		return nil, fmt.Errorf("Could not parse update interval as an integer: %w", err)
	}

	protocol := os.Getenv("PROTOCOL")
	if protocol == "" {
		protocol = "tcp"
	}
	if protocol != "tcp" && protocol != "udp" {
		return nil, fmt.Errorf("Invalid protocol: %s", protocol)
	}

	return &Settings{
		TargetList:       targetList,
		UpdateIntervalMs: updateIntervalMs,
		Protocol:         protocol,
	}, nil
}
