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

func runIperfTest(settings *Settings) {
	for _, target := range settings.TargetList {
		testsStarted.Value++

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
			testsFailed.Value++
			continue
		}

		var result IperfResult
		err = json.NewDecoder(bytes.NewReader(output)).Decode(&result)
		if err != nil {
			l.Error("Failed to decode iperf", "target", target, "error", err)
			testsFailed.Value++
			continue
		}

		liveMetrics.Push(Metric{Label: "iperf_sent_bytes", Tags: tags, Value: result.End.SumSent.Bytes})
		liveMetrics.Push(Metric{Label: "iperf_sent_seconds", Tags: tags, Value: result.End.SumSent.Seconds})
		if settings.Protocol == "udp" {
			liveMetrics.Push(Metric{Label: "iperf_sent_packets", Tags: tags, Value: result.End.SumSent.Packets})
			liveMetrics.Push(Metric{Label: "iperf_sent_lost_packets", Tags: tags, Value: result.End.SumSent.LostPackets})
		}

		liveMetrics.Push(Metric{Label: "iperf_received_bytes", Tags: tags, Value: result.End.SumReceived.Bytes})
		liveMetrics.Push(Metric{Label: "iperf_received_seconds", Tags: tags, Value: result.End.SumReceived.Seconds})
		if settings.Protocol == "udp" {
			liveMetrics.Push(Metric{Label: "iperf_received_packets", Tags: tags, Value: result.End.SumReceived.Packets})
			liveMetrics.Push(Metric{Label: "iperf_received_lost_packets", Tags: tags, Value: result.End.SumReceived.LostPackets})
		}

		testsFinished.Value++
	}
}
