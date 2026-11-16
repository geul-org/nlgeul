//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamGluedShort — in 패소자의 몫 the noun of one syllable with its particle written onto 패소 (자+의) comes first in the word 자의, with exactly bonusShortTaken more preference than the stream with no rules gives it, and the noun 자의 stays as it was behind it; no candidate is taken away and the trace says the rule stood
package token

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamGluedShort(t *testing.T) {
	const short, long = "자[열쇠:체언]+의[관형격:의]", "자의[열쇠:체언]"
	got, notes := ruledBranches(t, "words-glued-short.tsv", "패소자의 몫", "자의", true)
	plain, _ := ruledBranches(t, "words-glued-short.tsv", "패소자의 몫", "자의", false)
	if len(got) != 2 || got[0].String() != short || plain[0].String() != long || len(plain) != len(got) {
		t.Fatal(got, plain)
	}
	after, _ := quotedPref(got, short)
	before, _ := quotedPref(plain, short)
	if math.Abs(after-before-bonusShortTaken) > 1e-9 {
		t.Errorf("%v -> %v, want exactly bonusShortTaken more", before, after)
	}
	if was, _ := quotedPref(plain, long); got[1].Pref != was {
		t.Errorf("자의 stays as it was: %v, %v", was, got[1].Pref)
	}
	if !strings.Contains(strings.Join(notes, "\n"), short+"\t"+ruleGluedShort+"\t1.00\t통과") {
		t.Errorf("no 통과 line of %s in\n%s", ruleGluedShort, strings.Join(notes, "\n"))
	}
}
