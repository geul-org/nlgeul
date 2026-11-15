//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestComplementRuleAskDenied — the question-pronoun exception falls when the pronoun carries no 이/가 (반장이 누구를 됐어) or what follows it is no 되다·아니다 (반장이 누가 먹었어): 반장이 gets no 보격 preference, the 주격 stays first, and the trace names the reasons that stood
package token

import (
	"strings"
	"testing"
)

func TestComplementRuleAskDenied(t *testing.T) {
	cases := []struct{ sentence, reasons string }{
		{"반장이 누구를 됐어", "기각: 바로 뒤 어절이 되다·아니다가 아니다, 바로 뒤 어절이 의문 대명사다, 의문 대명사에 이·가가 붙지 않았다"},
		{"반장이 누가 먹었어", "기각: 바로 뒤 어절이 되다·아니다가 아니다, 바로 뒤 어절이 의문 대명사다, 의문 대명사 뒤 어절이 되다·아니다가 아니다"},
	}
	for _, c := range cases {
		plain, _ := complementNodes(complementStream(t, c.sentence, false, false))
		ruled, notes := complementNodes(complementStream(t, c.sentence, true, true))
		comp, _ := branchNamed(ruled[0], "반장[열쇠:체언]+이[보격:이]")
		old, _ := branchNamed(plain[0], "반장[열쇠:체언]+이[보격:이]")
		if ruled[0].Contest.Branches[0].String() != "반장[열쇠:체언]+이[주격:이]" || !almost(comp.Pref, old.Pref) {
			t.Fatalf("%s: %v", c.sentence, ruled[0].Contest.Branches)
		}
		if !strings.Contains(notes[0], "반장[열쇠:체언]+이[보격:이]\t문장:되다·아니다 앞 보격\t-1.00\t"+c.reasons) {
			t.Fatalf("%s: %q", c.sentence, notes[0])
		}
	}
}
