//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamOverUnit — end to end with no word table: the approximate 여 right after a number and its unit is read as the approximation 어림:여 (5분여, 1시간여, 7년여가), the reading with the copula that came first before stays among the candidates (5분이어); right after a number it is read as before (10여명) and no candidate appears twice
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamOverUnit(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, first string }{
		{"5분여 지났다", "5[수]+분[단위:분]+여[어림:여]"},
		{"1시간여 지났다", "1[수]+시간[단위:시간]+여[어림:여]"},
		{"7년여가 지났다", "7[수]+년[단위:년]+여[어림:여]+가[주격:가]"},
		{"10여명이 왔다", "10[수]+여[어림:여]+명[단위:명]+이[주격:이]"},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(r, c.sentence), "\n") + "\n"
		if !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
		if all := streamText(r, c.sentence); strings.Count(all, c.first) != 1 {
			t.Errorf("%s: %q appears %d times\n%s", c.sentence, c.first, strings.Count(all, c.first), all)
		}
	}
	if all := streamText(r, "5분여 지났다"); !strings.Contains(all, "5[수]+분[단위:분]+이[서술격:이]+(어)[맺음:어(합침)]") {
		t.Errorf("5분여: the reading with the copula is gone\n%s", all)
	}
}
