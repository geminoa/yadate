package cmd

import (
	"fmt"
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
