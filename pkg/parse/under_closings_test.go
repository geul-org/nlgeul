//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestUnderClosings — 는 put under a closing's last syllable as ㄴ is one more piece (는) after the closing: after the linking 어서 (먹어선 = 먹+어서+(는)) and 다가 (가다간), after a closing that asks (먹을진 = 먹+을지+(는)), also under a final consonant's closing (갈진 = 가+(ㄹ)지+(는)); nothing at a unit's first letters, under a syllable with no ㄴ, before a Hangul letter (먹어선가), under a syllable that ends none of the closings (먹은), under the bare 어 (먹언), for a pair that is a shortened closing itself (먹진: 진 is 지는), after a closing fused wholly into the stem's syllable (나간 has no 나+(아)+(는)), nor after a closing 는 is not written onto (먹든간: 든가)
package parse

import (
	"slices"
	"testing"
)

func TestUnderClosings(t *testing.T) {
	for _, c := range []struct {
		in, closing string
		j           int
	}{
		{"먹어선", "맺음:어서", 3},
		{"가다간", "맺음:다가", 3},
		{"먹을진", "맺음:을지", 3},
		{"갈진", "맺음:ㄹ지", 2},
	} {
		got := scanOf(c.in).underClosings(c.j)
		ok := slices.ContainsFunc(got, func(t tail) bool {
			ps := t.pieces
			return len(ps) >= 2 && ps[len(ps)-2].names == c.closing && ps[len(ps)-1].text == underText && ps[len(ps)-1].names == underName && ps[len(ps)-1].verdict > 0
		})
		if !ok {
			t.Fatalf("%s: no %s with (는) after it: %v", c.in, c.closing, got)
		}
	}
	for _, c := range []struct {
		in string
		j  int
	}{
		{"선", 1},
		{"먹어서", 3},
		{"먹어선가", 3},
		{"먹은", 2},
		{"먹언", 2},
		{"먹진", 2},
		{"나간", 2},
		{"먹든간", 3},
	} {
		if got := scanOf(c.in).underClosings(c.j); len(got) != 0 {
			t.Fatalf("%s: %v", c.in, got)
		}
	}
	sc := scanOf("먹어선")
	if sc.underClosings(3); string(sc.s) != "먹어선" || sc.letters(0, 3) != "먹어선" {
		t.Fatal("the unit's own letters are changed", string(sc.s))
	}
}
