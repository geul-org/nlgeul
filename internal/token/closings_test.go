//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosings — 먹었다 gives 다 alone and 었+다; 간다 a closing under the consonant; nothing at the start; 곰이다 has the copula reading too; 했잖아 has 지+(않) before the closing; 가야겠다 has (하) between the fused 어야 and 겠; with a word table 가심 has the honorific noun form
package token

import "testing"

func TestClosings(t *testing.T) {
	sc := scanOf("먹었다")
	got := sc.closings(3, false)
	starts := map[int]bool{}
	for _, c := range got {
		starts[c.start] = true
	}
	if !starts[2] || !starts[1] {
		t.Fatal(got)
	}
	sc = scanOf("간다")
	var fused bool
	for _, c := range sc.closings(2, false) {
		fused = fused || (c.pieces[0].text == "(ㄴ)다" && c.base == '가')
	}
	if !fused || sc.closings(0, false) != nil {
		t.Fatal(fused)
	}
	var copula bool
	for _, c := range scanOf("곰이다").closings(3, false) {
		copula = copula || (c.start == 1 && c.pieces[0].names == copulaName)
	}
	if !copula {
		t.Fatal("no copula reading")
	}
	var neg bool
	for _, c := range scanOf("했잖아").closings(3, false) {
		neg = neg || (len(c.pieces) == 4 && c.pieces[2].text == "(않)" && c.base == '하')
	}
	if !neg {
		t.Fatal("no 잖 reading")
	}
	var must bool
	for _, c := range scanOf("가야겠다").closings(4, false) {
		must = must || (c.start == 1 && c.base == '가' && len(c.pieces) == 4 && c.pieces[0].names == "맺음:어야(합침)" && c.pieces[1].text == "(하)" && c.pieces[2].names == guessName)
	}
	if !must {
		t.Fatal("no 겠 read after 어야")
	}
	var lone bool
	for _, c := range wordsScan("가심", layerCost).closings(2, false) {
		lone = lone || (c.gate != nil && c.pieces[0].names == honorName)
	}
	if !lone {
		t.Fatal("no honorific under 심")
	}
}
