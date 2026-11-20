//ff:func feature=parse type=reader control=sequence
//ff:what TestSignWords — a sign between letters is a word of its own after the candidates kept before it, named 기호, with preference 0 and its own verdict 1 (the candidate holds as much as its weakest piece); a run of the sign (——) is one word ending where the run ends, and nothing ends inside it; a place with no such sign adds nothing
package parse

import "testing"

func TestSignWords(t *testing.T) {
	sc := scanOf("구름·돌")
	best := [][]*path{{{verdict: 1}}, nil, nil, nil, nil}
	got := sc.signWords(nil, best, 3)
	if len(got) != 1 || got[0].last.end != 3 || got[0].last.word.End != 3 || got[0].last.pref != 0 || got[0].verdict != got[0].prev.verdict || got[0].last.verdict != 1 {
		t.Fatal(got)
	}
	n := got[0].last.word.Parts[0].Name
	if n.Kind != midName || got[0].last.word.Parts[0].Text != "·" {
		t.Fatal("named until the stretches are made", n)
	}
	if got = sc.signWords(nil, best, 2); len(got) != 0 {
		t.Fatal("no sign there", got)
	}
	run := scanOf("구름——돌")
	best = [][]*path{{{verdict: 1}}, nil, nil, nil, nil, nil}
	if got = run.signWords(nil, best, 4); len(got) != 1 || got[0].last.word.End != 4 || got[0].last.word.Parts[0].Text != "——" {
		t.Fatal("the run is one word", got)
	}
	if got = run.signWords(nil, best, 3); len(got) != 0 {
		t.Fatal("not inside the run", got)
	}
}
