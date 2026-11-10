//ff:func feature=parse type=rule control=sequence
//ff:what TestTopicSubject — 이것은 and 아이들 carry no case particle; 집에는 and 매우 do not qualify
package parse

import "testing"

func TestTopicSubject(t *testing.T) {
	yes, _ := topicSubject(ctxFor(readWord("이것은", false, &sentence{}), tailFact{}), nil)
	suffix, _ := topicSubject(ctxFor(readWord("아이들", false, &sentence{}), tailFact{}), nil)
	kase, _ := topicSubject(ctxFor(readWord("집에는", false, &sentence{}), tailFact{}), nil)
	none, _ := topicSubject(ctxFor(readWord("매우", false, &sentence{}), tailFact{}), nil)
	if yes != 1 || suffix != 1 || kase != -1 || none != -1 {
		t.Fatal(yes, suffix, kase, none)
	}
}
