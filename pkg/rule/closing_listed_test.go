//ff:func feature=rule type=rule control=sequence
//ff:what TestClosingListed — applies to a closing of the list only
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClosingListed(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("adverbial", []string{"게"})}
	yes, _ := ClosingListed(ctxFor(&Word{}, Tail{Closing: "게", Size: 2}), specs)
	no, _ := ClosingListed(ctxFor(&Word{}, Tail{Closing: "고", Size: 2}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
