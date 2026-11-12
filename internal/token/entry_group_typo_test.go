//ff:func feature=token type=reader control=sequence
//ff:what TestEntryGroupTypo — the key of a closing written wrong ends with the letters it is written with, so that it is not the key of the right closing of the same form and lead
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntryGroupTypo(t *testing.T) {
	right := &entry{name: "맺음:에요", form: "에요"}
	wrong := &entry{name: "맺음:에요", form: "에요", wrote: "에요", right: "예요"}
	if entryGroup(right) != "·에요·0·맺음:에요" || entryGroup(wrong) != "·에요·0·오표기:에요·맺음:에요" || entryGroup(right) == entryGroup(wrong) {
		t.Fatal(entryGroup(right), entryGroup(wrong))
	}
	if got := entryGroup(&entry{form: "께", lead: rule.JongL, wrote: "ㄹ께", name: "맺음:께"}); got != "ㄹ·께·0·오표기:ㄹ께·맺음:께" {
		t.Fatal(got)
	}
}
