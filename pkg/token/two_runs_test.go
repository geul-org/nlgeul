//ff:func feature=token type=reader control=sequence
//ff:what TestTwoRuns — letters that read as one run of markers ending in the middle and another from there to the end (는 then 게) are two runs; one run alone, a first run with nothing after it, a second run with nothing before it, no run at all and a span of one letter are not
package token

import "testing"

func TestTwoRuns(t *testing.T) {
	sc := scanOf("곰는게")
	if sc.twoRuns(1, 3) {
		t.Fatal("no run at all")
	}
	sc.marks[3] = map[int]bool{1: true}
	if sc.twoRuns(1, 3) {
		t.Fatal("one run alone is the old rule's, not this one's")
	}
	sc.marks[2] = map[int]bool{1: true}
	sc.marks[3] = map[int]bool{}
	if sc.twoRuns(1, 3) {
		t.Fatal("a first run with nothing after it")
	}
	sc.marks[2] = map[int]bool{}
	sc.marks[3] = map[int]bool{2: true}
	if sc.twoRuns(1, 3) {
		t.Fatal("a second run with nothing before it")
	}
	sc.marks[2] = map[int]bool{1: true}
	if !sc.twoRuns(1, 3) {
		t.Fatal("는 and then 게")
	}
	if sc.twoRuns(2, 3) {
		t.Fatal("a span of one letter has no middle")
	}
}
