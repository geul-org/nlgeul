//ff:func feature=mini type=rule control=sequence
//ff:what TestComplementCase — 곰이 is a complement before 아니다 and 되었다, not before 먹었다; 곰을 never
package mini

import "testing"

func TestComplementCase(t *testing.T) {
	words := readSentence("곰이 곰을 아니다")
	yes, _ := complementCase(ctxFor(words[0], tailFact{}), nil)
	obj, _ := complementCase(ctxFor(words[1], tailFact{}), nil)
	become, _ := complementCase(ctxFor(readSentence("곰이 되었다")[0], tailFact{}), nil)
	no, _ := complementCase(ctxFor(readSentence("곰이 먹었다")[0], tailFact{}), nil)
	if yes != 1 || obj != -1 || become != 1 || no != -1 {
		t.Fatal(yes, obj, become, no)
	}
}
