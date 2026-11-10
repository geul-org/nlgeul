//ff:func feature=rule type=helper control=sequence
//ff:what TestMarkOf — the sign set in the context comes back; without one, an empty sign
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestMarkOf(t *testing.T) {
	if MarkOf(toulmin.NewContext()).Run != 0 || MarkOf(markCtx(Mark{Run: 3})).Run != 3 {
		t.Fatal("MarkOf")
	}
}
