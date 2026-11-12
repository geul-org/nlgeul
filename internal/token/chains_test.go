//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestChains — 학교에서는 gives 는 alone and 에서+는; 결과가 gives 가 alone and not 과+가; 마당에서 gives 에서 and 서 but not 에+서; nothing at the start; a chain is as long as its particles follow one another (에게서+부터+는+요); 이건요 reads the fused form before 요; with a word table 선 is read after a nominal's place with its gate, and a particle after 심 takes the honorific noun form; a particle after a unit takes the unit with it; the emphatic 가 ends a chain only after a particle (집에서가), not after a number's unit, an honorific or a closing that names or asks
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
	if len(got) != 2 || got[0].start != 2 || got[1].start != 2 || got[0].pieces[0].names == got[1].pieces[0].names {
		t.Fatal(got)
	}
	for _, c := range scanOf("마당에서").chains(4, 0, -1) {
		if len(c.pieces) == 2 && c.pieces[0].text == "에" {
			t.Fatal("에+서 read as a chain", c)
		}
	}
	if sc.chains(0, 0, -1) != nil {
		t.Fatal("limits")
	}
	var four bool
	for _, c := range scanOf("친구에게서부터는요").chains(9, 0, -1) {
		four = four || (c.start == 2 && len(c.pieces) == 4 && c.pieces[0].text == "에게서")
	}
	if !four {
		t.Fatal("no chain of four particles")
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
	emph := func(sc *scan, end int) (n int) {
		for _, c := range sc.chains(end, 0, -1) {
			if c.pieces[len(c.pieces)-1].names == emphName {
				n++
			}
		}
		return n
	}
	if emph(scanOf("집에서가"), 4) == 0 {
		t.Fatal("no emphatic 가 after a particle")
	}
	for _, c := range []struct {
		text string
		end  int
		sc   *scan
	}{
		{"3건가", 3, scanOf("3건가")},
		{"두개가", 3, scanOf("두개가")},
		{"가심가", 3, wordsScan("가심가", layerCost)},
		{"먹기가", 3, wordsScan("먹기가", layerCost)},
		{"먹는지가", 4, wordsScan("먹는지가", layerCost)},
	} {
		if n := emph(c.sc, c.end); n != 0 {
			t.Fatal("emphatic 가 read after a unit, an honorific or a closing:", c.text, n)
		}
	}
}
