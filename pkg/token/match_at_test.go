//ff:func feature=token type=helper control=sequence
//ff:what TestMatchAt — a word that stands in the letters from a place ends where its letters end; one that differs, or runs past the end, does not stand; a word of no letters ends where it begins
package token

import "testing"

func TestMatchAt(t *testing.T) {
	s := []rune("곰이 물을")
	if q := matchAt(s, 0, "곰이"); q != 2 {
		t.Fatal(q)
	}
	if q := matchAt(s, 3, "물을"); q != 5 {
		t.Fatal(q)
	}
	if q := matchAt(s, 0, "곰을"); q != -1 {
		t.Fatal("differs", q)
	}
	if q := matchAt(s, 3, "물을요"); q != -1 {
		t.Fatal("runs past the end", q)
	}
	if q := matchAt(s, 5, ""); q != 5 {
		t.Fatal(q)
	}
}
