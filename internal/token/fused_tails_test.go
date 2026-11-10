//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFusedTails — a tail for each stem the fused syllable may come from, the piece in brackets before the rest of the form; a fused 어/아 that changes no letter is marked as leaving no trace; 셨 is read as the honorific too
package token

import "testing"

func TestFusedTails(t *testing.T) {
	sc := scanOf("갔다")
	got := sc.fusedTails(piece{names: "선어말:ㅆ"}, &entry{lead: jongSS}, 1)
	if len(got) != 2 || got[0].base != '가' || got[0].pieces[0].text != "(았)" || got[1].base != '그' || got[0].start != 1 {
		t.Fatal(got)
	}
	if unseen := scanOf("가").fusedTails(piece{}, &entry{lead: leadVowel}, 1); len(unseen) != 2 || !unseen[0].pieces[0].unseen {
		t.Fatal(unseen)
	}
	if seen := scanOf("봐").fusedTails(piece{}, &entry{lead: leadVowel}, 1); len(seen) != 1 || seen[0].pieces[0].unseen {
		t.Fatal(seen)
	}
	sc = scanOf("마신다")
	got = sc.fusedTails(piece{names: "맺음:ㄴ다"}, &entry{form: "다", lead: jongN, noYo: true}, 2)
	if len(got) != 1 || got[0].base != '시' || got[0].pieces[0].text != "(ㄴ)다" || !got[0].noYo {
		t.Fatal(got)
	}
	var honor bool
	for _, c := range scanOf("가셨다").fusedTails(piece{}, &entry{lead: jongSS}, 2) {
		honor = honor || (c.start == 1 && c.pieces[0].names == honorName)
	}
	if !honor {
		t.Fatal("no honorific under 셨")
	}
}
