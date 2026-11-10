//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSpecial — an address, a mail address, a path and an option are told and end before a particle; a word, a number with a slash, two Hangul words with a slash, and a lone dash are none
package token

import "testing"

func TestSpecial(t *testing.T) {
	cases := []struct {
		in, name string
		end      int
	}{
		{"https://a.org/x에", "주소", 15}, {"a@b.com으로", "전자우편", 7}, {"~/a/b를", "경로", 5}, {"./가/나.md는", "경로", 8},
		{"a/b/c", "경로", 5}, {"--words", "옵션", 7}, {"-v를", "옵션", 2}, {"-v=3", "옵션", 4},
		{"고양이가", "", 0}, {"11/10", "", 0}, {"2026/11/10", "", 0}, {"가/나/다", "", 0}, {"a/b", "", 0}, {"-", "", 0}, {"-3", "", 0}, {"---x", "", 0}, {"a@b", "", 0}, {"a@b.", "", 0}, {"@a.b", "", 0}, {"a@b@c.d", "", 0}, {"a@b!c.d", "", 0}, {"가/a/b", "", 0},
	}
	for _, c := range cases {
		if name, end := special([]rune(c.in)); name != c.name || end != c.end {
			t.Errorf("%s: %s %d", c.in, name, end)
		}
	}
}
