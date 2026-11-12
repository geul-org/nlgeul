//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderBundlesDots — a middle dot at a word's edge or apart with spaces is a sign of a unit's edge named 가운뎃점 (also ㆍ, which Unicode counts as a letter), a bundle of its own; a dot between letters in a row of two is not read as a word of its own; the dots between three words are two words of their own
package token

import (
	"reflect"
	"testing"
)

func TestReaderBundlesDots(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		text  string
		texts []string
		sign  int
	}{{"밥· 국", []string{"밥", "·", "국"}, 1}, {"밥 ㆍ 국", []string{"밥", "ㆍ", "국"}, 1}, {"ㆍ밥", []string{"ㆍ", "밥"}, 0}, {"밥・", []string{"밥", "・"}, 1}} {
		got := r.Bundles(c.text)
		if !reflect.DeepEqual(textsOf(got), c.texts) || kindsOf(got[c.sign])[0] != "가운뎃점:" {
			t.Errorf("%q: %+v", c.text, got)
		}
	}
	for _, br := range r.Bundles("밥··국")[0].Branches {
		if len(br.Words) != 1 {
			t.Fatal("a run of two is not read apart", br)
		}
	}
	three := r.Bundles("곰·솔·별")[0]
	two := false
	for _, br := range three.Branches {
		two = two || len(br.Words) == 5 && br.Words[1].Parts[0].Name.Kind == "가운뎃점" && br.Words[3].Parts[0].Name.Kind == "가운뎃점"
	}
	if !two {
		t.Fatal(three)
	}
}
