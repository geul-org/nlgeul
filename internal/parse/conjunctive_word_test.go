//ff:func feature=parse type=rule control=sequence
//ff:what TestConjunctiveWord — 그리고 and 하지만 are conjunctions; 먹고 is not
package parse

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestConjunctiveWord(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("conjunctions", conjunctions)}
	yes, _ := conjunctiveWord(ctxFor(readWord("그리고", false, &sentence{}), tailFact{}), specs)
	also, _ := conjunctiveWord(ctxFor(readWord("하지만,", false, &sentence{}), tailFact{}), specs)
	no, _ := conjunctiveWord(ctxFor(readWord("먹고", false, &sentence{}), tailFact{}), specs)
	if yes != 1 || also != 1 || no != -1 {
		t.Fatal(yes, also, no)
	}
}
