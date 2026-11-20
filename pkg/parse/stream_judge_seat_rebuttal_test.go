//ff:func feature=parse type=rule control=sequence
//ff:what TestStreamJudgeSeatRebuttal — a seat judged by a rule of rebuttal: 주격 이 of 학생이 before 되었다 gets a lift with the force and nothing added, the only lift; the seat then reads as its first candidate after it (보격 이), and the next seat reads it so; two rebuttals that stand give a lift for each
package parse

import "testing"

func TestStreamJudgeSeatRebuttal(t *testing.T) {
	st := sentenceStream(t, "", true, subjectRebuttalRule(), afterComplementRule())
	seats := seatsOf(sentenceRegions(t, "학생이 되었다"))
	if seats[0].near.Last != "주격:이" {
		t.Fatal("주격 first before judging", seats[0].near)
	}
	lifts, notes := st.judgeSeat(seats, 0)
	if len(lifts) != 1 || lifts[0] != (lift{span: Span{0, 3}, word: "학생[열쇠:체언]+이[주격:이]", force: forceUngrammatical}) || len(notes) != 1 {
		t.Fatal(lifts, notes)
	}
	if seats[0].near.Last != "보격:이" {
		t.Fatal("the seat reads as its first candidate after the rebuttal", seats[0].near)
	}
	if lifts, _ := st.judgeSeat(seats, 1); len(lifts) != 1 || lifts[0].add != 3 {
		t.Fatal("the next seat reads 보격 이 before it", lifts)
	}
	twice := subjectRebuttalRule()
	twice.name = "둘째"
	st.rules = []sentenceRule{subjectRebuttalRule(), twice}
	again := seatsOf(sentenceRegions(t, "학생이 되었다"))
	if lifts, _ := st.judgeSeat(again, 0); len(lifts) != 2 || lifts[1].force != forceUngrammatical || lifts[1].add != 0 {
		t.Fatal(lifts)
	}
}
