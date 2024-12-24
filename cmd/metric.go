package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type MetricsHandler struct {
	latestMetrics map[string]Metric
	totalMetrics  map[string]Metric
}

func NewMetricsHandler() MetricsHandler {
	return MetricsHandler{
		latestMetrics: map[string]Metric{},
		totalMetrics:  map[string]Metric{},
	}
}

func (h *MetricsHandler) Submit(m Metric) {
	key := m.StableKey()

	// always update the totals metrics
	oldTotalMetric, ok := h.totalMetrics[key]
	if !ok {
		oldTotalMetric = Metric{
			Label: m.Label + "_total",
			Tags:  m.Tags,
			Value: m.Value,
		}
	} else {
		oldTotalMetric.Value += m.Value
	}
	h.totalMetrics[key] = oldTotalMetric

	// only update latest metrics for relevant labels
	if !strings.HasPrefix(m.Label, "iperf_tests_") {
		h.latestMetrics[key] = m
	}
}

func (h *MetricsHandler) GetAll() []Metric {
	output := make([]Metric, 0)

	for k, m := range h.latestMetrics {
		output = append(output, m)
		if settings.DisardStaleResults {
			delete(h.latestMetrics, k)
		}
	}

	for _, m := range h.totalMetrics {
		output = append(output, m)
	}

	return output
}

type Metric struct {
	Label string
	Tags  map[string]string
	Value float32
}

func (m *Metric) StableKey() string {
	key := m.Label

	tagKeys := slices.Collect(maps.Keys(m.Tags))
	slices.Sort(tagKeys)
	for _, tagKey := range tagKeys {
		key = key + fmt.Sprintf("/%s=%s", tagKey, m.Tags[tagKey])
	}

	return key
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

var metricTypes = map[string]string{
	// running stats
	"iperf_tests_started_total":  "counter",
	"iperf_tests_finished_total": "counter",
	"iperf_tests_failed_total":   "counter",

	// latest test results
	"iperf_sent_bytes":            "gauge",
	"iperf_sent_packets":          "gauge",
	"iperf_sent_lost_packets":     "gauge",
	"iperf_sent_seconds":          "gauge",
	"iperf_received_bytes":        "gauge",
	"iperf_received_packets":      "gauge",
	"iperf_received_lost_packets": "gauge",
	"iperf_received_seconds":      "gauge",

	// summed test results
	"iperf_sent_bytes_total":            "gauge",
	"iperf_sent_packets_total":          "gauge",
	"iperf_sent_lost_packets_total":     "gauge",
	"iperf_sent_seconds_total":          "gauge",
	"iperf_received_bytes_total":        "gauge",
	"iperf_received_packets_total":      "gauge",
	"iperf_received_lost_packets_total": "gauge",
	"iperf_received_seconds_total":      "gauge",
}

var metricHelps = map[string]string{
	// running stats
	"iperf_tests_started_total":  "Number of tests that been started.",
	"iperf_tests_finished_total": "Number of tests that have finished successfully.",
	"iperf_tests_failed_total":   "Number of tests that have failed.",

	// latest test results
	"iperf_sent_bytes":            "Number of bytes sent during the latest test.",
	"iperf_sent_packets":          "Number of packets sent during the latest test (UDP only).",
	"iperf_sent_lost_packets":     "Number of packet lost during the latest test (UDP only).",
	"iperf_sent_seconds":          "Duration of the latest test on the sending side, in seconds.",
	"iperf_received_bytes":        "Number of bytes received during the latest test.",
	"iperf_received_packets":      "Number of packets received during the latest test (UDP only).",
	"iperf_received_lost_packets": "Number of packets lost by the receiver during the latest test (UDP only).",
	"iperf_received_seconds":      "Duration of the latest test on the receiving side, in seconds.",

	// summed test results
	"iperf_sent_bytes_total":            "Number of bytes sent during all completed tests.",
	"iperf_sent_packets_total":          "Number of packets sent during all completed tests (UDP only).",
	"iperf_sent_lost_packets_total":     "Number of packet lost during all completed tests (UDP only).",
	"iperf_sent_seconds_total":          "Duration of all completed tests on the sending side, in seconds.",
	"iperf_received_bytes_total":        "Number of bytes received during all completed tests.",
	"iperf_received_packets_total":      "Number of packets received during all completed tests (UDP only).",
	"iperf_received_lost_packets_total": "Number of packets lost by the receiver during all completed tests (UDP only).",
	"iperf_received_seconds_total":      "Duration of all completed tests on the receiving side, in seconds.",
}
