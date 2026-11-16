//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamUnmodified — before 및 and at the end of a sentence (nothing an adnominal modifies follows) the noun 감정인 of the table comes first in the word 감정인, with exactly bonusShortTaken more preference than the stream with no rules gives it, and the copula's modifying form 감정+이+ㄴ stays as it was behind it; the trace says the rule stood
package token

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamUnmodified(t *testing.T) {
	const noun, copula = "감정인[열쇠:체언]", "감정[열쇠:체언]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]"
	for _, in := range []string{"감정인 및 사람", "사람은 감정인."} {
		got, notes := ruledBranches(t, "words-unmodified.tsv", in, "감정인", true)
		plain, _ := ruledBranches(t, "words-unmodified.tsv", in, "감정인", false)
		if len(got) == 0 || got[0].String() != noun || plain[0].String() != copula || len(got) != len(plain) {
			t.Fatalf("%q: %v, %v", in, got, plain)
		}
		after, _ := quotedPref(got, noun)
		before, _ := quotedPref(plain, noun)
		if math.Abs(after-before-bonusShortTaken) > 1e-9 {
			t.Errorf("%q: %v -> %v, want exactly bonusShortTaken more", in, before, after)
		}
		rest, _ := quotedPref(got, copula)
		if was, _ := quotedPref(plain, copula); rest != was {
			t.Errorf("%q: the copula reading stays as it was: %v, %v", in, was, rest)
		}
		if !strings.Contains(strings.Join(notes, "\n"), noun+"\t"+ruleUnmodified+"\t1.00\t통과") {
			t.Errorf("%q: no 통과 line of %s in\n%s", in, ruleUnmodified, strings.Join(notes, "\n"))
		}
	}
}
