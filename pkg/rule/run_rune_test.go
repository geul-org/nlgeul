//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestRunRune — Latin letters, digits and the underscore join a run; Hangul, a dot and a hyphen do not
package rule

import "testing"

func TestRunRune(t *testing.T) {
	for _, r := range "aZ9_é" {
		if !RunRune(r) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range "가.- " {
		if RunRune(r) {
			t.Errorf("%c", r)
		}
	}
}
