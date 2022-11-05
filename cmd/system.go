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
	systemContrast = &cobra.Command{
		Use:   "contrast",
		Short: "lcd contrast level",
		Run:   systemContrastCmdRun,
	}
	systemWeather = &cobra.Command{
		Use:   "weather",
		Short: "weather priority",
		Run:   systemWeatherAlertCmdRun,
	}
)

func init() {
	systemCmd.AddCommand(systemVolume)
	systemCmd.AddCommand(systemSquelch)
	systemCmd.AddCommand(systemContrast)
	systemCmd.AddCommand(systemWeather)
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

func systemContrastCmdRun(_ *cobra.Command, _ []string) {
	term := Terminal()
	defer term.Close()

	cmd := "CNT"
	defaultContrast := term.Write(cmd)
	contrast := SelectNum("select lcd contrast", 1, 15, defaultContrast)
	term.Write(cmd, contrast)
}

func systemWeatherAlertCmdRun(_ *cobra.Command, _ []string) {
	term := Terminal()
	defer term.Close()

	cmd := "WXS"
	defaultAlert := term.Write(cmd)
	alert := SelectNum("select weather alert (0 - off, 1 - on)", 0, 1, defaultAlert)
	term.Write(cmd, alert)
}
