//ff:func feature=rule type=rule control=sequence
//ff:what TestChainCase — applies when the chain's case shares a flag with the criterion; the evidence is the last particle
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestChainCase(t *testing.T) {
	w := &Word{Chain: Chain{Size: 2, Last: "에서", Kase: CaseSubj | CaseAdv}}
	subj, ev := ChainCase(ctxFor(w, Tail{}), toulmin.Specs{&Bits{Of: "case", Mask: CaseSubj}})
	obj, _ := ChainCase(ctxFor(w, Tail{}), toulmin.Specs{&Bits{Of: "case", Mask: CaseObj}})
	none, _ := ChainCase(ctxFor(&Word{}, Tail{}), toulmin.Specs{&Bits{Of: "case", Mask: CaseSubj}})
	if subj != 1 || ev != "에서" || obj != -1 || none != -1 {
		t.Fatal(subj, ev, obj, none)
	}
}
