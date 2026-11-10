//ff:func feature=parse type=reader control=sequence
//ff:what TestMatchPlain — 습니다 after 좋 is three syllables; a word that does not end with it gives nothing
package parse

import "testing"

func TestMatchPlain(t *testing.T) {
	e := ending{form: "습니다", after: afterBatchim, kind: kindFinal}
	if got, ok := matchPlain([]rune("좋습니다"), 4, e); !ok || got.size != 6 || got.stem != "좋" || got.marked != "좋+습+니다" {
		t.Fatal(got, ok)
	}
	if _, ok := matchPlain([]rune("좋았다"), 3, e); ok {
		t.Fatal("좋았다")
	}
}
