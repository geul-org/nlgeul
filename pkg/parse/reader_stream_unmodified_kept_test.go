//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamUnmodifiedKept — before the noun 사람 (an adnominal may modify it) the rule does not stand: the copula's modifying form 감정+이+ㄴ stays first in the word 감정인 and every candidate keeps the preference the stream with no rules gives it
package parse

import "testing"

func TestReaderStreamUnmodifiedKept(t *testing.T) {
	got, _ := ruledBranches(t, "words-unmodified.tsv", "감정인 사람", "감정인", true)
	plain, _ := ruledBranches(t, "words-unmodified.tsv", "감정인 사람", "감정인", false)
	if len(got) == 0 || len(got) != len(plain) || got[0].String() != "감정[열쇠:체언]+이[서술격:이]+(ㄴ)[꾸밈:ㄴ]" {
		t.Fatal(got, plain)
	}
	for i := range got {
		if got[i].String() != plain[i].String() || got[i].Pref != plain[i].Pref {
			t.Errorf("%d: %v %v, %v %v", i, got[i], got[i].Pref, plain[i], plain[i].Pref)
		}
	}
}
