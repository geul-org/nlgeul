//ff:func feature=rule type=helper control=sequence
//ff:what TestListedFactNextKind — ListNextKind gives the kind of the name of the stem of the word right after the candidate, up to the colon (닫기 of 닫기:1, 열쇠 of 열쇠:체언); a name with no colon is given whole, and no word after it, or a context with no sentence, gives ""
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactNextKind(t *testing.T) {
	shut := toulmin.NewContext()
	shut.Set(KeyAround, Around{Self: Near{Stem: "시"}, After: []Near{{Stem: "]", Kind: "닫기:1"}, {Stem: "책", Kind: "열쇠:체언"}}})
	if got := listedFact(shut, ListNextKind); got != "닫기" {
		t.Errorf("ListNextKind = %q", got)
	}
	noun := toulmin.NewContext()
	noun.Set(KeyAround, Around{After: []Near{{Stem: "책", Kind: "열쇠:체언"}}})
	if got := listedFact(noun, ListNextKind); got != "열쇠" {
		t.Errorf("ListNextKind = %q", got)
	}
	bare := toulmin.NewContext()
	bare.Set(KeyAround, Around{After: []Near{{Stem: "수", Kind: "수"}}})
	if got := listedFact(bare, ListNextKind); got != "수" {
		t.Errorf("a name with no colon: %q", got)
	}
	none := toulmin.NewContext()
	none.Set(KeyAround, Around{Self: Near{Stem: "시"}, Before: []Near{{Kind: "열기:1"}}})
	if listedFact(none, ListNextKind) != "" || listedFact(toulmin.NewContext(), ListNextKind) != "" {
		t.Error("no word after, no sentence")
	}
}
