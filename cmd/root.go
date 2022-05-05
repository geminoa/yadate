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

			if dateOpt, err := cmd.Flags().GetString("date"); err == nil {
				resTime = modDate(resTime, dateOpt)
			}

			if utcOpt, err := cmd.Flags().GetBool("utc"); err == nil {
				if utcOpt {
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

func modDate(t time.Time, dOpt string) time.Time {
	var (
		year  int           = 0
		month int           = 0
		day   int           = 0
		hour  time.Duration = 0
		min   time.Duration = 0
		sec   time.Duration = 0
	)

	// TODO use regexp to support more flexible format
	if dOpt == "yesterday" {
		day -= 1
	} else if dOpt == "tomorrow" {
		day += 1
	} else if dOpt == "1 year ago" {
		year -= 1
	}

	t = t.AddDate(year, month, day)
	t = t.Add(time.Hour*hour + time.Minute*min + time.Second*sec)
	return t
}
