//ff:func feature=token type=reader control=sequence
//ff:what TestHonors — 가셨 reads 가+시 under the letter; 읽으셨 also 읽+으+시; a letter that is not listed, a consonant before it, or nothing before it give no such candidate; with a word table 신 is read so too, at the unit's end or under a whole closing, not before a particle
package token

import "testing"

func TestHonors(t *testing.T) {
	p := piece{text: "(었)"}
	got := scanOf("가셨다").honors(p, 2, false)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != honorName || got[0].pieces[1].text != "(었)" {
		t.Fatal(got)
	}
	got = scanOf("읽으셨다").honors(p, 3, false)
	if len(got) != 2 || got[0].start != 2 || got[1].start != 1 || got[1].pieces[0].names != euName || got[1].pieces[1].text != "시" || len(got[1].pieces) != 3 {
		t.Fatal(got)
	}
	if got = scanOf("가신다").honors(p, 2, false); got != nil {
		t.Fatal("a letter not listed", got)
	}
	if got = scanOf("읽셨다").honors(p, 2, false); len(got) != 0 {
		t.Fatal("after a consonant", got)
	}
	if got = scanOf("셨다").honors(p, 1, false); len(got) != 0 {
		t.Fatal("nothing before", got)
	}
	if got = scanOf("가으셨다").honors(p, 3, false); len(got) != 1 {
		t.Fatal("으 after a vowel is no link", got)
	}
	wide := wordsScan("가신", layerCost)
	got = wide.honors(p, 2, true)
	if len(got) != 1 || got[0].gate != wide.ix.pred || got[0].start != 1 {
		t.Fatal("신 at the unit's end", got)
	}
	if got = wordsScan("가신이", layerCost).honors(p, 2, true); len(got) != 0 {
		t.Fatal("신 before a particle", got)
	}
	if got = wordsScan("가신이", layerCost).honors(p, 2, false); len(got) != 1 {
		t.Fatal("신 under a whole closing", got)
	}
	if got = wordsScan("먹으신", layerCost).honors(p, 3, true); len(got) != 2 || got[1].gate == nil {
		t.Fatal("으 before 신", got)
	}
}
