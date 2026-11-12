//ff:func feature=token type=helper control=sequence
//ff:what TestUnder — a final consonant comes off the syllable it sits under; ㅁ also off ㄻ, leaving ㄹ; not where another consonant, or none, sits; a lead that is no consonant never
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestUnder(t *testing.T) {
	if r, ok := under('간', rule.JongN); !ok || r != '가' {
		t.Fatal(string(r), ok)
	}
	if r, ok := under('삶', rule.JongM); !ok || r != '살' {
		t.Fatal(string(r), ok)
	}
	if r, ok := under('함', rule.JongM); !ok || r != '하' {
		t.Fatal(string(r), ok)
	}
	if _, ok := under('갈', rule.JongN); ok {
		t.Fatal("another consonant")
	}
	if _, ok := under('가', rule.LeadVowel); ok {
		t.Fatal("no consonant")
	}
}
