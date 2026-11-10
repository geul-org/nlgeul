//ff:func feature=parse type=rule control=sequence
//ff:what TestTopicObject — 이것은 carries no case particle; 집에는 and 매우 do not qualify
package parse

import "testing"

func TestTopicObject(t *testing.T) {
	yes, _ := topicObject(ctxFor(readWord("이것은", false, &sentence{}), tailFact{}), nil)
	kase, _ := topicObject(ctxFor(readWord("집에는", false, &sentence{}), tailFact{}), nil)
	none, _ := topicObject(ctxFor(readWord("매우", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || kase != -1 || none != -1 {
		t.Fatal(yes, kase, none)
	}
}
