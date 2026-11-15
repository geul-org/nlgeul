//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamAfterDot — no particle or closing is read right after a middle dot (·, ㆍ, ‧): 동ㆍ리ㆍ읍ㆍ면 has no ㆍ+면[맺음:면] and has a way of four words between dots (the stream sends the whole letters as a new word first), 갑·을 has no ·+을[목적격:을], 벽보‧를 no ‧+를; the words between dots still take their particles (도+에, 건물+을, 군+에); 시 before 군 is named as the dependent noun of the closed set
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamAfterDot(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, unit, first, has, not string
	}{
		{"동ㆍ리ㆍ읍ㆍ면 단위로 둔다.", "동ㆍ리ㆍ읍ㆍ면", "동ㆍ리ㆍ읍ㆍ면", "\n동 ㆍ[가운뎃점] 리 ㆍ[가운뎃점] 읍 ㆍ[가운뎃점] 면\n", "ㆍ+면"},
		{"갑·을 사이다.", "갑·을", "", "\n갑·을\n", "·+을"},
		{"특별시ㆍ광역시ㆍ도에 둔다.", "", "", "ㆍ[가운뎃점] 도+에[부사격:에]", "ㆍ+도"},
		{"벽보‧를 붙였다.", "", "", "\n벽보‧를\n", "‧+를"},
		{"토지·건물을 샀다.", "", "", "\n토지 ·[가운뎃점] 건물+을[목적격:을]\n", "·+"},
		{"시·군에 낸다.", "", "", "\n시[의존명사:시] ·[가운뎃점] 군+에[부사격:에]\n", "·+"},
	} {
		got := streamText(r, c.in)
		if c.first != "" && stretchFirst(r, c.in, c.unit) != c.first {
			t.Errorf("%q: the first line of %q is not %q\n%s", c.in, c.unit, c.first, got)
		}
		if !strings.Contains(got, c.has) || strings.Contains(got, c.not) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
