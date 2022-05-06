package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type DateTime struct {
	Year, Month, Day int
	Hour, Min, Sec   time.Duration
}

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

func modDate(t time.Time, dOpt string) time.Time {
	var (
		dt = DateTime{0, 0, 0, 0, 0, 0}
	)

	reSpaces := regexp.MustCompile(`\s+`)
	dOpt = reSpaces.ReplaceAllString(dOpt, " ")
	dOptTerms := strings.Split(dOpt, " ")

	// TODO use regexp to support more flexible format
	if len(dOptTerms) == 1 {
		n, term := parseSingleDateOpt(dOpt)
		switch term {
		case "yesterday":
			dt.Day -= n * 1
		case "tomorrow":
			dt.Day += n * 1
		case "week":
			dt.Day += n * 7
		case "fortnight":
			dt.Day += n * 14
		case "day":
			dt.Month += n * 1
		case "month":
			dt.Month += n * 1
		case "year":
			dt.Year += n * 1
		default:
		}
	} else {
		// TODO
		r := regexp.MustCompile(`(\d+)\s+(\w+)\s*(\w*)`)
		if r.MatchString(dOpt) {
			fmt.Println(dOpt)
			a := r.FindAllSubmatch([]byte(dOpt), -1)
			fmt.Println(a[0])
			dt.Year -= 1
		}
	}

	t = t.AddDate(dt.Year, dt.Month, dt.Day)
	t = t.Add(time.Hour*dt.Hour + time.Minute*dt.Min + time.Second*dt.Sec)
	return t
}

func parseSingleDateOpt(dOpt string) (n int, term string) {
	n = 1
	term = ""

	r1 := regexp.MustCompile(`^(\d+)(\w+)`)
	r2 := regexp.MustCompile(`^(\w+)`)

	if r1.MatchString(dOpt) {
		a := r1.FindAllSubmatch([]byte(dOpt), -1)
		n, _ = strconv.Atoi(string(a[0][1]))
		term = string(a[0][2])
	} else if r2.MatchString(dOpt) {
		a := r2.FindAllSubmatch([]byte(dOpt), -1)
		term = string(a[0][1])
	}

	return n, term
}
