package cmd

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestModDate(t *testing.T) {
	now := time.Now()

	// TODO: support several patterns of dOpt
	dOpt := "1 hour"
	modTime, _ := modDate(now, dOpt)
	modTime2 := now.Add(time.Hour * 1)

	if (modTime.Year() != modTime2.Year()) || (modTime.Month() != modTime2.Month()) ||
		(modTime.Day() != modTime2.Day()) || (modTime.Hour() != modTime2.Hour()) ||
		(modTime.Minute() != modTime2.Minute()) || (modTime.Second() != modTime2.Second()) {
		t.Errorf("Mod time should be '%s', but '%s'.", modTime2, modTime)
	}
}

func TestUpdateTimeWithDOpt(t *testing.T) {
	loc := time.Now().Location()
	tm := time.Date(2022, 0, 0, 0, 0, 0, 0, loc)

	dtSet := map[string]time.Time{
		"": tm,
		"2022/05/10 11:22:33.1234": time.Date(
			2022, time.Month(5), 10, 11, 22, 33, PaddingZero(1234, 9), loc),
		"2022/05/10 11:22:33": time.Date(
			2022, time.Month(5), 10, 11, 22, 33, 0, loc),
		"2022/05/10 11:22": time.Date(
			2022, time.Month(5), 10, 11, 22, 0, 0, loc),
		"2022/05/10 11": time.Date(
			2022, time.Month(5), 10, 11, 0, 0, 0, loc),

		"05/10 11:22:33.1234": time.Date(
			tm.Year(), time.Month(5), 10, 11, 22, 33, PaddingZero(1234, 9), loc),
		"05/10 11:22:33": time.Date(
			tm.Year(), time.Month(5), 10, 11, 22, 33, 0, loc),
		"05/10 11:22": time.Date(
			tm.Year(), time.Month(5), 10, 11, 22, 0, 0, loc),
		"05/10 11": time.Date(
			tm.Year(), time.Month(5), 10, 11, 0, 0, 0, loc),

		"2022/05/10": time.Date(
			2022, time.Month(5), 10, 0, 0, 0, 0, loc),
		"05/10": time.Date(
			tm.Year(), time.Month(5), 10, 0, 0, 0, 0, loc),
	}

	for k, v := range dtSet {
		updated, _ := updateTimeWithoutModifier(tm, k)

		if v != updated {
			t.Errorf("'%s' != '%s' for '%s'.",
				v, updated, k)
		}
	}
}

func TestInitDateTime(t *testing.T) {
	// TODO
}

func TestParseSingleDateStringOpt(t *testing.T) {
	type NumTerm struct {
		n    int
		term string
	}
	terms := []NumTerm{
		{2, "day"},
		{-2, "month"},
	}
	for _, v := range terms {
		// test for an arg has no spaces
		s := fmt.Sprintf("%d%s", v.n, v.term)
		n, term := parseSingleDateStringOpt(s)
		if n != v.n || term != v.term {
			t.Errorf(
				"The result must be (%d, '%s') for '%s', "+
					"but (%d, '%s') unexpectedly.",
				v.n, v.term, s, n, term)
		}

		// test for an arg has spaces
		s = fmt.Sprintf("%d %s", v.n, v.term)
		n, term = parseSingleDateStringOpt(s)
		if n != v.n || term != v.term {
			t.Errorf(
				"The result must be (%d, '%s') for '%s', "+
					"but (%d, '%s') unexpectedly.",
				v.n, v.term, s, n, term)
		}
	}
}

func TestGetInitDateArray(t *testing.T) {
	strAry := []string{"1", "2", "3"}
	res := getInitDateArray(strAry)
	for i, _ := range res {
		si, _ := strconv.Atoi(strAry[i+1])
		if res[i] != si {
			t.Errorf("%d must be the same as %d.", res[i], si)
		}
	}
}
