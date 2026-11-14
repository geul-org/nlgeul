//ff:func feature=token type=reader control=sequence
//ff:what TestFindCues — with the rules of articles 42 to 45 and 48 on, the words of their lists in a unit are cues (겸 stands where it ends and shows the letters before it; 중, 님), the units after a Hangul numeral too (개 after 이백삼십), and the groups of a Hangul numeral (삼억오천만); with no rule on none is, though stands is made at every length; a place inside an atom is passed over
package token

import "testing"

func TestFindCues(t *testing.T) {
	on := spacingScan("뚜루뚜겸총무다", SpacingArticles...)
	if len(on.stands) != 8 || !on.stands[4][3] || !on.shows[3][4] || !on.stands[6][4] || !on.shows[4][6] {
		t.Fatal("겸, 총무", on.stands, on.shows)
	}
	if num := spacingScan("이백삼십개", SpacingArticles...); !num.stands[5][4] || !num.shows[4][5] || len(num.stands[2]) != 0 {
		t.Fatal("a unit after a numeral, and no groups with no 만 억 조", num.stands, num.shows)
	}
	if groups := spacingScan("삼억오천만원을", SpacingArticles...); !groups.stands[2][0] || !groups.stands[5][2] {
		t.Fatal("the groups", groups.stands)
	}
	off := spacingScan("뚜루뚜겸총무다")
	if len(off.stands) != 8 || len(off.shows) != 0 || off.stood(3, 4) {
		t.Fatal("no rule on", off.stands, off.shows)
	}
	atom := spacingScan("~/a/b겸총무", SpacingArticles...)
	if insideCount(atom) == 0 || len(atom.shows) == 0 {
		t.Fatal("an atom", insideCount(atom), atom.shows)
	}
}
