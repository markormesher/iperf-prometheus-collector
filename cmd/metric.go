package main

import "fmt"

type Metric struct {
	Label string
	Value float32
}

func (m *Metric) Format() string {
	return fmt.Sprintf("%s %f\n", m.Label, m.Value)
}
