//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderContestsTypoClosing — the structure of a closing written wrong: the part has the right closing as Text (게), what the letters do not show as Hidden (ㄹ), the letters it is written with as Written (ㄹ께), and only the closing's own name in Names (the mark is not a candidate); a noun written wrong has the right noun as Text and the letters as Written; where the right form stands with verdict 1 the one written wrong stands with 2/(1+forceTypo)-1; printed again by Lines
package token

import (
	"math"
	"reflect"
	"testing"
)

func TestReaderContestsTypoClosing(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bs := r.Contests("할께")
	if len(bs) != 1 {
		t.Fatal(bs)
	}
	first := bs[0].Branches[0]
	parts := first.Words[0].Parts
	if len(parts) != 2 || parts[0].Text != "하" || parts[0].Written != "" {
		t.Fatal(parts)
	}
	end := parts[1]
	if end.Text != "게" || end.Hidden != "ㄹ" || end.Written != "ㄹ께" || end.Stem || end.Name != (Name{"맺음", "ㄹ게"}) {
		t.Fatal(end)
	}
	right := r.Contests("할게")[0].Branches[0]
	if right.Verdict != 1 || math.Abs(first.Verdict-(2/(1+forceTypo)-1)) > 1e-9 || first.Pref != right.Pref {
		t.Fatal(first.Verdict, right.Verdict, first.Pref, right.Pref)
	}
	if got := r.Contests("할게")[0].Branches[0].Words[0].Parts[1]; got.Written != "" || got.Text != "게" || got.Hidden != "ㄹ" {
		t.Fatal("written right", got)
	}
	noun := r.Contests("역활")[0].Branches[0].Words[0].Parts[0]
	if noun.Text != "역할" || noun.Written != "역활" || !noun.Stem || noun.Name != (Name{"열쇠", "체언"}) {
		t.Fatal(noun)
	}
	var lines []string
	for _, b := range bs {
		lines = append(lines, render(b, 0)...)
	}
	if want := r.Lines("할께", false); !reflect.DeepEqual(lines, want) {
		t.Fatalf("printed again\n%q\nLines\n%q", lines, want)
	}
}
