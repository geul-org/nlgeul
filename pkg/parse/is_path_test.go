//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestIsPath — a path's head makes a path; so do two slashes between names, and a slash between names with a slash that ends the last name (pkg/parser/); one slash (a name and its end slash: files/), slashes between Hangul words, numbers only, and a text that begins with Hangul do not
package parse

import "testing"

func TestIsPath(t *testing.T) {
	for _, yes := range []string{"/a", "~/a", "./a", "../a", "a/b/c", "a.b/c-d/e_f", "pkg/parser/", "a/b/"} {
		if !isPath([]rune(yes)) {
			t.Errorf("%s is a path", yes)
		}
	}
	for _, no := range []string{"a/b", "가/나/다", "1/2/3", "가/a/b", "a//b", "a", "files/", "a/", "가/나/"} {
		if isPath([]rune(no)) {
			t.Errorf("%s is no path", no)
		}
	}
}
