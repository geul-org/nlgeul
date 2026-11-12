//ff:func feature=token type=reader control=sequence
//ff:what TestEntryGroup — the lead, the form, the kind and the name, parted by dots; candidates of different names have different keys
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryGroup(t *testing.T) {
	if entryGroup(&entry{form: "이", slot: 1, name: "주격:이"}) != "·이·1·주격:이" || entryGroup(&entry{form: "다", lead: rule.JongN, slot: 0, name: "맺음:다"}) != "ㄴ·다·0·맺음:다" {
		t.Fatal(entryGroup(&entry{form: "이", slot: 1, name: "주격:이"}))
	}
	if entryGroup(&entry{form: "이", slot: 1, name: "주격:이"}) == entryGroup(&entry{form: "이", slot: 1, name: "보격:이"}) {
		t.Fatal("candidates of different names are never filed together")
	}
}
