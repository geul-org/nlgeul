//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestUnderYo — 죠 at the end of a longer final closing is its 지 with 요 fused in: 먹어야죠 = 먹+어야지+(요), 가야죠 = 가+(아)야지+(요), 가라죠 = 가+라지+(요); the plain 지 is not read here (먹죠), nor a closing that is none of the few (먹는죠 has no 는지+(요)), nor anything at a unit's first letters or under another letter
package parse

import (
	"slices"
	"testing"
)

func TestUnderYo(t *testing.T) {
	for _, c := range []struct {
		in, closing string
		j           int
	}{
		{"먹어야죠", "맺음:어야지", 4},
		{"가야죠", "맺음:어야지(합침)", 3},
		{"가라죠", "맺음:라지", 3},
	} {
		got := scanOf(c.in).underYo(c.j)
		ok := slices.ContainsFunc(got, func(t tail) bool {
			ps := t.pieces
			return len(ps) >= 2 && ps[len(ps)-2].names == c.closing && ps[len(ps)-1].text == yoText && ps[len(ps)-1].names == yoName
		})
		if !ok {
			t.Fatalf("%s: no %s with (요) after it: %v", c.in, c.closing, got)
		}
	}
	for _, c := range []struct {
		in string
		j  int
	}{
		{"죠", 1},
		{"먹죠", 2},
		{"먹는죠", 3},
		{"먹어야지", 4},
	} {
		if got := scanOf(c.in).underYo(c.j); len(got) != 0 {
			t.Fatalf("%s: %v", c.in, got)
		}
	}
	sc := scanOf("가야죠")
	if sc.underYo(3); string(sc.s) != "가야죠" {
		t.Fatal("the unit's own letters are changed", string(sc.s))
	}
}
