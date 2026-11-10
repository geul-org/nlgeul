//ff:func feature=rule type=rule control=sequence
//ff:what TestChainNoCase — applies to a chain without a case particle; not to one with a case, not without a chain
package rule

import "testing"

func TestChainNoCase(t *testing.T) {
	yes, _ := ChainNoCase(ctxFor(&Word{Chain: Chain{Size: 2}}, Tail{}), nil)
	kase, _ := ChainNoCase(ctxFor(&Word{Chain: Chain{Size: 2, Kase: CaseAdv}}, Tail{}), nil)
	bare, _ := ChainNoCase(ctxFor(&Word{Chain: Chain{Size: 4, Kase: CaseBare}}, Tail{}), nil)
	none, _ := ChainNoCase(ctxFor(&Word{}, Tail{}), nil)
	if yes != 1 || kase != -1 || bare != -1 || none != -1 {
		t.Fatal(yes, kase, bare, none)
	}
}
