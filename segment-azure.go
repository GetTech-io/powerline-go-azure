package main

import (
	"os/exec"
	"strings"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentAZURE() []pwl.Segment {
	commandOutput, err := exec.Command("az", "account", "show", "--query", "name", "--output", "tsv").Output()
	if err != nil || len(commandOutput) == 0 {
		return []pwl.Segment{}
	}
	subscriptionName := strings.TrimSpace(string(commandOutput))
	if len(subscriptionName) == 0 {
		return []pwl.Segment{}
	}
	return []pwl.Segment{{
		Name:       "azure",
		Content:    subscriptionName,
		Foreground: 15,
		Background: 26,
	}}
}
