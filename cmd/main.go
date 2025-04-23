package main

import (
	"fmt"
	"net/http"
	"time"
)

var metrics = NewMetricsHandler()
var settings Settings

func main() {
	var err error
	err = loadSettings()
	if err != nil {
		l.Error("Failed to load settings")
		panic(err)
	}

	metrics.Submit(Metric{Label: "iperf_tests_started", Value: 0})
	metrics.Submit(Metric{Label: "iperf_tests_finished", Value: 0})
	metrics.Submit(Metric{Label: "iperf_tests_failed", Value: 0})

	go runTests()

	http.HandleFunc("/", httpHandler)
	err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", settings.ListenPort), nil)
	if err != nil {
		l.Error("Failed to start server")
		panic(err)
	}
}

func runTests() {
	for ; true; <-time.Tick(time.Duration(settings.TestIntervalMs) * time.Millisecond) {
		runIperfTest()
	}
}

func httpHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/plain")

	for _, m := range metrics.GetAll() {
		_, err := res.Write([]byte(m.Format()))
		if err != nil {
			l.Error("error writing HTTP response", "error", err)
			return
		}
	}
}
