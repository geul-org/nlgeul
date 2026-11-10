//ff:func feature=parse type=rule control=sequence
//ff:what TestBeforeBound — applies to a word with a modifying reading before a dependent noun; not before another word, not to a word without such a reading
package parse

import "testing"

func TestBeforeBound(t *testing.T) {
	words := readSentence("없을 때까지 사과가 것을 먹는 사과")
	yes, _ := beforeBound(ctxFor(words[0], tailFact{}), nil)
	plain, _ := beforeBound(ctxFor(words[2], tailFact{}), nil)
	other, _ := beforeBound(ctxFor(words[4], tailFact{}), nil)
	if yes != 1 || plain != -1 || other != -1 {
		t.Fatal(yes, plain, other)
	}
}
