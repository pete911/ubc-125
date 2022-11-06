package prompt

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

func Confirm(msg string) (bool, error) {
	ok := false
	err := survey.AskOne(&survey.Confirm{Message: msg}, &ok)
	return ok, err
}

func SelectNum(msg string, min, max int, def string) (string, error) {
	return Select(msg, getNumSlice(min, max), def)
}

func SelectOptions(msg string, options Options, def string) (string, error) {
	value, err := Select(msg, options.Values(), options.Value(def))
	if err != nil {
		return "", err
	}
	return options.Key(value), nil
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

func MultiSelect(msg string, options []string, def []string) ([]string, error) {
	prompt := &survey.MultiSelect{
		Message: msg,
		Options: options,
		Default: def,
	}
	var out []string
	if err := survey.AskOne(prompt, &out); err != nil {
		return nil, fmt.Errorf("error: %s: %v", msg, err)
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
