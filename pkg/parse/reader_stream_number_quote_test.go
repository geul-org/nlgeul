//ff:func feature=parse type=reader control=sequence
//ff:what TestReaderStreamNumberQuote — 라고 right on a circled number reads as the particle of quotation first, its preference exactly bonusSameLetters above what the stream with no rules gives it, while the stream with no rules reads it as the ending first; the trace says the rule stood
package parse

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamNumberQuote(t *testing.T) {
	const quote, ending = "②[수]+라고[부사격:라고]", "②[수]+라고[맺음:라고]"
	got, notes := copulaBranches(t, "이것을 ②라고 한다.", "②라고", true)
	plain, _ := copulaBranches(t, "이것을 ②라고 한다.", "②라고", false)
	if len(plain) == 0 || plain[0].String() != ending {
		t.Fatalf("with no rules the ending comes first, got %v", plain)
	}
	if len(got) == 0 || got[0].String() != quote {
		t.Fatalf("the first candidate is the particle of quotation, got %v", got)
	}
	after, ok := quotedPref(got, quote)
	before, _ := quotedPref(plain, quote)
	if !ok || math.Abs(after-before-bonusSameLetters) > 1e-9 {
		t.Errorf("%v -> %v, want exactly bonusSameLetters more", before, after)
	}
	if !strings.Contains(strings.Join(notes, "\n"), quote+"\t"+ruleNumberQuote+"\t1.00\t통과") {
		t.Errorf("no 통과 line of %s in\n%s", ruleNumberQuote, strings.Join(notes, "\n"))
	}
}
