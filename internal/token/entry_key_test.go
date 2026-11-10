//ff:func feature=token type=reader control=sequence
//ff:what TestEntryKey — the last token of the form; the lead's key for a piece with no token
package token

import "testing"

func TestEntryKey(t *testing.T) {
	if entryKey(&entry{form: "에서"}) != '서' || entryKey(&entry{lead: leadVowel}) != keyVowel || entryKey(&entry{lead: jongN}) != 'ㄴ' {
		t.Fatal("entryKey")
	}
}
