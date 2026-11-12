//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what TestPage — page 0 or below is the whole list; a page is cut by its place in the list, with the number of combinations before it; the last page holds what is left, full or not; a page past the last one, and any page of an empty list, is empty with the length of the list before it; a page size past the list makes one page; the greatest page and page size do not overflow
package hypo

import (
	"math"
	"slices"
	"testing"
)

func TestPage(t *testing.T) {
	// Seven combinations in the order List would give them. Their numbers
	// are not their places, so a page cut by number would show.
	listed := []Hypothesis{
		{Number: 4, Score: Score{Place: 6, Verdict: 1, Pref: 6}},
		{Number: 0, Score: Score{Place: 5, Verdict: 1, Pref: 5}},
		{Number: 6, Score: Score{Place: 4, Verdict: 1, Pref: 4}},
		{Number: 2, Score: Score{Place: 3, Verdict: 1, Pref: 3}},
		{Number: 5, Score: Score{Place: 2, Verdict: 1, Pref: 2}},
		{Number: 1, Score: Score{Place: 1, Verdict: 1, Pref: 1}},
		{Number: 3, Score: Score{Place: 0, Verdict: 1, Pref: 0}},
	}
	cases := []struct {
		n             int // the page is asked of listed[:n]
		page, perPage int
		first, end    int // want first before the page and listed[first:end]
	}{
		{7, 0, 3, 0, 7},
		{7, -2, 3, 0, 7},
		{7, math.MinInt, 3, 0, 7},
		{7, 1, 3, 0, 3},
		{7, 2, 3, 3, 6},
		{7, 3, 3, 6, 7},
		{7, 4, 3, 7, 7},
		{6, 2, 3, 3, 6},
		{6, 3, 3, 6, 6},
		{7, 1, 1, 0, 1},
		{7, 4, 1, 3, 4},
		{7, 7, 1, 6, 7},
		{7, 8, 1, 7, 7},
		{7, 1, 7, 0, 7},
		{7, 2, 7, 7, 7},
		{7, 1, 10, 0, 7},
		{7, 2, 10, 7, 7},
		{0, 0, 3, 0, 0},
		{0, 1, 3, 0, 0},
		{0, 1, 1, 0, 0},
		{0, 2, 3, 0, 0},
		{0, math.MaxInt, math.MaxInt, 0, 0},
		{7, math.MaxInt, 3, 7, 7},
		{7, math.MaxInt/3 + 2, 3, 7, 7},
		{7, 1, math.MaxInt, 0, 7},
		{7, 2, math.MaxInt, 7, 7},
		{7, math.MaxInt, math.MaxInt, 7, 7},
	}
	for _, c := range cases {
		first, got := Page(listed[:c.n], c.page, c.perPage)
		if first != c.first || !slices.Equal(got, listed[c.first:c.end]) {
			t.Errorf("page %d of %d of %d combinations: %d before, %+v; want %d before, %+v", c.page, c.perPage, c.n, first, got, c.first, listed[c.first:c.end])
		}
	}
}
