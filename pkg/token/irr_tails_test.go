//ff:func feature=token type=reader control=sequence
//ff:what TestIrrTails — 은 and 으면 after the ㄹ a ㄷ stem is left with are read by the stem's own ㄷ and ask for that change; 은 after the vowel a ㅅ stem is left with is read by the stem's own last syllable and asks for that change; nothing without a word table, after a consonant, or after a vowel no stem changes into
package token

import "testing"

func TestIrrTails(t *testing.T) {
	got := wordsScan("지은", layerCost).irrTails(2, false)
	if len(got) != 1 || got[0].start != 1 || got[0].irr != irrS || got[0].pieces[0].names != "꾸밈:은" || got[0].close == nil || got[0].head != headEuForm {
		t.Fatal(got)
	}
	got = wordsScan("들은들으면", layerCost).irrTails(2, false)
	if len(got) != 1 || got[0].irr != irrD || got[0].pieces[0].names != "꾸밈:은" {
		t.Fatal("은 after the ㄹ a ㄷ stem is left with", got)
	}
	if got = wordsScan("들은들으면", layerCost).irrTails(5, false); len(got) != 1 || got[0].irr != irrD || got[0].pieces[0].names != "맺음:으면" {
		t.Fatal("으면 after the ㄹ a ㄷ stem is left with", got)
	}
	if got = scanOf("지은").irrTails(2, false); len(got) != 0 {
		t.Fatal("no word table", got)
	}
	if got = wordsScan("먹은", layerCost).irrTails(2, false); len(got) != 0 {
		t.Fatal("after a consonant the closing is read as it is", got)
	}
	if got = wordsScan("가은", layerCost).irrTails(2, false); len(got) != 0 {
		t.Fatal("no stem changes into this syllable", got)
	}
}
