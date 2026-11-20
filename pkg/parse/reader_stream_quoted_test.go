//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamQuoted — right after a closing quotation mark 라고 and 이라고 read as the particle of direct quotation first (the closing candidate of the same word follows) and the preference of that candidate is exactly bonusSameLetters above what the stream with no rules gives it, the others stay as they were; the trace says the rule stood
package parse

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamQuoted(t *testing.T) {
	for _, c := range []struct {
		in, text, quote, other string
	}{
		{"\"좋다\"라고 했다", "라고", "+라고[부사격:라고]", "+라고[맺음:라고]"},
		{"\"책\"이라고 했다", "이라고", "+이라고[부사격:이라고]", "+이[서술격:이]+라고[맺음:라고]"},
	} {
		got, notes := quotedBranches(t, c.in, c.text, true)
		plain, _ := quotedBranches(t, c.in, c.text, false)
		if len(got) == 0 || got[0].String() != c.quote {
			t.Fatalf("%q: the first candidate is the particle of quotation, got %v", c.in, got)
		}
		after, ok := quotedPref(got, c.quote)
		before, _ := quotedPref(plain, c.quote)
		if !ok || math.Abs(after-before-bonusSameLetters) > 1e-9 {
			t.Errorf("%q: %v -> %v, want exactly bonusSameLetters more", c.in, before, after)
		}
		rest, ok := quotedPref(got, c.other)
		was, _ := quotedPref(plain, c.other)
		if !ok || rest != was {
			t.Errorf("%q: %s stays after it as it was (%v, %v)", c.in, c.other, was, rest)
		}
		if len(got) != len(plain) {
			t.Errorf("%q: no candidate is taken away: %d of %d", c.in, len(got), len(plain))
		}
		if !strings.Contains(strings.Join(notes, "\n"), c.quote+"\t"+ruleQuoted+"\t1.00\t통과") {
			t.Errorf("%q: no 통과 line of %s in\n%s", c.in, ruleQuoted, strings.Join(notes, "\n"))
		}
	}
}
