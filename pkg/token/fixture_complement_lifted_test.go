//ff:func feature=token type=helper control=sequence
//ff:what complementLifted — checks that in a sentence read with complementRule the word given has the candidate top first, its preference up by exactly bonusSameLetters and the candidate alt's as it was without the rule, and that the trace says 통과 for top
package token

import (
	"strings"
	"testing"
)

func complementLifted(t *testing.T, sentence string, word int, top, alt string) {
	t.Helper()
	plain, _ := complementNodes(complementStream(t, sentence, false, false))
	ruled, notes := complementNodes(complementStream(t, sentence, true, true))
	n := ruled[word]
	if n.Contest.Branches[0].String() != top {
		t.Fatalf("%s: %v", sentence, n.Contest.Branches)
	}
	now, _ := branchNamed(n, top)
	was, _ := branchNamed(plain[word], top)
	other, _ := branchNamed(n, alt)
	otherWas, _ := branchNamed(plain[word], alt)
	if !almost(now.Pref, was.Pref+bonusSameLetters) || !almost(other.Pref, otherWas.Pref) {
		t.Fatalf("%s: only 보격 is raised, by bonusSameLetters: %v", sentence, n.Contest.Branches)
	}
	if !strings.Contains(notes[word], top+"\t문장:되다·아니다 앞 보격\t1.00\t통과") {
		t.Fatalf("%s: %q", sentence, notes[word])
	}
}
