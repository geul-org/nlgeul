//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestHeldAfterText — a word whose candidate a rule of a sentence speaks of by the word's letters (haOneRule: 하다 with ㄴ written as 한) opens the wait, as judgeSeat gives the letters to the rules; the word before it (일을) does not
package token

import "testing"

func TestHeldAfterText(t *testing.T) {
	st := streamReader(t, "words-ha-one.tsv").StreamTrace("일을 한 번 했다")
	for st.cur != nil || st.read < len(st.units) {
		st.pull()
	}
	rules := []sentenceRule{haOneRule()}
	var one, obj Node
	for _, rg := range st.queue {
		n := firstOf(t, rg)
		switch n.Contest.Text {
		case "한":
			one = n
		case "일을":
			obj = n
		}
	}
	if one.Contest.Text == "" || obj.Contest.Text == "" {
		t.Fatal("no 한 or 일을 among the regions")
	}
	if !heldAfter(false, one, rules) {
		t.Error("한 does not open the wait")
	}
	if heldAfter(false, obj, rules) {
		t.Error("일을 opens the wait")
	}
}
