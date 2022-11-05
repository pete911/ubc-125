package cmd

import (
	"github.com/spf13/cobra"
)

var (
	systemCmd = &cobra.Command{
		Use:   "system",
		Short: "system settings",
	}
	systemVolume = &cobra.Command{
		Use:   "volume",
		Short: "volume level",
		Run:   systemVolumeCmdRun,
	}
	systemSquelch = &cobra.Command{
		Use:   "squelch",
		Short: "squelch level",
		Run:   systemSquelchCmdRun,
	}
)

func init() {
	systemCmd.AddCommand(systemVolume)
	systemCmd.AddCommand(systemSquelch)
}

func systemVolumeCmdRun(_ *cobra.Command, _ []string) {
	term := Terminal()
	defer term.Close()

	cmd := "VOL"
	defaultVolume := term.Write(cmd)
	volume := SelectNum("select volume", 0, 15, defaultVolume)
	term.Write(cmd, volume)
}

func systemSquelchCmdRun(_ *cobra.Command, _ []string) {
	term := Terminal()
	defer term.Close()

	cmd := "SQL"
	defaultSquelch := term.Write(cmd)
	squelch := SelectNum("select squelch (0 - open, 15 - close)", 0, 15, defaultSquelch)
	term.Write(cmd, squelch)
}
