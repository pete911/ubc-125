package cmd

import (
	"fmt"
	"github.com/pete911/ubc-125/port"
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
	RootCmd.AddCommand(systemCmd)
}

func initLog(verbose bool) {
	logger = log.New(io.Discard, "", log.Ltime|log.Lmicroseconds|log.Lshortfile)
	if verbose {
		logger.SetOutput(os.Stderr)
	}
}

func Terminal() port.Terminal {
	return port.NewTerminal(logger, openPort())
}

func openPort() port.Port {
	if dryRun {
		return port.OpenDryRun()
	}

	if portName == "" {
		portName = selectPort()
	}
	p, err := port.Open(portName)
	if err != nil {
		Fatal(fmt.Sprintf("error: open port: %v", err))
	}
	return p
}

func selectPort() string {
	portNames, err := port.List()
	if err != nil {
		Fatal(fmt.Sprintf("error: list ports: %v", err))
	}
	return Select("select port:", portNames, "")
}

func Fatal(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
