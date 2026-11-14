//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestEntriesAt — the entries under a token and under its keys as a fused token, the same whether the index has worked them out (addEnds) or not
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
	worked := len(sc.entriesAt(1))
	sc.ix.ends = nil
	if worked == 0 || len(sc.entriesAt(1)) != worked {
		t.Fatal("an index that has not worked them out gives the same", worked)
	}
}
