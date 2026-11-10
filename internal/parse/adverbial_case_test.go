//ff:func feature=parse type=rule control=sequence
//ff:what TestAdverbialCase — 지침으로 and 집에는 are adverbial; 학교에서의 and 이것은 are not
package parse

import "testing"

func TestAdverbialCase(t *testing.T) {
	yes, _ := adverbialCase(ctxFor(readWord("지침으로", false, &sentence{}), tailFact{}), nil)
	aux, _ := adverbialCase(ctxFor(readWord("집에는", false, &sentence{}), tailFact{}), nil)
	adn, _ := adverbialCase(ctxFor(readWord("학교에서의", false, &sentence{}), tailFact{}), nil)
	no, _ := adverbialCase(ctxFor(readWord("이것은", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || aux != 1 || adn != -1 || no != -1 {
		t.Fatal(yes, aux, adn, no)
	}
}
