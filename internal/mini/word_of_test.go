//ff:func feature=mini type=helper control=sequence
//ff:what TestWordOf — the word set in the context comes back; without one, an empty word in an empty sentence
package mini

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWordOf(t *testing.T) {
	if w := wordOf(toulmin.NewContext()); w.chain.size != 0 || w.sent == nil {
		t.Fatal("empty context")
	}
	w := &word{text: "가"}
	if wordOf(ctxFor(w, tailFact{})) != w {
		t.Fatal("word")
	}
}
