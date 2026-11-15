//ff:func feature=token type=rule control=sequence
//ff:what TestComplementRuleNotBecome — when the word right after 이/가 is no 되다 or 아니다 (학생이 밥을 먹었다) the 주격 stays first, the 보격 is not raised, and the trace says 기각 with the reason that the next word is not 되다·아니다
package token

import (
	"strings"
	"testing"
)

func TestComplementRuleNotBecome(t *testing.T) {
	plain, _ := complementNodes(complementStream(t, "학생이 밥을 먹었다", false, false))
	ruled, notes := complementNodes(complementStream(t, "학생이 밥을 먹었다", true, true))
	n := ruled[0]
	if n.Contest.Branches[0].String() != "학생[열쇠:체언]+이[주격:이]" {
		t.Fatal(n.Contest.Branches)
	}
	comp, _ := branchNamed(n, "학생[열쇠:체언]+이[보격:이]")
	old, _ := branchNamed(plain[0], "학생[열쇠:체언]+이[보격:이]")
	if !almost(comp.Pref, old.Pref) {
		t.Fatal("보격 is not raised", comp.Pref, old.Pref)
	}
	if !strings.Contains(notes[0], "학생[열쇠:체언]+이[보격:이]\t문장:되다·아니다 앞 보격\t-1.00\t기각: 바로 뒤 어절이 되다·아니다가 아니다") {
		t.Fatalf("%q", notes[0])
	}
	if strings.Contains(notes[1], "문장:되다·아니다 앞 보격") {
		t.Fatalf("밥을 is not asked: %q", notes[1])
	}
}
