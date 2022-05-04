package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "yadate",
		Short: "Yet another date command",
		Long: `Yet another date command for providing more flexible ways for ` +
			`standard date command.`,
		Run: func(cmd *cobra.Command, args []string) {
			if bUtc, err := cmd.Flags().GetBool("utc"); err == nil {
				if bUtc {
					tz := "America/New_York"
					a, err := time.LoadLocation(tz)
					if err != nil {
						panic(err)
					}
					now := time.Now().In(a)
					fmt.Println(now)
				} else {
					now := time.Now()
					fmt.Println(now)
				}
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP(
		"utc", "u", false,
		"print or set Coordinated Universal Time (UTC)")
	rootCmd.Flags().StringP(
		"date", "d", "",
		"display time described by STRING, not 'now'")
}

func initConfig() {
	// init config with viper
	// https://github.com/spf13/cobra/blob/master/user_guide.md
}
