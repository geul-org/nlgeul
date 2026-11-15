//ff:func feature=rule type=helper control=sequence
//ff:what TestListedFactNext — ListNextLast gives the name of the last piece of the word right after the candidate and ListNext2Stem the stem of the word after that; ListPrevKind (not covered elsewhere) the kind of the word before it up to the colon; a candidate with fewer words after it, or a context with no sentence, gives ""
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestListedFactNext(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(KeyAround, Around{
		Self:  Near{Stem: "반장", Last: "보격:이"},
		After: []Near{{Stem: "누구", Last: "주격:가"}, {Stem: "되", Last: "맺음:어"}, {Stem: "먼"}},
	})
	if got := listedFact(ctx, ListNextLast); got != "주격:가" {
		t.Errorf("ListNextLast = %q", got)
	}
	if got := listedFact(ctx, ListNext2Stem); got != "되" {
		t.Errorf("ListNext2Stem = %q", got)
	}
	before := toulmin.NewContext()
	before.Set(KeyAround, Around{Before: []Near{{Kind: "열쇠:체언"}}})
	if listedFact(before, ListPrevKind) != "열쇠" {
		t.Error("ListPrevKind keeps the kind before the colon")
	}
	one := toulmin.NewContext()
	one.Set(KeyAround, Around{After: []Near{{Stem: "되", Last: "맺음:다"}}})
	if listedFact(one, ListNextLast) != "맺음:다" || listedFact(one, ListNext2Stem) != "" {
		t.Error("one word after")
	}
	none := toulmin.NewContext()
	none.Set(KeyAround, Around{Self: Near{Stem: "학생"}})
	if listedFact(none, ListNextLast) != "" || listedFact(none, ListNext2Stem) != "" {
		t.Error("no word after")
	}
	if listedFact(toulmin.NewContext(), ListNextLast) != "" || listedFact(toulmin.NewContext(), ListNext2Stem) != "" {
		t.Error("no sentence")
	}
}
