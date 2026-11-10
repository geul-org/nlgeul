//ff:func feature=mini type=helper control=sequence
//ff:what TestTailOf — the tail set in the context comes back; without one, an unread tail
package mini

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestTailOf(t *testing.T) {
	if tailOf(toulmin.NewContext()).size != 0 {
		t.Fatal("empty context")
	}
	if tailOf(ctxFor(&word{}, tailFact{size: 4})).size != 4 {
		t.Fatal("tail")
	}
}
