//ff:func feature=token type=rule control=sequence
//ff:what TestBound — two known words with a slash or a hyphen between them are one stem named by the sign's candidate, no better founded than the weaker word; a run of Latin letters counts as a word; not when a word is unknown, when the last word does not take what follows, without such a sign, or without a word table
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestBound(t *testing.T) {
	sc := wordsScan("x", layerCost)
	if r := sc.bound("입력/출력", rule.TailCase); len(r) != 1 || r[0].name != "대비:/" || r[0].v <= 0 || r[0].v > 1 {
		t.Fatal(r)
	}
	if r := sc.bound("고양이-마당", rule.TailBare); len(r) != 1 || r[0].name != "붙임:-" || r[0].v <= 0 {
		t.Fatal(r)
	}
	if r := sc.bound("abc/출력", rule.TailBare); len(r) != 1 || r[0].name != "대비:/" || r[0].v != sc.ix.run {
		t.Fatal("a run of Latin letters as a word", r)
	}
	if r := sc.bound("입력/없는말", rule.TailBare); len(r) != 0 {
		t.Fatal("an unknown word")
	}
	if r := sc.bound("입력/출력", rule.TailPred); len(r) != 0 {
		t.Fatal("a noun before a closing")
	}
	if r := sc.bound("입력출력", rule.TailBare); len(r) != 0 {
		t.Fatal("no sign")
	}
	if r := scanOf("x").bound("입력/출력", rule.TailBare); len(r) != 0 {
		t.Fatal("no word table")
	}
}
