package main

import (
	"fmt"
	"strings"
)

var metricTypes = map[string]string{
	"iperf_tests_started":         "counter",
	"iperf_tests_finished":        "counter",
	"iperf_tests_failed":          "counter",
	"iperf_sent_bytes":            "gauge",
	"iperf_sent_packets":          "gauge",
	"iperf_sent_lost_packets":     "gauge",
	"iperf_sent_seconds":          "gauge",
	"iperf_received_bytes":        "gauge",
	"iperf_received_packets":      "gauge",
	"iperf_received_lost_packets": "gauge",
	"iperf_received_seconds":      "gauge",
}

var metricHelps = map[string]string{
	"iperf_tests_started":         "Number of tests that been started.",
	"iperf_tests_finished":        "Number of tests that have finished successfully.",
	"iperf_tests_failed":          "Number of tests that have failed.",
	"iperf_sent_bytes":            "Number of bytes sent during the test.",
	"iperf_sent_packets":          "Number of packets sent during the test (UDP only).",
	"iperf_sent_lost_packets":     "Number of packet lost during the test (UDP only).",
	"iperf_sent_seconds":          "Duration of the test on the sending side, in seconds.",
	"iperf_received_bytes":        "Number of bytes received during the test.",
	"iperf_received_packets":      "Number of packets received during the test (UDP only).",
	"iperf_received_lost_packets": "Number of packets lost by the receiver during the test (UDP only).",
	"iperf_received_seconds":      "Duration of the test on the receiving side, in seconds.",
}

type Metric struct {
	Label string
	Tags  map[string]string
	Value float32
}

func (m *Metric) Format() string {
	typeStr := ""
	metricType, ok := metricTypes[m.Label]
	if ok {
		typeStr = fmt.Sprintf("# TYPE %s %s\n", m.Label, metricType)
	}

	helpStr := ""
	metricHelp, ok := metricHelps[m.Label]
	if ok {
		helpStr = fmt.Sprintf("# HELP %s %s\n", m.Label, metricHelp)
	}

	tags := make([]string, 0)
	for key, value := range m.Tags {
		tags = append(tags, fmt.Sprintf("%s=\"%s\"", key, value))
	}

	tagStr := ""
	if len(tags) > 0 {
		tagStr = "{" + strings.Join(tags, ",") + "}"
	}

	return fmt.Sprintf("%s%s%s%s %f\n", typeStr, helpStr, m.Label, tagStr, m.Value)
}
