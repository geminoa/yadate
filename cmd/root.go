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
			var resTime time.Time

			resTime = time.Now()

			if bUtc, err := cmd.Flags().GetBool("utc"); err == nil {
				if bUtc {
					a, err := time.LoadLocation("UTC")
					if err != nil {
						panic(err)
					}
					resTime = resTime.In(a)
				}
			}

			printDatenize(resTime)
		},
	}
)

func weekToChineseChar(key time.Weekday) string {
	var wtc = []string{
		"日", "月", "火", "水", "木", "金", "土",
	}
	return wtc[key]
}

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

func printDatenize(d time.Time) {
	fmt.Printf("%s %2d %2d %02d:%02d:%02d UTC %d\n",
		weekToChineseChar(d.Weekday()), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(), d.Year())
}
