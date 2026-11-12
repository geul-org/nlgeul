//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEntriesAt — the entries under a token and under its keys as a fused token
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestEntriesAt(t *testing.T) {
	sc := scanOf("간")
	var plain, fused bool
	for _, e := range sc.entriesAt(1) {
		plain = plain || e.lead == 0
		fused = fused || e.lead == rule.JongN
	}
	if !fused || plain {
		t.Fatal(plain, fused)
	}
}
