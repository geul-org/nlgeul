//ff:func feature=token type=reader control=sequence
//ff:what TestSideOf — a run of letters, a run of digits, a Hangul syllable, and anything else
package token

import (
	"testing"

	"github.com/geul-org/koparse/pkg/rule"
)

func TestSideOf(t *testing.T) {
	if sideOf("task") != rule.SideAlnum || sideOf("14") != rule.SideAlnum|rule.SideDigit || sideOf("다") != rule.SideHangul || sideOf(".") != 0 || sideOf("") != 0 {
		t.Fatal("sideOf")
	}
}
