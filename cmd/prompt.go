package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func SelectNum(msg string, min, max int, def string) (string, error) {
	prompt := &survey.Select{
		Message: msg,
		Options: getNumSlice(min, max),
		Default: def,
	}
	var out int
	if err := survey.AskOne(prompt, &out); err != nil {
		return "", fmt.Errorf("error: %s: %v", msg, err)
	}
	return fmt.Sprintf("%d", out), nil
}

func Select(msg string, options []string, def string) (string, error) {
	if def == "" {
		def = options[0]
	}
	prompt := &survey.Select{
		Message: msg,
		Options: options,
		Default: def,
	}
	var out string
	if err := survey.AskOne(prompt, &out); err != nil {
		return "", fmt.Errorf("error: %s: %v", msg, err)
	}
	return out, nil
}

func getNumSlice(min, max int) []string {
	a := make([]string, max-min+1)
	for i := range a {
		a[i] = fmt.Sprintf("%d", min+i)
	}
	return a
}
