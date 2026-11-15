//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestAroundIs — applies when the candidate and the words around it have one of the facts asked for (its word is glued to the word before, the word before ends in a nominal, an adnominal may stand before the word after); not when they have none of them, when there is no word before or after, or when the context holds no candidate; Not turns both round
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAroundIs(t *testing.T) {
	glued := Near{Stem: "자", Glued: true}
	noun := Near{Stem: "패소", Nominal: true}
	head := Near{Stem: "및"}
	cases := []struct {
		name string
		a    Around
		mask int
		not  bool
		want float64
	}{
		{"glued", Around{Self: glued}, AroundGlued, false, 1},
		{"spaced", Around{Self: Near{Stem: "자"}}, AroundGlued, false, -1},
		{"after a nominal", Around{Self: glued, Before: []Near{noun}}, AroundAfterNominal, false, 1},
		{"after a particle", Around{Self: glued, Before: []Near{{Stem: "패소", Last: "주격:가"}}}, AroundAfterNominal, false, -1},
		{"nothing before", Around{Self: glued}, AroundAfterNominal, false, -1},
		{"only the nearest before", Around{Self: glued, Before: []Near{{Stem: "을"}, noun}}, AroundAfterNominal, false, -1},
		{"before a nominal", Around{Self: Near{Stem: "감정인"}, After: []Near{{Stem: "사람", Modifiable: true}}}, AroundBeforeModifiable, false, 1},
		{"before an adverb", Around{Self: Near{Stem: "감정인"}, After: []Near{head}}, AroundBeforeModifiable, false, -1},
		{"at the end", Around{Self: Near{Stem: "감정인"}}, AroundBeforeModifiable, false, -1},
		{"one of two", Around{Self: glued}, AroundGlued | AroundAfterNominal, false, 1},
		{"not, held", Around{Self: glued}, AroundGlued, true, -1},
		{"not, missing", Around{Self: Near{Stem: "감정인"}, After: []Near{head}}, AroundBeforeModifiable, true, 1},
	}
	for _, c := range cases {
		ctx := toulmin.NewContext()
		ctx.Set(KeyAround, c.a)
		specs := toulmin.Specs{&Bits{Of: "around", Mask: c.mask}}
		if c.not {
			specs = append(specs, &Not{})
		}
		if got, fact := AroundIs(ctx, specs); got != c.want || fact != nil {
			t.Errorf("%s: %v %v, want %v", c.name, got, fact, c.want)
		}
	}
	if got, _ := AroundIs(toulmin.NewContext(), toulmin.Specs{&Bits{Of: "around", Mask: AroundGlued}}); got != -1 {
		t.Error("no candidate in the context", got)
	}
}
