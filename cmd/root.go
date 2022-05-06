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
	Year, Month, Day  int
	Hour, Min, Second time.Duration
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
			dt.Day -= 1
		case "tomorrow":
			dt.Day += 1
		case "week", "weeks":
			dt.Day += n * 7
		case "fortnight", "fortnights":
			dt.Day += n * 14
		case "day", "days":
			dt.Month += n * 1
		case "month", "months":
			dt.Month += n * 1
		case "year", "years":
			dt.Year += n * 1
		default:
			// do nothing
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
	t = t.Add(time.Hour*dt.Hour + time.Minute*dt.Min + time.Second*dt.Second)
	return t
}

func parseSingleDateOpt(dOpt string) (n int, term string) {
	n = 1
	term = ""

	r := regexp.MustCompile(`^(-?)(\d*)(\w*)`)

	if r.MatchString(dOpt) {
		a := r.FindAllSubmatch([]byte(dOpt), -1)
		n, _ = strconv.Atoi(string(a[0][2]))
		if n == 0 { // n is 0 if no num in dOpt, but should be 1 in this case.
			n = 1
		}
		if len(a[0][1]) != 0 { // including '-' for negative number
			n = -n
		}
		term = string(a[0][3])
	}

	return n, term
}
