//ff:func feature=token type=rule control=sequence
//ff:what TestComplementRuleWaits — a stream with complementRule reads on to the predicate before it sends the first 이/가 word (the first Next of 학생이 밥을 먹었다 has begun 3 units where the unruled stream has begun fewer), and one with no 이/가 candidate in the sentence (밥을 먹었다) waits for nothing
package token

import "testing"

func TestComplementRuleWaits(t *testing.T) {
	plain := complementStream(t, "학생이 밥을 먹었다", false, false)
	plain.Next()
	ruled := complementStream(t, "학생이 밥을 먹었다", false, true)
	seg, ok := ruled.Next()
	if !ok || segmentLine(seg) != "학생이[0-3]" {
		t.Fatal(segmentLine(seg))
	}
	if begun(ruled) != 3 || begun(plain) >= begun(ruled) {
		t.Fatal("read up to 먹었다", begun(ruled), begun(plain))
	}
	if !ruled.sent[0].judged {
		t.Fatal("judged before it is sent")
	}
	free := complementStream(t, "밥을 먹었다", false, true)
	free.Next()
	if begun(free) != 1 {
		t.Fatal("nothing to wait for", begun(free))
	}
}
