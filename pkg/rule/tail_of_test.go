//ff:func feature=rule type=helper control=sequence
//ff:what TestTailOf — the tail set in the context comes back; without one, an unread tail
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestTailOf(t *testing.T) {
	if TailOf(toulmin.NewContext()).Size != 0 || TailOf(ctxFor(&Word{}, Tail{Size: 4})).Size != 4 {
		t.Fatal("TailOf")
	}
}
