//ff:func feature=token type=reader control=sequence
//ff:what TestEntryGroup — the lead, the form and the kind, parted by dots
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryGroup(t *testing.T) {
	if entryGroup(&entry{form: "이", slot: 1}) != "·이·1" || entryGroup(&entry{form: "다", lead: rule.JongN, slot: 0}) != "ㄴ·다·0" {
		t.Fatal(entryGroup(&entry{form: "이", slot: 1}))
	}
}
