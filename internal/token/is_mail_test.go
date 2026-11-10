//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestIsMail — one @ between names with a dot after it is a mail address; no @, two, nothing before it, no dot after it, a dot or @ at the end, or another sign is none
package token

import "testing"

func TestIsMail(t *testing.T) {
	if !isMail([]rune("a.b@c-d.org")) {
		t.Fatal("a mail address")
	}
	for _, no := range []string{"a.b", "a@b@c.d", "@b.c", "a@bc", "a@b.", "a.b@", "a@b/c.d"} {
		if isMail([]rune(no)) {
			t.Errorf("%s is no mail address", no)
		}
	}
}
