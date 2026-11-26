//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamHangulUnits — a number directly followed by a Hangul unit (미터, 톤, 인치, 퍼센트, 매) is read as number + unit with the particle after it, and the noun reading stays a candidate; a spaced number or a bare noun has no unit reading
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamHangulUnits(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct{ in, has, not string }{
		{"12미터를 갔다.", "12[수]+미터[단위:미터]+를[목적격:를]", ""},
		{"3.5톤의 짐이다.", "3.5[소수점]+톤[단위:톤]+의[관형격:의]", ""},
		{"8퍼센트로 올랐다.", "8[수]+퍼센트[단위:퍼센트]+로[부사격:로]", ""},
		{"7인치 화면이다.", "7[수]+인치[단위:인치]", ""},
		{"4매를 냈다.", "4[수]+매[단위:매]+를[목적격:를]", ""},
		{"12 미터를 갔다.", "12[수]", "+미터[단위:미터]"},
		{"미터를 쟀다.", "미터", "[단위:미터]"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || (c.not != "" && strings.Contains(got, c.not)) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
