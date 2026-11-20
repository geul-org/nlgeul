//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCopulaEnd — 야 and 라고 that end a sentence (before a full stop or a question mark, or at the end) read as endings first, their preference exactly bonusShortTaken above what the stream with no rules gives them, while the stream with no rules reads them as particles first; the trace says the rule stood; before a comma, an exclamation mark or another word the first candidate stays as with no rules
package parse

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamCopulaEnd(t *testing.T) {
	for _, c := range []struct {
		in, text, ending string
	}{
		{"차야.", "차야", "차[열쇠:체언]+(이)[서술격:이(빠짐)]+야[맺음:야]"},
		{"그건 차야?", "차야", "차[열쇠:체언]+(이)[서술격:이(빠짐)]+야[맺음:야]"},
		{"그건 차야", "차야", "차[열쇠:체언]+(이)[서술격:이(빠짐)]+야[맺음:야]"},
		{"책이라고.", "책이라고", "책[열쇠:체언]+이[서술격:이]+라고[맺음:라고]"},
	} {
		got, notes := copulaBranches(t, c.in, c.text, true)
		plain, _ := copulaBranches(t, c.in, c.text, false)
		if len(got) == 0 || got[0].String() != c.ending {
			t.Fatalf("%q: the first candidate is the ending, got %v", c.in, got)
		}
		if len(plain) == 0 || plain[0].String() == c.ending {
			t.Fatalf("%q: with no rules a particle comes first, got %v", c.in, plain)
		}
		after, ok := quotedPref(got, c.ending)
		before, _ := quotedPref(plain, c.ending)
		if !ok || math.Abs(after-before-bonusShortTaken) > 1e-9 {
			t.Errorf("%q: %v -> %v, want exactly bonusShortTaken more", c.in, before, after)
		}
		if !strings.Contains(strings.Join(notes, "\n"), c.ending+"\t"+ruleCopulaEnd+"\t1.00\t통과") {
			t.Errorf("%q: no 통과 line of %s in\n%s", c.in, ruleCopulaEnd, strings.Join(notes, "\n"))
		}
	}
	for _, in := range []string{"차야, 가자.", "차야!", "책이라고 했다."} {
		text := strings.Fields(strings.NewReplacer(",", " ", "!", " ").Replace(in))[0]
		got, _ := copulaBranches(t, in, text, true)
		plain, _ := copulaBranches(t, in, text, false)
		if len(got) == 0 || len(plain) == 0 || got[0].String() != plain[0].String() {
			t.Errorf("%q: the first candidate stays as with no rules: %v, %v", in, got, plain)
		}
	}
}
