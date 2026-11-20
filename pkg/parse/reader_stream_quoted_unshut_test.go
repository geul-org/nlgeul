//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamQuotedUnshut — with no closing sign right before it 라고 and 이라고 gain nothing (책이라고, 좋다라고) and the trace says 기각 with the reason that the word before is no closing sign; after a closing bracket the rule stands too (the list asks only for the kind 닫기), and a word before that closes nothing gives no line to the claimed candidates
package parse

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamQuotedUnshut(t *testing.T) {
	for _, c := range []struct {
		in, text, cand string
	}{
		{"책이라고 했다", "책이라고", "책[열쇠:체언]+이라고[부사격:이라고]"},
		{"좋다라고 했다", "좋다라고", "좋다+라고[부사격:라고]"},
	} {
		got, notes := quotedBranches(t, c.in, c.text, true)
		plain, _ := quotedBranches(t, c.in, c.text, false)
		after, ok := quotedPref(got, c.cand)
		before, _ := quotedPref(plain, c.cand)
		if !ok || after != before {
			t.Errorf("%q: %s gains nothing (%v, %v)", c.in, c.cand, before, after)
		}
		want := c.cand + "\t" + ruleQuoted + "\t-1.00\t기각: " + whyNotShut
		if !strings.Contains(strings.Join(notes, "\n"), want) {
			t.Errorf("%q: no line %q in\n%s", c.in, want, strings.Join(notes, "\n"))
		}
	}
	got, notes := quotedBranches(t, "(책)라고 했다", "라고", true)
	plain, _ := quotedBranches(t, "(책)라고 했다", "라고", false)
	after, ok := quotedPref(got, "+라고[부사격:라고]")
	before, _ := quotedPref(plain, "+라고[부사격:라고]")
	if !ok || math.Abs(after-before-bonusSameLetters) > 1e-9 {
		t.Errorf("after a closing bracket: %v -> %v", before, after)
	}
	if !strings.Contains(strings.Join(notes, "\n"), "+라고[부사격:라고]\t"+ruleQuoted+"\t1.00\t통과") {
		t.Errorf("no 통과 line after a closing bracket in\n%s", strings.Join(notes, "\n"))
	}
}
