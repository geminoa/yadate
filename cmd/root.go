package cmd

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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

	if layout := findLayout(dOpt); layout != "" {
		t, _ = time.Parse(layout, dOpt)
	}

	dOptTerms := splitWithSpace(dOpt)
	agoPos := []int{} // Positions of "ago" in dOptTerms
	for i, t := range dOptTerms {
		if t == "ago" {
			agoPos = append(agoPos, i)
		}
	}
	// Change numbers to negative if it's mentions as "ago".
	for _, v := range agoPos {
		if v-1 < 0 {
			// TODO: raise a panic because it cannot be happend.
		} else if v-2 < 0 { // case of first value with no number such as "-d year ago"
			n, t := parseSingleDateOpt(dOptTerms[v-1])
			dOptTerms[v-1] = fmt.Sprintf("%d%s", -n, t)
		} else {
			if n, err := strconv.Atoi(dOptTerms[v-2]); err == nil {
				dOptTerms[v-2] = fmt.Sprintf("%d", -n)
			} else {
				n, t := parseSingleDateOpt(dOptTerms[v-1])
				dOptTerms[v-1] = fmt.Sprintf("%d%s", -n, t)
			}
		}
	}

	// Remove entries "ago" from dOptTerms.
	sort.Sort(sort.Reverse(sort.IntSlice(agoPos)))
	for _, v := range agoPos {
		dOptTerms = append(dOptTerms[:v], dOptTerms[v+1:]...)
	}

	//rDigits := regexp.MustCompile(`^(\d{1,14})$`)
	//rDate := regexp.MustCompile(`^(\d{1,4})/(\d{1,2})/(\d{1,2})`)

	rNum := regexp.MustCompile(`^-?(\d+)`)
	rStr := regexp.MustCompile(`^(\w+)`)
	for i := 0; i < len(dOptTerms)-1; i++ {
		// TODO: combine num and term if there are separated
		if rNum.MatchString(dOptTerms[i]) && rStr.MatchString(dOptTerms[i+1]) {
			dOptTerms[i] = dOptTerms[i] + dOptTerms[i+1]
			dOptTerms[i+1] = ""
		}

	}
	dTerms := []string{}
	for _, v := range dOptTerms {
		if v != "" {
			dTerms = append(dTerms, v)
		}
	}

	for _, v := range dTerms {
		n, term := parseSingleDateOpt(v)
		dt = updateDateTime(dt, n, term)
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

func updateDateTime(dt DateTime, n int, term string) DateTime {
	switch term {
	case "yesterday":
		dt.Day -= 1
	case "tomorrow":
		dt.Day += 1
	case "week", "weeks":
		dt.Day += n * 7
	case "fortnight", "fortnights":
		dt.Day += n * 14
	case "year", "years":
		dt.Year += n * 1
	case "month", "months":
		dt.Month += n * 1
	case "day", "days":
		dt.Day += n * 1
	case "hour", "hours":
		dt.Hour += time.Duration(n * 1)
	case "minute", "minutes":
		dt.Min += time.Duration(n * 1)
	case "second", "seconds":
		dt.Second += time.Duration(n * 1)
	default:
		// TODO
	}
	return dt
}
