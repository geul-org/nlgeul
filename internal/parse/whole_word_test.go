//ff:func feature=parse type=rule control=sequence
//ff:what TestWholeWord — applies to a word of the list only
package parse

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWholeWord(t *testing.T) {
	specs := toulmin.Specs{toulmin.NewStringListSpec("wholeWords", conjunctions)}
	yes, _ := wholeWord(ctxFor(readWord("그리고", false, &sentence{}), tailFact{}), specs)
	no, _ := wholeWord(ctxFor(readWord("먹고", false, &sentence{}), tailFact{}), specs)
	if yes != 1 || no != -1 {
		t.Fatal(yes, no)
	}
}
