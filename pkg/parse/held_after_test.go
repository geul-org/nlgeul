//ff:func feature=parse type=helper control=sequence
//ff:what TestHeldAfter — a word with a candidate that waits (학생이 has 보격 이, not first) opens the wait; a word whose first candidate is a predicate (되었다) closes an open wait and opens none; a word that is neither keeps the wait as it was; with no rule nothing opens
package parse

import "testing"

func TestHeldAfter(t *testing.T) {
	rules := []sentenceRule{waitingComplementRule()}
	nodes := sentenceRegions(t, "학생이 밥을 되었다")
	subj, obj, pred := firstOf(t, nodes[0]), firstOf(t, nodes[1]), firstOf(t, nodes[2])
	if !heldAfter(false, subj, rules) || !heldAfter(true, subj, rules) || heldAfter(false, subj, nil) {
		t.Fatal("학생이 opens the wait")
	}
	if !heldAfter(true, obj, rules) || heldAfter(false, obj, rules) {
		t.Fatal("밥을 keeps the wait as it was")
	}
	if heldAfter(true, pred, rules) || heldAfter(false, pred, rules) {
		t.Fatal("되었다 closes the wait")
	}
}
