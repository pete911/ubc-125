package cmd

import (
	"github.com/pete911/ubc-125/prompt"
	"github.com/spf13/cobra"
	"strings"
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
	systemBattery = &cobra.Command{
		Use:   "battery",
		Short: "battery charge time",
		RunE:  systemBatteryCmdRunE,
	}
	systemMemory = &cobra.Command{
		Use:   "memory",
		Short: "clear all memory",
		RunE:  systemMemoryCmdRunE,
	}
	systemKey = &cobra.Command{
		Use:   "key",
		Short: "key beep and key lock",
		RunE:  systemKeyCmdRunE,
	}
	systemPriority = &cobra.Command{
		Use:   "priority",
		Short: "priority mode",
		RunE:  systemPriorityCmdRunE,
	}
)

func init() {
	systemCmd.AddCommand(systemVolume)
	systemCmd.AddCommand(systemSquelch)
	systemCmd.AddCommand(systemContrast)
	systemCmd.AddCommand(systemWeather)
	systemCmd.AddCommand(systemBacklight)
	systemCmd.AddCommand(systemBattery)
	systemCmd.AddCommand(systemMemory)
	systemCmd.AddCommand(systemKey)
	systemCmd.AddCommand(systemPriority)
}

func systemVolumeCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(false)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "VOL"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	volume, err := prompt.SelectNum("select volume", 0, 15, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, volume)
	return err
}

func systemSquelchCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(false)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "SQL"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	squelch, err := prompt.SelectNum("select squelch (0 - open, 15 - close)", 0, 15, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, squelch)
	return err
}

func systemContrastCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "CNT"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	contrast, err := prompt.SelectNum("select lcd contrast", 1, 15, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, contrast)
	return err
}

func systemWeatherAlertCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "WXS"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	options := prompt.Options{"0": "off", "1": "on"}
	weatherAlert, err := prompt.SelectOptions("select weather alert", options, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, weatherAlert)
	return err
}

func systemBacklightCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "BLT"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}

	options := prompt.Options{"AO": "Always On", "AF": "Always Off", "KY": "Key Press", "SQ": "Squelch", "KS": "Key+SQL"}
	backlight, err := prompt.SelectOptions("select backlight", options, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, backlight)
	return err
}

func systemBatteryCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "BSV"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	chargeTime, err := prompt.SelectNum("select battery charge time", 1, 16, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, chargeTime)
	return err
}

func systemMemoryCmdRunE(_ *cobra.Command, _ []string) error {
	ok, err := prompt.Confirm("do you want to clear all memory?")
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	_, err = term.WriteE("CLR")
	return err
}

func systemKeyCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "KBP"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}
	var defLevel, defLock string
	if parts := strings.Split(def, ","); len(parts) == 2 {
		defLevel = parts[0]
		defLock = parts[1]
	}

	levelOptions := prompt.Options{"0": "Auto", "99": "Off"}
	level, err := prompt.SelectOptions("select key beep", levelOptions, defLevel)
	if err != nil {
		return err
	}

	lockOptions := prompt.Options{"0": "Off", "1": "On"}
	lock, err := prompt.SelectOptions("select key lock", lockOptions, defLock)
	if err != nil {
		return err
	}

	_, err = term.WriteE(cmd, level, lock)
	return err
}

func systemPriorityCmdRunE(_ *cobra.Command, _ []string) error {
	term, err := Terminal(true)
	if err != nil {
		return err
	}
	defer term.Close()

	cmd := "PRI"
	def, err := term.WriteE(cmd)
	if err != nil {
		return err
	}

	options := prompt.Options{"0": "Off", "1": "On", "2": "Plus On", "3": "DND"}
	priority, err := prompt.SelectOptions("select priority mode", options, def)
	if err != nil {
		return err
	}
	_, err = term.WriteE(cmd, priority)
	return err
}
