//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderStreamAlone — a word alone in a pair of brackets reads the letters of a bound noun or a determiner as the other word first (詩[시]를: 시[열쇠:체언] before 시[의존명사:시]; (이): 이[열쇠:체언] before 이[관형사:이]; [시를]: with a particle too), whose preference is exactly bonusSentence above what the stream with no rules gives it while the bound noun or determiner stays as it was and no candidate is taken away, and the trace says the rule stood; with no closing sign right after it ((이 책)) or no opening sign right before it (그 시를) the order is that of the stream with no rules
package token

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamAlone(t *testing.T) {
	for _, c := range []struct {
		in, text, free, barred string
		lifted                 bool
	}{
		{"詩[시]를 읽었다.", "시", "시[열쇠:체언]", "시[의존명사:시]", true},
		{"책 (이) 장", "이", "이[열쇠:체언]", "이[관형사:이]", true},
		{"책 [시를] 읽었다.", "시를", "시[열쇠:체언]+를[목적격:를]", "시[의존명사:시]+를[목적격:를]", true},
		{"책 (이 책)", "이", "이[열쇠:체언]", "이[관형사:이]", false},
		{"그 시를 읽었다.", "시를", "시[열쇠:체언]+를[목적격:를]", "시[의존명사:시]+를[목적격:를]", false},
	} {
		got, notes := aloneBranches(t, c.in, c.text, true)
		plain, _ := aloneBranches(t, c.in, c.text, false)
		if len(got) == 0 || len(got) != len(plain) {
			t.Fatalf("%q: %v of %v", c.in, got, plain)
		}
		after, ok := quotedPref(got, c.free)
		before, _ := quotedPref(plain, c.free)
		kept, _ := quotedPref(got, c.barred)
		was, _ := quotedPref(plain, c.barred)
		if !ok || kept != was {
			t.Errorf("%q: %s stays as it was (%v, %v)", c.in, c.barred, was, kept)
		}
		stood := strings.Contains(strings.Join(notes, "\n"), c.free+"\t"+ruleAlone+"\t1.00\t통과")
		if !c.lifted {
			if after != before || got[0].String() != plain[0].String() || stood {
				t.Errorf("%q: not lifted, got %v", c.in, got)
			}
			continue
		}
		if got[0].String() != c.free || math.Abs(after-before-bonusSentence) > 1e-9 || !stood {
			t.Errorf("%q: %v -> %v, first %s\n%s", c.in, before, after, got[0].String(), strings.Join(notes, "\n"))
		}
	}
}
