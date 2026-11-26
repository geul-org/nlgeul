//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPlainDots — two numbers joined by a full stop with no full stop after them are read as a date (a month and a day: 6.25, a year and a month: 1971.3) and a time (an hour and a minute: 18.30) behind the decimal and the version of the same letters (plainShapes), with a word or a particle right after them too; numbers out of the ranges (13.40 is no date, 6.75 and 0.5 no time, 25.61 neither) get no such candidate
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamPlainDots(t *testing.T) {
	r := streamReader(t, "words-stream.tsv")
	for _, c := range []struct {
		text, nums string
		date, time bool
	}{
		{"6.25", "6.25", true, true}, {"7.4", "7.4", true, false}, {"1971.3", "1971.3", true, false},
		{"18.30", "18.30", false, true}, {"13.40", "13.40", false, true},
		{"6.75", "6.75", false, false}, {"0.5", "0.5", false, false}, {"25.61", "25.61", false, false}, {"1971.13", "1971.13", false, false},
		{"6.25에", "6.25", true, true}, {"11.20.", "11.20", true, true},
	} {
		seg, _ := r.Stream(c.text).Next()
		got := streamText(r, c.text)
		if first := seg.Nodes[0].Contest.Branches[0].String(); !strings.HasPrefix(first, c.nums+"[소수점]") {
			t.Errorf("%s: the decimal stays first: %q", c.text, first)
		}
		if strings.Contains(got, "\n"+c.nums+"[날짜]") != c.date || strings.Contains(got, "\n"+c.nums+"[시각]") != c.time {
			t.Errorf("%s: date %v, time %v\n%s", c.text, c.date, c.time, got)
		}
	}
}
