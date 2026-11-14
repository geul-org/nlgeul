//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBareSign — a contest of one candidate of one word of one part with no name is signs sent up as their letters; a named sign, a contest of two candidates, a word of two parts, and a contest with no branch are not
package token

import "testing"

func TestBareSign(t *testing.T) {
	bare := signBranch(Part{Text: "->", Stem: true}, 0)
	named := signBranch(Part{Text: "?", Name: Name{Kind: "물음표"}}, 0)
	two := Branch{Words: []Eojeol{{Parts: []Part{{Text: "a"}, {Text: "b"}}}}, Verdict: 1}
	cases := []struct {
		b    Contest
		want bool
	}{
		{Contest{Text: "->", Branches: []Branch{bare}}, true},
		{Contest{Text: "?", Branches: []Branch{named}}, false},
		{Contest{Text: "#", Branches: []Branch{named, bare}}, false},
		{Contest{Text: "ab", Branches: []Branch{two}}, false},
		{Contest{Text: "x", Branches: []Branch{{}}}, false},
		{Contest{Text: "x"}, false},
	}
	for _, c := range cases {
		if got := bareSign(c.b); got != c.want {
			t.Errorf("%+v: %v, want %v", c.b, got, c.want)
		}
	}
}
