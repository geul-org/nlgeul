//ff:func feature=rule type=helper control=sequence
//ff:what TestWordOf — the word set in the context comes back; without one, an empty word in an empty sentence
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWordOf(t *testing.T) {
	if w := WordOf(toulmin.NewContext()); w.Chain.Size != 0 || w.Sent == nil {
		t.Fatal("empty context")
	}
	w := &Word{Text: "가"}
	if WordOf(ctxFor(w, Tail{})) != w {
		t.Fatal("word")
	}
}
