//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTrimUnit — the signs at a unit's edges go, but a sign at a number's front and a unit sign at its end stay; a hyphen before letters, a sign after a run, and a percent sign after a syllable go
package token

import "testing"

func TestTrimUnit(t *testing.T) {
	cases := map[string]string{
		"(곰이)": "곰이", "-3이다.": "-3이다", "$40가": "$40가", "50%": "50%", "(50%)": "50%", "-p": "-p", "-곰": "곰", "a-3": "a-3", "(~/a/b/)": "~/a/b/", "@a": "@a", "a/": "a", "a/b/": "a/b/", "곰%": "곰", "...": "", "3℃,": "3℃",
	}
	for in, want := range cases {
		if got := string(trimUnit(in)); got != want {
			t.Errorf("%s: %s", in, got)
		}
	}
}
