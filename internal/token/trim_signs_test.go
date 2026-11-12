//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestTrimSigns — punctuation, symbols and the named signs (ㆍ, a letter for Unicode) are taken off the two ends of a field, and never from between letters
package token

import "testing"

func TestTrimSigns(t *testing.T) {
	for in, want := range map[string]string{"ㆍ밥ㆍ": "밥", "·밥·": "밥", "(밥).": "밥", "밥ㆍ국": "밥ㆍ국", "ㆍ": "", "": "", "밥": "밥"} {
		if got := string(trimSigns(in)); got != want {
			t.Errorf("%q: %q, want %q", in, got, want)
		}
	}
}
