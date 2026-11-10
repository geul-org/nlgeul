//ff:func feature=parse type=rule control=sequence
//ff:what TestRangeAdverbial — 때까지, 처음부터 and 어제까지는 may be adverbials of range; 어제까지의 (a case particle) and 이것은 are not
package parse

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestRangeAdverbial(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("rangeParticles", rangeParticles)}
	yes, _ := rangeAdverbial(ctxFor(readWord("때까지", false, &sentence{}), tailFact{}), specs)
	from, _ := rangeAdverbial(ctxFor(readWord("처음부터", false, &sentence{}), tailFact{}), specs)
	inner, _ := rangeAdverbial(ctxFor(readWord("어제까지는", false, &sentence{}), tailFact{}), specs)
	kase, _ := rangeAdverbial(ctxFor(readWord("어제까지의", false, &sentence{}), tailFact{}), specs)
	no, _ := rangeAdverbial(ctxFor(readWord("이것은", false, &sentence{}), tailFact{}), specs)
	if yes != 1 || from != 1 || inner != 1 || kase != -1 || no != -1 {
		t.Fatal(yes, from, inner, kase, no)
	}
}
