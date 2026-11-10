//ff:func feature=parse type=rule control=sequence
//ff:what TestClosingPredicate — a final or linking closing makes a predicate; a modifying one or no tail does not
package parse

import "testing"

func TestClosingPredicate(t *testing.T) {
	final, ev := closingPredicate(ctxFor(&word{}, tailFact{closing: "다", kind: kindFinal, size: 2}), nil)
	link, _ := closingPredicate(ctxFor(&word{}, tailFact{closing: "고", kind: kindLink, size: 2}), nil)
	adn, _ := closingPredicate(ctxFor(&word{}, tailFact{closing: "는", kind: kindAdn, size: 2}), nil)
	none, _ := closingPredicate(ctxFor(&word{}, tailFact{}), nil)
	if final != 1 || ev != "다" || link != 1 || adn != -1 || none != -1 {
		t.Fatal(final, ev, link, adn, none)
	}
}
