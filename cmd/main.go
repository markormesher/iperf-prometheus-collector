package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	testsStarted  = Metric{Label: "iperf_tests_started", Value: 0}
	testsFinished = Metric{Label: "iperf_tests_finished", Value: 0}
	testsFailed   = Metric{Label: "iperf_tests_failed", Value: 0}
)

var liveMetrics Queue[Metric]

func main() {
	settings, err := getSettings()
	if err != nil {
		l.Error("Failed to get settings")
		panic(err)
	}

	go runTests(settings)

	http.HandleFunc("/", httpHandler)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", settings.ListenPort), nil)
	if err != nil {
		l.Error("Failed to start server")
		panic(err)
	}
}

func runTests(settings *Settings) {
	for {
		runIperfTest(settings)
		time.Sleep(time.Duration(settings.TestIntervalMs) * time.Millisecond)
	}
}

func httpHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")

	// fixed metrics
	fmt.Fprintf(res, testsStarted.Format())
	fmt.Fprintf(res, testsFinished.Format())
	fmt.Fprintf(res, testsFailed.Format())

	// buffered metrics
	for {
		metric, ok := liveMetrics.Pop()
		if !ok {
			break
		}

		fmt.Fprintf(res, metric.Format())
	}
}
