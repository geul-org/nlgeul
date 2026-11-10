//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestChains — 학교에서는 gives 는 alone and 에서+는; 결과가 gives 가 alone and not 과+가; 마당에서 gives 에서 and 서 but not 에+서; nothing at the start, nothing past three particles; 이건요 reads the fused form before 요; with a word table 선 is read after a nominal's place with its gate, and a particle after 심 takes the honorific noun form; a particle after a unit takes the unit with it
package token

import "testing"

func TestChains(t *testing.T) {
	sc := scanOf("학교에서는")
	starts := map[int]int{}
	for _, c := range sc.chains(5, 0, -1) {
		starts[c.start] = max(starts[c.start], len(c.pieces))
	}
	if starts[4] != 1 || starts[2] != 2 || starts[1] != 0 {
		t.Fatal(starts)
	}
	got := scanOf("결과가").chains(3, 0, -1)
	if len(got) != 1 || got[0].start != 2 {
		t.Fatal(got)
	}
	for _, c := range scanOf("마당에서").chains(4, 0, -1) {
		if len(c.pieces) == 2 && c.pieces[0].text == "에" {
			t.Fatal("에+서 read as a chain", c)
		}
	}
	if sc.chains(0, 0, -1) != nil || sc.chains(5, 3, -1) != nil {
		t.Fatal("limits")
	}
	var fixed bool
	for _, c := range scanOf("이건요").chains(3, 0, -1) {
		fixed = fixed || (c.shown == "이것" && len(c.pieces) == 2)
	}
	if !fixed {
		t.Fatal("no fused form before 요")
	}
	var short, noun bool
	for _, c := range wordsScan("학교선", layerCost).chains(3, 0, -1) {
		short = short || (c.gate != nil && c.pieces[0].text == "선")
	}
	for _, c := range wordsScan("가심을", layerCost).chains(3, 0, -1) {
		noun = noun || (len(c.pieces) == 3 && c.pieces[1].text == "(ㅁ)")
	}
	if !short || !noun {
		t.Fatal(short, noun)
	}
	var unit bool
	for _, c := range scanOf("3개를").chains(3, 0, -1) {
		unit = unit || (c.start == 1 && len(c.pieces) == 2 && c.pieces[0].names == "단위:개")
	}
	if !unit {
		t.Fatal("no unit before the particle")
	}
}
