//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamGluedShortSpaced — in 패소 자의 몫 the word 자의 has a space before it: the rule of a glued noun of one syllable does not stand (no 통과 line of it), and every candidate keeps the place and the preference the stream with no rules gives it
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamGluedShortSpaced(t *testing.T) {
	got, notes := ruledBranches(t, "words-glued-short.tsv", "패소 자의 몫", "자의", true)
	plain, _ := ruledBranches(t, "words-glued-short.tsv", "패소 자의 몫", "자의", false)
	if len(got) != 2 || len(got) != len(plain) {
		t.Fatal(got, plain)
	}
	for i := range got {
		if got[i].String() != plain[i].String() || got[i].Pref != plain[i].Pref {
			t.Errorf("%d: %v %v, %v %v", i, got[i], got[i].Pref, plain[i], plain[i].Pref)
		}
	}
	if strings.Contains(strings.Join(notes, "\n"), ruleGluedShort+"\t1.00\t통과") {
		t.Errorf("the rule stood with a space before the word:\n%s", strings.Join(notes, "\n"))
	}
}
