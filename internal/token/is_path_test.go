//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestIsPath — a path's head makes a path; so do two slashes between names; one slash, slashes between Hangul words, numbers only, and a text that begins with Hangul do not
package token

import "testing"

func TestIsPath(t *testing.T) {
	for _, yes := range []string{"/a", "~/a", "./a", "../a", "a/b/c", "a.b/c-d/e_f"} {
		if !isPath([]rune(yes)) {
			t.Errorf("%s is a path", yes)
		}
	}
	for _, no := range []string{"a/b", "가/나/다", "1/2/3", "가/a/b", "a//b", "a"} {
		if isPath([]rune(no)) {
			t.Errorf("%s is no path", no)
		}
	}
}
