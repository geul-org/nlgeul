//ff:func feature=rule type=helper control=sequence
//ff:what TestListedFactSelfText — ListSelfText gives the letters of the word of the candidate judged in its sentence (한 of 하다 with ㄴ), not its stem, and Listed finds them in its list; a candidate with no letters given, or a context with no sentence, gives ""
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactSelfText(t *testing.T) {
	one := toulmin.NewContext()
	one.Set(KeyAround, Around{Self: Near{Text: "한", Stem: "하", Kind: "열쇠:동사", Last: "꾸밈:ㄴ"}, After: []Near{{Stem: "번"}}})
	if got := listedFact(one, ListSelfText); got != "한" {
		t.Errorf("ListSelfText = %q", got)
	}
	if v, ev := Listed(one, toulmin.Specs{&List{Name: "수 관형사", Of: ListSelfText, Items: []string{"한", "두"}}}); v != 1 || ev != "한" {
		t.Errorf("Listed on ListSelfText = %v, %v", v, ev)
	}
	bare := toulmin.NewContext()
	bare.Set(KeyAround, Around{Self: Near{Stem: "하"}})
	if listedFact(bare, ListSelfText) != "" || listedFact(toulmin.NewContext(), ListSelfText) != "" {
		t.Error("no letters, no sentence")
	}
}
