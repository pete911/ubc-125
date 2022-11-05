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
		RunE:  systemVolumeCmdRunE,
	}
	systemSquelch = &cobra.Command{
		Use:   "squelch",
		Short: "squelch level",
		RunE:  systemSquelchCmdRunE,
	}
	systemContrast = &cobra.Command{
		Use:   "contrast",
		Short: "lcd contrast level",
		RunE:  systemContrastCmdRunE,
	}
	systemWeather = &cobra.Command{
		Use:   "weather",
		Short: "weather priority",
		RunE:  systemWeatherAlertCmdRunE,
	}
)

func init() {
	systemCmd.AddCommand(systemVolume)
	systemCmd.AddCommand(systemSquelch)
	systemCmd.AddCommand(systemContrast)
	systemCmd.AddCommand(systemWeather)
}

func systemVolumeCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "VOL"
	defaultVolume, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	volume, err := SelectNum("select volume", 0, 15, defaultVolume)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, volume)
	return err
}

func systemSquelchCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "SQL"
	defaultSquelch, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	squelch, err := SelectNum("select squelch (0 - open, 15 - close)", 0, 15, defaultSquelch)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, squelch)
	return err
}

func systemContrastCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "CNT"
	defaultContrast, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	contrast, err := SelectNum("select lcd contrast", 1, 15, defaultContrast)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, contrast)
	return err
}

func systemWeatherAlertCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "WXS"
	defaultAlert, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	alert, err := SelectNum("select weather alert (0 - off, 1 - on)", 0, 1, defaultAlert)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, alert)
	return err
}
