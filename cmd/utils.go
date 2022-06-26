package cmd

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Return a weekday in chinese char. If append is not "", add this value it.
// For example, WeekToChineseChar(time., "曜日") #=> "月曜日".
func WeekToChineseChar(key time.Weekday, append string) string {
	var wtc = []string{
		"日", "月", "火", "水", "木", "金", "土",
	}

	if append != "" {
		for i, v := range wtc {
			wtc[i] = v + append
		}
	}
	return wtc[key]
}

func WeekdayAbbreviated(key time.Weekday) string {
	var wtc = []string{
		"Sun", "Mon", "Tue", "Wed", "Thr", "Fri", "Sat",
	}
	return wtc[key]
}

// Return the number of digits for a given number.
func NofDigits(n int) int {
	return len(strconv.Itoa(n))
}

// Padding n with 0 if n is less given digit. For example, PaddingZero(1234, 9)
// returns 123400000. This func is mainly used for calculating nano sec NN
// in HH:MM:SS.NN but useful for more usecases.
func PaddingZero(n int, digit int) int {
	nofDig := NofDigits(n)
	if nofDig < digit {
		return n * int(math.Pow(10, float64(digit-nofDig)))
	} else {
		return n
	}
}

// Print the result as similar to date command
// such as "火  5 10 00:00:00 JST 2022".
func PrintDatenize(d time.Time, platform string) {
	tzName, _ := d.Zone()
	if platform == "macOS" {
		fmt.Printf("%s %2d %2d %02d:%02d:%02d %s %d\n",
			WeekToChineseChar(d.Weekday(), ""), d.Month(), d.Day(),
			d.Hour(), d.Minute(), d.Second(), tzName, d.Year())
	} else { // GNU
		fmt.Printf("%d年%3d月%3d日 %s %02d:%02d:%02d %s\n",
			d.Year(), d.Month(), d.Day(), WeekToChineseChar(d.Weekday(), "曜日"),
			d.Hour(), d.Minute(), d.Second(), tzName)
	}
}

// Print date in RFC5322 format, for example, "Sun, 26 Jun 2022 20:18:50 +0900".
func PrintRFC5322Format(d time.Time) {
	fmt.Printf("%s, %02d %s %4d %02d:%02d:%02d %s\n",
		WeekdayAbbreviated(d.Weekday()), d.Day(), d.Month(), d.Year(),
		d.Hour(), d.Minute(), d.Second(), GetOffsetInString(d))
}

// Return "+0900" for JST.
func GetOffsetInString(d time.Time) string {
	_, offset := d.Zone()

	pref := ""
	if offset < 0 {
		pref = "-"
		offset = -offset
	} else {
		pref = "+"
	}
	h := offset / 3600
	s := fmt.Sprintf("%s%02d00", pref, h)

	return s
}

// Return true if given string s is found in array.
func Included(s string, ary []string) bool {
	for i := range ary {
		if s == ary[i] {
			return true
		}
	}
	return false
}

// Return an array of terms split with single/multiple space(s) from given s.
// "a   b   c" #=> ["a", "b", "c"]
func SplitWithSpace(s string) []string {
	// Remove all redundant spaces.
	reSpaces := regexp.MustCompile(`\s+`)
	s = reSpaces.ReplaceAllString(s, " ")
	terms := strings.Split(s, " ")

	return terms
}

// Return time.Layout of given formatted string of datetime.
func FindLayout(s string) string {
	var (
		res string
	)

	if _, err := time.Parse(time.Layout, s); err == nil {
		res = time.Layout
	} else if _, err := time.Parse(time.ANSIC, s); err == nil {
		res = time.ANSIC
	} else if _, err := time.Parse(time.UnixDate, s); err == nil {
		res = time.UnixDate
	} else if _, err := time.Parse(time.RubyDate, s); err == nil {
		res = time.RubyDate
	} else if _, err := time.Parse(time.RFC822, s); err == nil {
		res = time.RFC822
	} else if _, err := time.Parse(time.RFC822Z, s); err == nil {
		res = time.RFC822Z
	} else if _, err := time.Parse(time.RFC850, s); err == nil {
		res = time.RFC850
	} else if _, err := time.Parse(time.RFC1123, s); err == nil {
		res = time.RFC1123
	} else if _, err := time.Parse(time.RFC1123Z, s); err == nil {
		res = time.RFC1123Z
	} else if _, err := time.Parse(time.RFC3339, s); err == nil {
		res = time.RFC3339
	} else if _, err := time.Parse(time.RFC3339Nano, s); err == nil {
		res = time.RFC3339Nano
	} else if _, err := time.Parse(time.Kitchen, s); err == nil {
		res = time.Kitchen
	} else if _, err := time.Parse(time.Stamp, s); err == nil {
		res = time.Stamp
	} else if _, err := time.Parse(time.StampMilli, s); err == nil {
		res = time.StampMilli
	} else if _, err := time.Parse(time.StampMicro, s); err == nil {
		res = time.StampMicro
	} else if _, err := time.Parse(time.StampNano, s); err == nil {
		res = time.StampNano
	} else {
		return ""
	}

	return res
}
