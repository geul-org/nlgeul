//ff:func feature=token type=helper control=sequence
//ff:what TestUnder — a final consonant comes off the syllable it sits under; ㅁ also off ㄻ, leaving ㄹ; not where another consonant, or none, sits; a lead that is no consonant never
package token

import "testing"

func TestUnder(t *testing.T) {
	if r, ok := under('간', jongN); !ok || r != '가' {
		t.Fatal(string(r), ok)
	}
	if r, ok := under('삶', jongM); !ok || r != '살' {
		t.Fatal(string(r), ok)
	}
	if r, ok := under('함', jongM); !ok || r != '하' {
		t.Fatal(string(r), ok)
	}
	if _, ok := under('갈', jongN); ok {
		t.Fatal("another consonant")
	}
	if _, ok := under('가', leadVowel); ok {
		t.Fatal("no consonant")
	}
}
