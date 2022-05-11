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
	// TODO
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
