//ff:func feature=rule type=rule control=sequence
//ff:what TestClosingKind — applies when the closing's kind shares a flag with the criterion; not without a tail
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClosingKind(t *testing.T) {
	pred := toulmin.Specs{&Bits{Of: "kind", Mask: KindFinal | KindLink}}
	yes, ev := ClosingKind(ctxFor(&Word{}, Tail{Closing: "고", Kind: KindLink, Size: 2}), pred)
	adn, _ := ClosingKind(ctxFor(&Word{}, Tail{Closing: "는", Kind: KindAdn, Size: 2}), pred)
	none, _ := ClosingKind(ctxFor(&Word{}, Tail{}), pred)
	if yes != 1 || ev != "고" || adn != -1 || none != -1 {
		t.Fatal(yes, ev, adn, none)
	}
}
