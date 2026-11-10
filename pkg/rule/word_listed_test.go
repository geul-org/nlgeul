//ff:func feature=rule type=rule control=sequence
//ff:what TestWordListed — applies to a word of the list only; the evidence is the word
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWordListed(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("words", []string{"그리고", "이"})}
	yes, ev := WordListed(ctxFor(&Word{Text: "그리고"}, Tail{}), specs)
	no, _ := WordListed(ctxFor(&Word{Text: "먹고"}, Tail{}), specs)
	if yes != 1 || ev != "그리고" || no != -1 {
		t.Fatal(yes, ev, no)
	}
}
