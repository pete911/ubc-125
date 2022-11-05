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
	systemBacklight = &cobra.Command{
		Use:   "backlight",
		Short: "backlight",
		RunE:  systemBacklightCmdRunE,
	}
)

func init() {
	systemCmd.AddCommand(systemVolume)
	systemCmd.AddCommand(systemSquelch)
	systemCmd.AddCommand(systemContrast)
	systemCmd.AddCommand(systemWeather)
	systemCmd.AddCommand(systemBacklight)
}

func systemVolumeCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "VOL"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	volume, err := SelectNum("select volume", 0, 15, def)
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
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	squelch, err := SelectNum("select squelch (0 - open, 15 - close)", 0, 15, def)
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
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	contrast, err := SelectNum("select lcd contrast", 1, 15, def)
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
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	options := Options{"0": "off", "1": "on"}
	value, err := Select("select weather alert", options.Values(), options.Value(def))
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, options.Key(value))
	return err
}

func systemBacklightCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal()
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "BLT"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}

	options := Options{"AO": "Always On", "AF": "Always Off", "KY": "Key Press", "SQ": "Squelch", "KS": "Key+SQL"}
	value, err := Select("select backlight", options.Values(), options.Value(def))
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, options.Key(value))
	return err
}
