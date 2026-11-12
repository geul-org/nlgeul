//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestUnitsOf — a sentence is cut from its first letter on at spaces and at the signs of a pair: an opening bracket alone, a closing one with the particle written onto it, a pair that spans a space, a span of code with its spaces as one unit, a sign with no partner as an opening sign, a sign not read as a sign of a pair (x < y) in its word; each sign carries its candidates with depths and whether it is also its letter; an empty sentence has no units
package token

import (
	"strings"
	"testing"
)

func TestUnitsOf(t *testing.T) {
	ix, _ := newIndex(Options{})
	brief := func(sentence string) string {
		var out []string
		for _, u := range unitsOf(ix, sentence, false) {
			s := u.text
			if u.name != "" {
				s += "/" + u.name
			}
			for _, rd := range u.candidates {
				s += "/" + rd.name
				s += string(rune('0' + rd.depth))
			}
			if u.bare {
				s += "/글자"
			}
			if u.first {
				s += "/첫"
			}
			if u.spaced {
				s += "/빈칸"
			}
			out = append(out, s)
		}
		return strings.Join(out, " ")
	}
	for in, want := range map[string]string{
		"곰(큰 곰)이 `a b`를 본다": "곰/첫 (/묶음/열기1 큰/빈칸 곰 )이/묶음/닫기1/빈칸 `a b`를/코드/빈칸 본다",
		"곰(이 \"간다":          "곰/첫 (/묶음/열기1 이/빈칸 \"/묶음/열기2 간다",
		"「책」을":              "「/묶음/열기1/첫 책 」을/묶음/닫기1",
		"x < y":             "x/첫/빈칸 </빈칸 y",
		"'I'm (O'Brien)'":   "'/묶음/열기1/글자/첫 I'm/빈칸 (/묶음/열기1/열기2 O'Brien )/묶음/닫기1/닫기2 '/묶음/열기1/닫기1/글자",
		"5'10\"":            "5'10/첫 \"/묶음/열기1",
		"첫줄\n(둘째":           "첫줄/첫/빈칸 (/묶음/열기1/첫 둘째",
		"  ":                "",
		"":                  "",
	} {
		if got := brief(in); got != want {
			t.Errorf("%q:\n got %s\nwant %s", in, got, want)
		}
	}
}
