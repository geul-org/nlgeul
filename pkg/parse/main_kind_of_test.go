//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestMainKindOf — 있 is read as the adjective (열쇠:형용사) and every other stem of an auxiliary verb (보, 주, 두, 놓, 버리, 지) as the verb (열쇠:동사)
package parse

import "testing"

func TestMainKindOf(t *testing.T) {
	if got := mainKindOf("있"); got != "열쇠:형용사" {
		t.Fatal(got)
	}
	for _, s := range []string{"보", "주", "두", "놓", "버리", "지"} {
		if got := mainKindOf(s); got != "열쇠:동사" {
			t.Fatal(s, got)
		}
	}
	r := mainUnlinkedRule([]string{"있"})
	brs, _ := auxLinkedLast(t, "집에 있다", true, false)
	verb := auxLinkedPick(t, brs, "있[열쇠:동사]+다[맺음:다]")
	if r.about(nearOf(verb)) {
		t.Fatal("있[열쇠:동사] gets no preference")
	}
}
