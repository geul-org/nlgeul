//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBareSign — a bundle of one reading of one word of one part with no name is signs sent up as their letters; a named sign, a contest, a word of two parts, and a bundle with no branch are not
package token

import "testing"

func TestBareSign(t *testing.T) {
	bare := signBranch(Part{Text: "->", Stem: true}, 0)
	named := signBranch(Part{Text: "?", Name: Name{Kind: "물음표"}}, 0)
	two := Branch{Words: []Eojeol{{Parts: []Part{{Text: "a"}, {Text: "b"}}}}, Verdict: 1}
	cases := []struct {
		b    Bundle
		want bool
	}{
		{Bundle{Text: "->", Branches: []Branch{bare}}, true},
		{Bundle{Text: "?", Branches: []Branch{named}}, false},
		{Bundle{Text: "#", Branches: []Branch{named, bare}}, false},
		{Bundle{Text: "ab", Branches: []Branch{two}}, false},
		{Bundle{Text: "x", Branches: []Branch{{}}}, false},
		{Bundle{Text: "x"}, false},
	}
	for _, c := range cases {
		if got := bareSign(c.b); got != c.want {
			t.Errorf("%+v: %v, want %v", c.b, got, c.want)
		}
	}
}
