package collection

import (
	"github.com/xh3b4sd/anna/factory/input"
	"github.com/xh3b4sd/anna/factory/output"
	"github.com/xh3b4sd/anna/spec"
)

func (c *collection) SelectFirstCharacterOfSentence() ([]spec.Challenge, error) {
	configs := []struct {
		Input   string
		Outputs []string
	}{
		{
			Input: "Select the first character of this sentence.",
			Outputs: []string{
				"S",
			},
		},
		{
			Input: "Show me the first character of this sentence.",
			Outputs: []string{
				"S",
			},
		},
		{
			Input: "Select the character on the first position of this sentence.",
			Outputs: []string{
				"S",
			},
		},
		{
			Input: "Show me the character on the first position of this sentence.",
			Outputs: []string{
				"S",
			},
		},
		{
			Input: "Select the first character of this sentence.",
			Outputs: []string{
				"S",
			},
		},
		{
			Input: "What is the first character of this sentence?",
			Outputs: []string{
				"W",
			},
		},
		{
			Input: "What is the first character of the first word of this sentence?",
			Outputs: []string{
				"W",
			},
		},
	}

	var newChallenges []spec.Challenge

	for _, c := range configs {
		newConfig := DefaultConfig()
		newConfig.Input = c.Input
		newConfig.Outputs = c.Outputs
		newInputsOutputs, err := input.New(newConfig)
		if err != nil {
			return nil, maskAny(err)
		}
	}

	return newChallenges, nil
}
