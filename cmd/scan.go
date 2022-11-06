package cmd

import (
	"github.com/pete911/ubc-125/prompt"
	"github.com/spf13/cobra"
	"strings"
)

var (
	scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "scan settings",
	}
	scanGroup = &cobra.Command{
		Use:   "group",
		Short: "scan channel group",
		RunE:  scanGroupCmdRunE,
	}
)

func init() {
	scanCmd.AddCommand(scanGroup)
}

func scanGroupCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(false)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "SCG"
	out, err := term.WriteE(cmd)
	if err != nil {
		return err
	}

	options := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	selected := getSelectedScanGroups(options, out)
	scanGroups, err := prompt.MultiSelect("select scan channel groups", options, selected)
	if err != nil {
		return err
	}
	args := toScanGroups(options, scanGroups)
	_, err = term.WriteE(cmd, args)
	return err
}

// getSelectedScanGroups returns slice of labels that are selected in 'in' input
func getSelectedScanGroups(labels []string, in string) []string {
	var selected []string
	for i, v := range []rune(in) {
		if v == '0' { // 0 - valid, 1 - invalid
			selected = append(selected, labels[i])
		}
	}
	return selected
}

// toScanGroups returns scan group string (e.g. 1011101111) from supplied labels and selected labels
func toScanGroups(labels, selected []string) string {
	selectedSet := make(map[string]struct{})
	for _, v := range selected {
		selectedSet[v] = struct{}{}
	}

	out := make([]string, len(labels))
	for i, v := range labels {
		if _, ok := selectedSet[v]; ok {
			out[i] = "0" // 0 - valid, 1 - invalid
		} else {
			out[i] = "1" // 0 - valid, 1 - invalid
		}
	}
	return strings.Join(out, "")
}
