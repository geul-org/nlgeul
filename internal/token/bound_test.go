//ff:func feature=token type=rule control=sequence
//ff:what TestBound — two known words with a slash or a hyphen between them are one stem named by the sign's reading, no better founded than the weaker word; a run of Latin letters counts as a word; not when a word is unknown, when the last word does not take what follows, without such a sign, or without a word table
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestBound(t *testing.T) {
	sc := wordsScan("x", layerCost)
	if name, v := sc.bound("입력/출력", rule.TailCase); name != "대비:/" || v <= 0 || v > 1 {
		t.Fatal(name, v)
	}
	if name, v := sc.bound("고양이-마당", rule.TailBare); name != "붙임:-" || v <= 0 {
		t.Fatal(name, v)
	}
	if name, v := sc.bound("abc/출력", rule.TailBare); name != "대비:/" || v != sc.ix.run {
		t.Fatal("a run of Latin letters as a word", name, v)
	}
	if _, v := sc.bound("입력/없는말", rule.TailBare); v != 0 {
		t.Fatal("an unknown word")
	}
	if _, v := sc.bound("입력/출력", rule.TailPred); v != 0 {
		t.Fatal("a noun before a closing")
	}
	if _, v := sc.bound("입력출력", rule.TailBare); v != 0 {
		t.Fatal("no sign")
	}
	if _, v := scanOf("x").bound("입력/출력", rule.TailBare); v != 0 {
		t.Fatal("no word table")
	}
}
