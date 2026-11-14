//ff:func feature=parse type=rule control=sequence
//ff:what TestComplementShare — before 되다 one 이/가 word is surely the complement (+1), each of two is half so (0); no value with another predicate or no 이/가 word
package parse

import "testing"

func TestComplementShare(t *testing.T) {
	one, _ := complementShare(ctxFor(readSentence("곰이 되었다")[0], tailFact{}), nil)
	two, _ := complementShare(ctxFor(readSentence("물이 얼음이 되었다")[0], tailFact{}), nil)
	other, _ := complementShare(ctxFor(readSentence("곰이 먹었다")[0], tailFact{}), nil)
	none, _ := complementShare(ctxFor(readSentence("곰은 되었다")[0], tailFact{}), nil)
	if one != 1 || two != 0 || other != -1 || none != -1 {
		t.Fatal(one, two, other, none)
	}
}
