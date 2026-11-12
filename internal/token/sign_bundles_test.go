//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSignBundles — the signs at one edge of a unit sent up, none dropped: each stretch of one sign in a row is a bundle of its own with the readings that stand (?! is ? and !, ... is one 줄임표, a lead # is 샵 and 태그); stretches no reading stands for are joined into one bundle of letters with no name (->); the end of a reading is 0 before the letters and the letters' end after them
package token

import (
	"reflect"
	"testing"
)

func TestSignBundles(t *testing.T) {
	cases := []struct {
		u     unit
		core  string
		signs string
		at    int
		after bool
		want  []string
	}{
		{unit{text: "뭐?!"}, "뭐", "?!", 1, true, []string{`?: "물음표"/stem=false/depth=0/end=1`, `!: "느낌표"/stem=false/depth=0/end=1`}},
		{unit{text: "뭐..."}, "뭐", "...", 1, true, []string{`...: "줄임표"/stem=false/depth=0/end=1`}},
		{unit{text: "뭐."}, "뭐", ".", 1, true, []string{`.: "마침표"/stem=false/depth=0/end=1`}},
		{unit{text: "#태그"}, "태그", "#", 0, false, []string{`#: "샵"/stem=false/depth=0/end=0; "태그"/stem=false/depth=0/end=0`}},
		{unit{text: "->"}, "", "->", 0, true, []string{`->: ""/stem=true/depth=0/end=0`}},
		{unit{text: "-->"}, "", "-->", 0, true, []string{`-->: ""/stem=true/depth=0/end=0`}},
		{unit{text: "?->!"}, "", "?->!", 0, true, []string{`?: "물음표"/stem=false/depth=0/end=0`, `->: ""/stem=true/depth=0/end=0`, `!: "느낌표"/stem=false/depth=0/end=0`}},
		{unit{text: "뭐"}, "뭐", "", 1, true, nil},
	}
	for _, c := range cases {
		var got []string
		for _, b := range scanOf(c.core).signBundles(c.u, []rune(c.signs), c.at, c.after) {
			got = append(got, b.Text+": "+signBrief(b))
		}
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("%q:\n got %q\nwant %q", c.u.text, got, c.want)
		}
	}
}
