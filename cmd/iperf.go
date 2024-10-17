package main

type IperfTcpResult struct {
	Error string `json:"error"`
	End   struct {
		SumSent struct {
			Bytes   int `json:"bytes"`
			Seconds int `json:"seconds"`
		} `json:"sum_sent"`
		SumReceived struct {
			Bytes   int `json:"bytes"`
			Seconds int `json:"seconds"`
		} `json:"sum_received"`
	} `json:"end"`
}

func runIperfTest(settings *Settings) {

}
