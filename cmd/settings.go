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
	Options          string
	ListenPort       int
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

	protocol := os.Getenv("TEST_PROTOCOL")
	if protocol == "" {
		protocol = "tcp"
	}
	if protocol != "tcp" && protocol != "udp" {
		return nil, fmt.Errorf("Invalid protocol: %s", protocol)
	}

	options := strings.TrimSpace(os.Getenv("TEST_OPTIONS"))

	listenPortStr := os.Getenv("LISTEN_PORT")
	if listenPortStr == "" {
		listenPortStr = "9030"
	}
	listenPort, err := strconv.Atoi(listenPortStr)
	if err != nil {
		return nil, fmt.Errorf("Could not parse listen port as an integer: %w", err)
	}

	return &Settings{
		TargetList:       targetList,
		UpdateIntervalMs: updateIntervalMs,
		Protocol:         protocol,
		Options:          options,
		ListenPort:       listenPort,
	}, nil
}
