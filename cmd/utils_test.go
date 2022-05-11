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
		wd := WeekToChineseChar(time.Weekday(i), "")
		if wd != v {
			t.Errorf("The result must be '%s', but '%s", v, wd)
		}
	}
}

func TestNofDigits(t *testing.T) {
	if NofDigits(1000) != 4 {
		t.Errorf("'%d' must be %d digits.", 1000, 4)
	}
}

func TestPaddingZero(t *testing.T) {
	res := PaddingZero(1234, 9)
	if NofDigits(res) != 9 {
		t.Errorf("'%d' must be %d digits.", res, 9)
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
	formats := map[string][]string{
		"RFC822":   []string{time.RFC822, "10 Nov 09 23:00 UTC"},
		"Kitchen":  []string{time.Kitchen, "11:00PM"},
		"UnixDate": []string{time.UnixDate, "Tue Nov 10 23:00:00 UTC 2009"},
	}

	for k, v := range formats {
		layout := FindLayout(v[1])
		if layout != v[0] {
			t.Errorf("'%s' must be '%s'.", v, k)
		}
	}
}
