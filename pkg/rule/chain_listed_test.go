//ff:func feature=rule type=rule control=sequence
//ff:what TestChainListed — applies when a listed particle is in a chain without a case particle
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestChainListed(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("range", []string{"까지", "부터"})}
	yes, _ := ChainListed(ctxFor(&Word{Chain: Chain{Size: 6, Parts: "까지+는"}}, Tail{}), specs)
	kase, _ := ChainListed(ctxFor(&Word{Chain: Chain{Size: 6, Parts: "까지+의", Kase: CaseAdn}}, Tail{}), specs)
	other, _ := ChainListed(ctxFor(&Word{Chain: Chain{Size: 2, Parts: "은"}}, Tail{}), specs)
	if yes != 1 || kase != -1 || other != -1 {
		t.Fatal(yes, kase, other)
	}
}
