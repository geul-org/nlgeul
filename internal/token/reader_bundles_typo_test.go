//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBundlesTypo — the structure of a stem written wrong: its stem part has the right form as Text, the form it is written in as Written, and only the word's own kind in Names (the mark is not a reading); the same stem written right has no Written; the bundles printed again are the lines Lines gives
package token

import (
	"reflect"
	"testing"
)

func TestReaderBundlesTypo(t *testing.T) {
	r, err := New(Options{Words: "testdata/words-typo.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	var lines []string
	bs := r.Bundles("맞겨")
	for _, b := range bs {
		lines = append(lines, render(b, 0)...)
	}
	if want := r.Lines("맞겨", false); !reflect.DeepEqual(lines, want) {
		t.Fatalf("printed again\n%q\nLines\n%q", lines, want)
	}
	if len(bs) != 1 || len(bs[0].Branches) < 2 {
		t.Fatal(bs)
	}
	first := bs[0].Branches[0].Words[0]
	stem := first.Parts[0]
	if stem.Text != "맡기" || stem.Written != "맞기" || !stem.Stem || stem.Name != (Name{"열쇠", "동사"}) {
		t.Fatal(stem)
	}
	if first.Parts[1].Written != "" || first.Parts[1].Text != "어" && first.Parts[1].Hidden != "어" {
		t.Fatal(first.Parts[1])
	}
	if right := r.Bundles("맡겨")[0].Branches[0].Words[0].Parts[0]; right.Text != "맡기" || right.Written != "" {
		t.Fatal("written right", right)
	}
}
