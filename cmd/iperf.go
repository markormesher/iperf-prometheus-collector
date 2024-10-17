package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type IperfTcpResult struct {
	End struct {
		SumSent struct {
			Bytes   float32 `json:"bytes"`
			Seconds float32 `json:"seconds"`
		} `json:"sum_sent"`
		SumReceived struct {
			Bytes   float32 `json:"bytes"`
			Seconds float32 `json:"seconds"`
		} `json:"sum_received"`
	} `json:"end"`
}

func runIperfTest(settings *Settings) {
	for _, target := range settings.TargetList {
		testsStarted.Value++

		tags := map[string]string{
			"target": target,
		}

		cmd := exec.Command("bash", "-c", fmt.Sprintf("iperf3 -c %s --json", target))

		output, err := cmd.Output()
		if err != nil {
			l.Error("Failed to run iperf test", "target", target, "error", err, "output", string(output))
			testsFailed.Value++
			continue
		}

		var result IperfTcpResult
		err = json.NewDecoder(bytes.NewReader(output)).Decode(&result)
		if err != nil {
			l.Error("Failed to decode iperf", "target", target, "error", err)
			testsFailed.Value++
			continue
		}

		liveMetrics.Push(Metric{Label: "iperf_sent_bytes", Tags: tags, Value: result.End.SumSent.Bytes})
		liveMetrics.Push(Metric{Label: "iperf_sent_seconds", Tags: tags, Value: result.End.SumSent.Seconds})
		liveMetrics.Push(Metric{Label: "iperf_received_bytes", Tags: tags, Value: result.End.SumReceived.Bytes})
		liveMetrics.Push(Metric{Label: "iperf_received_seconds", Tags: tags, Value: result.End.SumReceived.Seconds})
		testsFinished.Value++
	}
}
