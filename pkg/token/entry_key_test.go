//ff:func feature=token type=reader control=sequence
//ff:what TestEntryKey — the last token of the form; the lead's key for a piece with no token
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryKey(t *testing.T) {
	if entryKey(&entry{form: "에서"}) != '서' || entryKey(&entry{lead: rule.LeadVowel}) != keyVowel || entryKey(&entry{lead: rule.JongN}) != 'ㄴ' {
		t.Fatal("entryKey")
	}
}
