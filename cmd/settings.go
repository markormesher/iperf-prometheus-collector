package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	TargetList     []string
	TestIntervalMs int
	Protocol       string
	Options        string
	ListenPort     int
}

func getSettings() (*Settings, error) {
	targetListRaw := os.Getenv("TARGET_LIST")
	targetList := strings.Split(targetListRaw, ",")

	testIntervalMsStr := os.Getenv("TEST_INTERVAL_MS")
	if testIntervalMsStr == "" {
		testIntervalMsStr = "600000"
	}
	updateIntervalMs, err := strconv.Atoi(testIntervalMsStr)
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
		TargetList:     targetList,
		TestIntervalMs: updateIntervalMs,
		Protocol:       protocol,
		Options:        options,
		ListenPort:     listenPort,
	}, nil
}
