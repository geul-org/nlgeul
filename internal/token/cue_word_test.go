//ff:func feature=token type=reader control=sequence
//ff:what TestCueWord — a word in the letters from a place stands where it ends, and shows the letters before it when the letter right before it is a Hangul syllable; at the unit's head or after a sign it only stands; a word that is not there, runs past the end or ends inside an atom is no cue
package token

import "testing"

func TestCueWord(t *testing.T) {
	fresh := func(text string) *scan {
		sc := spacingScan(text)
		sc.stands, sc.shows = make([]map[int]bool, len(sc.s)+1), map[int]map[int]bool{}
		return sc
	}
	sc := fresh("뚜루뚜겸총무")
	sc.cueWord(3, "겸")
	if !sc.stands[4][3] || !sc.shows[3][4] {
		t.Fatal(sc.stands, sc.shows)
	}
	head := fresh("겸총무")
	head.cueWord(0, "겸")
	if !head.stands[1][0] || len(head.shows) != 0 {
		t.Fatal("at the head", head.stands, head.shows)
	}
	sign := fresh("솔·겸총무")
	sign.cueWord(2, "겸")
	if !sign.stands[3][2] || len(sign.shows) != 0 {
		t.Fatal("after a sign", sign.stands, sign.shows)
	}
	skip := fresh("뚜루뚜겸총무")
	skip.cueWord(3, "및")
	skip.cueWord(5, "총무다")
	skip.inside[4] = true
	skip.cueWord(3, "겸")
	if !standsEmpty(skip.stands) || len(skip.shows) != 0 {
		t.Fatal("no cue", skip.stands, skip.shows)
	}
}
