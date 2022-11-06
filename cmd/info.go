package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	infoCmd = &cobra.Command{
		Use:   "info",
		Short: "system information",
		RunE:  infoCmdRunE,
	}
)

func infoCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(false)
	if err != nil {
		return err
	}
	defer term.Close()

	model, err := term.WriteE("MDL")
	if err != nil {
		return err
	}
	version, err := term.WriteE("VER")
	if err != nil {
		return err
	}

	fmt.Printf("Model: %s\nVersion: %s\n", model, version)
	return nil
}
