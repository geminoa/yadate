package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func weekToChineseChar(key time.Weekday) string {
	var wtc = []string{
		"日", "月", "火", "水", "木", "金", "土",
	}
	return wtc[key]
}

func printDatenize(d time.Time) {
	tzName, _ := d.Zone()
	fmt.Printf("%s %2d %2d %02d:%02d:%02d %s %d\n",
		weekToChineseChar(d.Weekday()), d.Month(), d.Day(),
		d.Hour(), d.Minute(), d.Second(), tzName, d.Year())
}

func included(s string, ary []string) bool {
	for i := range ary {
		if s == ary[i] {
			return true
		}
	}
	return false
}

func splitWithSpace(s string) []string {
	// Remove all redundant spaces.
	reSpaces := regexp.MustCompile(`\s+`)
	s = reSpaces.ReplaceAllString(s, " ")
	terms := strings.Split(s, " ")

	return terms
}

func findLayout(s string) string {
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
