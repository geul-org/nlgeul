//ff:func feature=token type=reader control=sequence
//ff:what TestAuxLinkedRuleTrace — the trace of 먹어 있다 has 통과 for the auxiliary 있 under 문장:연결 꼴 뒤 보조 용언 and 기각 with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝난다 for the main adjective 있 under 문장:연결 꼴 뒤가 아닌 본용언; that of 집에 있다 has 기각 with the reason 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다 for the auxiliary and 통과 for the main adjective; each rule writes one line for a candidate it speaks of and none for 있[열쇠:동사] under the first rule
package token

import "testing"

func TestAuxLinkedRuleTrace(t *testing.T) {
	_, linked := auxLinkedLast(t, "먹어 있다", true, true)
	auxLinkedHas(t, linked,
		"  · 있[보조 용언:있]+다[맺음:다]\t문장:연결 꼴 뒤 보조 용언\t1.00\t통과",
		"  · 있[열쇠:형용사]+다[맺음:다]\t문장:연결 꼴 뒤가 아닌 본용언\t-1.00\t기각: 앞 어절이 연결 꼴(어/아, 고)로 끝난다")
	_, apart := auxLinkedLast(t, "집에 있다", true, true)
	auxLinkedHas(t, apart,
		"  · 있[보조 용언:있]+다[맺음:다]\t문장:연결 꼴 뒤 보조 용언\t-1.00\t기각: 앞 어절이 연결 꼴(어/아, 고)로 끝나지 않는다",
		"  · 있[열쇠:형용사]+다[맺음:다]\t문장:연결 꼴 뒤가 아닌 본용언\t1.00\t통과")
	if auxLinkedCount(linked) != 2 || auxLinkedCount(apart) != 2 {
		t.Fatalf("one line for each candidate a rule speaks of: %q %q", linked, apart)
	}
}
