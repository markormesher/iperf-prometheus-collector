package main

import (
	"fmt"
	"strings"
)

type Metric struct {
	Label string
	Tags  map[string]string
	Value float32
}

func (m *Metric) Format() string {
	tags := make([]string, 0)
	for key, value := range m.Tags {
		tags = append(tags, fmt.Sprintf("%s=\"%s\"", key, value))
	}

	tagStr := ""
	if len(tags) > 0 {
		tagStr = "{" + strings.Join(tags, ",") + "}"
	}

	return fmt.Sprintf("%s%s %f\n", m.Label, tagStr, m.Value)
}
