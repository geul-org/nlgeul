//ff:func feature=token type=rule control=sequence
//ff:what TestStreamJudgeSeat — each candidate of a seat is judged with the seats after it (학생이 before 되었다: 보격 이 gains 2 and is the only lift and the only trace line, since no rule speaks of the others) and the seat reads as its first candidate after the lift; the next seat reads the seats before it as they stand after their own judging (되었다 gains 3 after 보격 이), and with no seat before it the words judged earlier (behind); a seat nothing is gained at gives no lift
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestStreamJudgeSeat(t *testing.T) {
	st := sentenceStream(t, "", true, complementRule(), afterComplementRule())
	seats := seatsOf(sentenceRegions(t, "학생이 되었다"))
	if seats[0].near.Last != "주격:이" {
		t.Fatal("주격 first before judging", seats[0].near)
	}
	lifts, notes := st.judgeSeat(seats, 0)
	if len(lifts) != 1 || lifts[0] != (lift{span: Span{0, 3}, word: "학생[열쇠:체언]+이[보격:이]", add: 2}) || len(notes) != 1 {
		t.Fatal(lifts, notes)
	}
	if seats[0].near.Last != "보격:이" {
		t.Fatal("the seat reads as its first candidate after the lift", seats[0].near)
	}
	lifts, notes = st.judgeSeat(seats, 1)
	if len(lifts) != 1 || lifts[0] != (lift{span: Span{3, 6}, word: "되[열쇠:동사]+었[선어말:었]+다[맺음:다]", add: 3}) || len(notes) != 1 || notes[0] != "  · 되[열쇠:동사]+었[선어말:었]+다[맺음:다]\t되 앞 보격\t1.00\t통과" {
		t.Fatal(lifts, notes)
	}
	alone := seatsOf(sentenceRegions(t, "되었다"))
	if lifts, _ := st.judgeSeat(alone, 0); lifts != nil {
		t.Fatal("nothing before 되었다", lifts)
	}
	st.behind = []rule.Near{{Stem: "학생", Last: "보격:이"}}
	if lifts, _ := st.judgeSeat(alone, 0); len(lifts) != 1 || lifts[0].add != 3 {
		t.Fatal("the word judged earlier is before it", lifts)
	}
}
