//ff:func feature=token type=reader control=sequence
//ff:what TestJoinChainHyphen — a name joined by a hyphen and a dot is one atom whose dot's reading comes before 붙임 (geul-goal.md: 파일명 then 붙임); a name joined by a hyphen alone is 붙임
package token

import "testing"

func TestJoinChainHyphen(t *testing.T) {
	sc := scanOf("geul-goal.md")
	if a := sc.atoms[0]; a.end != 12 || atomNames(a) != "파일명,붙임" {
		t.Fatal(a)
	}
	sc = scanOf("h-Cat")
	if a := sc.atoms[0]; a.end != 5 || atomNames(a) != "붙임" {
		t.Fatal(a)
	}
}
