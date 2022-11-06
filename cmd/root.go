package cmd

import (
	"fmt"
	"github.com/pete911/ubc-125/port"
	"github.com/pete911/ubc-125/prompt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
)

var (
	dryRun   bool
	verbose  bool
	portName string
	logger   *log.Logger

	RootCmd = &cobra.Command{
		Use:   "ubc-125",
		Short: "Uniden UBC125 Programming",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initLog(verbose)
		},
	}
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "dry run, do not connect")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose, print logs")
	RootCmd.PersistentFlags().StringVarP(&portName, "port", "p", "", "serial port name")

	RootCmd.AddCommand(infoCmd)
	RootCmd.AddCommand(systemCmd)
}

func initLog(verbose bool) {
	logger = log.New(io.Discard, "", log.Ltime|log.Lmicroseconds|log.Lshortfile)
	if verbose {
		logger.SetOutput(os.Stderr)
	}
}

func Terminal(programMode bool) (port.Terminal, error) {
	p, err := openPort()
	if err != nil {
		return port.Terminal{}, err
	}
	return port.NewTerminal(logger, p, programMode)
}

func openPort() (port.Port, error) {
	if dryRun {
		return port.OpenDryRun(), nil
	}

	if portName == "" {
		selectedPort, err := selectPort()
		if err != nil {
			return nil, err
		}
		portName = selectedPort
	}
	p, err := port.Open(portName)
	if err != nil {
		return nil, fmt.Errorf("error: open port: %v", err)
	}
	return p, nil
}

func selectPort() (string, error) {
	portNames, err := port.List()
	if err != nil {
		return "", fmt.Errorf("error: list ports: %v", err)
	}
	return prompt.Select("select port:", portNames, "")
}
