package cmd

import (
	"fmt"
	"testing"
	"time"
)

func TestWeekToChineseChar(t *testing.T) {
	var wtc = []string{
		"日", "月", "火", "水", "木", "金", "土",
	}

	for i, v := range wtc {
		wd := WeekToChineseChar(time.Weekday(i))
		if wd != v {
			t.Errorf("The result must be '%s', but '%s", v, wd)
		}
	}
}

func TestIncluded(t *testing.T) {
	ary := []string{"aaa", "bbb", "ccc"}
	s := ary[0]
	if !Included(s, ary) {
		t.Errorf("'%s' should be included in '%s'", s, ary)
	}
}

func TestSplitWithSpace(t *testing.T) {
	ary := []string{"many", "spaces", "around", "there"}
	manySpaces := fmt.Sprintf(
		"%s   %s    %s     %s", ary[0], ary[1], ary[2], ary[3])
	split := SplitWithSpace(manySpaces)
	for i, v := range split {
		if v != ary[i] {
			t.Errorf("'%s' should be included in '%s'", v, ary[i])
		}
	}
}

func TestFindLayout(t *testing.T) {
	// TODO: test all formats just there are still three ...
	formats := map[string]string{
		"RFC822":   time.RFC822,
		"Kitchen":  time.Kitchen,
		"UnixDate": time.UnixDate,
	}

	examples := map[string]string{
		"RFC822":   "10 Nov 09 23:00 UTC",
		"Kitchen":  "11:00PM",
		"UnixDate": "Tue Nov 10 23:00:00 UTC 2009",
	}

	for k, v := range examples {
		layout := FindLayout(v)
		if layout != formats[k] {
			t.Errorf("'%s' must be '%s'.", v, k)
		}
	}
}
