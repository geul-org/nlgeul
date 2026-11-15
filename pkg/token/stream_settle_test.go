//ff:func feature=token type=rule control=sequence
//ff:what TestStreamSettle — the regions queued and not judged are judged as one line: each gains its own lifts (학생이 its 보격 이, then 되었다 after it) with its trace lines kept after the lines of its reading, is marked judged, the wait is closed, and the words of the line stand behind what comes next, the nearest first; judging again does nothing; with no rule the regions are not judged and only the wait is closed
package token

import "testing"

func TestStreamSettle(t *testing.T) {
	st := sentenceStream(t, "학생이 되었다", true, complementRule(), afterComplementRule())
	st.pull()
	st.pull()
	st.pull()
	a, b := st.queue[0], st.queue[1]
	readA, readB := len(a.notes), len(b.notes)
	if len(st.queue) != 2 || a.judged || b.judged {
		t.Fatal(len(st.queue))
	}
	st.held = true
	st.settle()
	if st.held || !a.judged || !b.judged || len(a.lifts) != 1 || a.lifts[0].add != 2 || len(b.lifts) != 1 || b.lifts[0].add != 3 {
		t.Fatal(st.held, a.lifts, b.lifts)
	}
	if len(a.notes) != readA+1 || a.notes[readA] != "  · 학생[열쇠:체언]+이[보격:이]\t보격 뒤 되\t1.00\t통과" || len(b.notes) != readB+1 {
		t.Fatalf("%q", a.notes)
	}
	if len(st.behind) != 2 || st.behind[0].Stem != "되" || st.behind[1].Last != "보격:이" {
		t.Fatal(st.behind)
	}
	st.held = true
	st.settle()
	if st.held || len(a.lifts) != 1 || len(b.notes) != readB+1 || len(st.behind) != 2 {
		t.Fatal("all judged already")
	}
	bare := sentenceStream(t, "학생이", false)
	bare.pull()
	bare.held = true
	bare.settle()
	if bare.held || bare.queue[0].judged || bare.queue[0].lifts != nil || bare.behind != nil {
		t.Fatal("no rule")
	}
}
