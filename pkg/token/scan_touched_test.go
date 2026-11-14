//ff:func feature=token type=reader control=sequence
//ff:what TestScanTouched — at the place where a unit headed by a sign of a pair begins, the sign is given once for each way it is read and once more as its letter; anywhere else (head not set) the candidates stay as they are; the order is kept
package token

import "testing"

func TestScanTouched(t *testing.T) {
	sc, _, _ := netUnit(t, "", "가")
	head := Branch{Words: []Eojeol{{Parts: []Part{{Text: "(", Name: Name{Kind: pairName}, Stem: true}}, End: 1}}, Verdict: 1, Pref: -5}
	other := Branch{Words: []Eojeol{{Parts: []Part{{Text: "(a", Stem: true}}, End: 2}}, Verdict: 0.3, Pref: -9}
	u := unit{text: "(", head: 1, name: pairName, candidates: []pairing{{openName, 1}, {openName, 2}}, bare: true}
	got := sc.touched([]Branch{head, other}, u, true)
	if len(got) != 4 || got[0].Words[0].Parts[0].Name.Form != "1" || got[1].Words[0].Parts[0].Name.Form != "2" || got[2].Words[0].Parts[0].Name != (Name{}) || got[3].String() != other.String() {
		t.Fatal(got)
	}
	if kept := sc.touched([]Branch{head, other}, u, false); len(kept) != 2 || kept[0].Words[0].Parts[0].Name.Kind != pairName {
		t.Fatal("not at the head", kept)
	}
}
