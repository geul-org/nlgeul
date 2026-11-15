//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestComplementRuleAsk — a question pronoun with 이/가 right before 되다 or 아니다 is the subject asked for: in 반장이 누가 됐어 반장이 has the 보격 first and 누가 (read 누구+가 by the contraction table) the 주격 first; in 누가 반장이 됐어 반장이 is the 보격; in 커서 뭐가 될래 (no 이/가 word before it) 뭐가 has the 보격 first; the lifted ones are 1.00 통과 in the trace
package token

import (
	"strings"
	"testing"
)

func TestComplementRuleAsk(t *testing.T) {
	cases := []struct {
		sentence string
		word     int
		top      string
	}{
		{"반장이 누가 됐어", 0, "반장[열쇠:체언]+이[보격:이]"},
		{"반장이 누가 됐어", 1, "누구[열쇠:체언]+가[주격:가]"},
		{"누가 반장이 됐어", 0, "누구[열쇠:체언]+가[주격:가]"},
		{"누가 반장이 됐어", 1, "반장[열쇠:체언]+이[보격:이]"},
		{"커서 뭐가 될래", 1, "뭐[열쇠:체언]+가[보격:가]"},
	}
	for _, c := range cases {
		nodes, notes := complementNodes(complementStream(t, c.sentence, true, true))
		if nodes[c.word].Contest.Branches[0].String() != c.top {
			t.Fatalf("%s %d: %v", c.sentence, c.word, nodes[c.word].Contest.Branches)
		}
		if strings.Contains(c.top, "보격") && !strings.Contains(notes[c.word], c.top+"\t문장:되다·아니다 앞 보격\t1.00\t통과") {
			t.Fatalf("%s %d: %q", c.sentence, c.word, notes[c.word])
		}
	}
	plain, _ := complementNodes(complementStream(t, "반장이 누가 됐어", false, false))
	ruled, _ := complementNodes(complementStream(t, "반장이 누가 됐어", false, true))
	subj, _ := branchNamed(ruled[1], "누구[열쇠:체언]+가[주격:가]")
	was, _ := branchNamed(plain[1], "누구[열쇠:체언]+가[주격:가]")
	if !almost(subj.Pref, was.Pref) {
		t.Fatal("the question word itself is not lifted", subj.Pref, was.Pref)
	}
	comp, _ := branchNamed(ruled[0], "반장[열쇠:체언]+이[보격:이]")
	old, _ := branchNamed(plain[0], "반장[열쇠:체언]+이[보격:이]")
	if !almost(comp.Pref, old.Pref+bonusSentence) {
		t.Fatal("one word back, by bonusSentence", comp.Pref, old.Pref)
	}
}
