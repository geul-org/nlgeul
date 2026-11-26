//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDropSlot — end to end with the word table of testdata, the copula's stem 이 is read as dropped before a slot (시, 겠, 더) and before a closing that begins with 시 (세요, 시오) after a nominal of the table that ends in a vowel (어디십니까, 세대겠다, 사과더냐, 어디셨다, 어디세요, 어디시오), not after a word no table has (고래겠다), not after a consonant (입력겠다), and not without a word table
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamDropSlot(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	bare, _ := New(Options{})
	for in, want := range map[string]string{
		"어디십니까": "어디[열쇠:체언]+(이)[서술격:이(빠짐)]+시[선어말:시]+(ㅂ)니까[맺음:ㅂ니까]",
		"세대겠다":  "세대[열쇠:체언]+(이)[서술격:이(빠짐)]+겠[선어말:겠]+다[맺음:다]",
		"사과더냐":  "사과[열쇠:체언]+(이)[서술격:이(빠짐)]+더[선어말:더]+냐[맺음:냐]",
		"어디셨다":  "어디[열쇠:체언]+(이)[서술격:이(빠짐)]+시[선어말:시]+(었)[선어말:ㅆ]+다[맺음:다]",
		"어디세요":  "어디[열쇠:체언]+(이)[서술격:이(빠짐)]+시[선어말:시]+(어)[맺음:어(합침)]+요[보조사:요]",
		"어디시오":  "어디[열쇠:체언]+(이)[서술격:이(빠짐)]+시[선어말:시]+오[맺음:오]",
	} {
		if got := streamText(r, in); !strings.Contains(got, "\n"+want+"\n") {
			t.Errorf("%s: no %q\n%s", in, want, got)
		}
		if got := streamText(bare, in); strings.Contains(got, want[strings.Index(want, "+(이)"):]) {
			t.Errorf("%s without a word table\n%s", in, got)
		}
	}
	for _, in := range []string{"고래겠다", "입력겠다"} {
		if got := streamText(r, in); strings.Contains(got, "(이)[서술격:이(빠짐)]+겠") {
			t.Errorf("%s: a dropped stem before 겠\n%s", in, got)
		}
	}
}
