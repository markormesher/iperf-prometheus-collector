package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type IperfResult struct {
	End struct {
		SumSent struct {
			Bytes       float32 `json:"bytes"`
			Seconds     float32 `json:"seconds"`
			Packets     float32 `json:"packets"`
			LostPackets float32 `json:"LostPackets"`
		} `json:"sum_sent"`
		SumReceived struct {
			Bytes       float32 `json:"bytes"`
			Seconds     float32 `json:"seconds"`
			Packets     float32 `json:"packets"`
			LostPackets float32 `json:"LostPackets"`
		} `json:"sum_received"`
	} `json:"end"`
}

func runIperfTest() {
	for _, target := range settings.TargetList {
		l.Info("Running iperf test", "target", target)
		metrics.Submit(Metric{Label: "iperf_tests_started", Value: 1})

		tags := map[string]string{
			"target":   target,
			"protocol": settings.Protocol,
			"options":  settings.Options,
		}

		udpOption := ""
		if settings.Protocol == "udp" {
			udpOption = "--udp"
		}

		cmd := exec.Command("bash", "-c", fmt.Sprintf("iperf3 --client %s --json %s %s", target, udpOption, settings.Options))

		output, err := cmd.Output()
		if err != nil {
			l.Error("Failed to run iperf test", "target", target, "error", err, "output", string(output))
			metrics.Submit(Metric{Label: "iperf_tests_failed", Value: 1})
			continue
		}

		var result IperfResult
		err = json.NewDecoder(bytes.NewReader(output)).Decode(&result)
		if err != nil {
			l.Error("Failed to decode iperf result", "target", target, "error", err)
			metrics.Submit(Metric{Label: "iperf_tests_failed", Value: 1})
			continue
		}

		metrics.Submit(Metric{Label: "iperf_sent_bytes", Tags: tags, Value: result.End.SumSent.Bytes})
		metrics.Submit(Metric{Label: "iperf_sent_seconds", Tags: tags, Value: result.End.SumSent.Seconds})
		if settings.Protocol == "udp" {
			metrics.Submit(Metric{Label: "iperf_sent_packets", Tags: tags, Value: result.End.SumSent.Packets})
			metrics.Submit(Metric{Label: "iperf_sent_lost_packets", Tags: tags, Value: result.End.SumSent.LostPackets})
		}

		metrics.Submit(Metric{Label: "iperf_received_bytes", Tags: tags, Value: result.End.SumReceived.Bytes})
		metrics.Submit(Metric{Label: "iperf_received_seconds", Tags: tags, Value: result.End.SumReceived.Seconds})
		if settings.Protocol == "udp" {
			metrics.Submit(Metric{Label: "iperf_received_packets", Tags: tags, Value: result.End.SumReceived.Packets})
			metrics.Submit(Metric{Label: "iperf_received_lost_packets", Tags: tags, Value: result.End.SumReceived.LostPackets})
		}

		l.Info("Finished iperf test", "target", target)
		metrics.Submit(Metric{Label: "iperf_tests_finished", Value: 1})
	}
}
