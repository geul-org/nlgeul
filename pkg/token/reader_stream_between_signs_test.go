//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderStreamBetweenSigns — the signs between words of a unit: the ellipsis … closes the word before it (올랐다+…[줄임표], 창고+……[줄임표]) and alone is …[줄임표]; a slash and a tilde (∼ too) between Hangul syllables are words of their own in a way of the stretch (압박골절 /[빗금] 가로돌기골절+로, 여름 ~[물결표] 가을; the stream sends the whole letters as a new word first), and so are the dashes ― — (―[줄표]) and the middle dot variants ‧ ․ • (가운뎃점, after the whole letters as well); a range after a number and its unit, a bound pair of known words, names, paths, addresses, a slash at a word's edge, the middle dot and box lines are read as before
package token

import (
	"strings"
	"testing"
)

func TestReaderStreamBetweenSigns(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, unit, first, has, not string
	}{
		{"값이 올랐다…그래서 샀다.", "올랐다…그래서", "", "\n올랐다+…[줄임표] 그래서[접속 부사:그래서]\n", ""},
		{"창고……건평이 넓다.", "창고……건평이", "", "\n창고+……[줄임표] 건평+이[주격:이]\n", ""},
		{"그는 떠났다… 아무도 몰랐다.", "", "", "\n…[줄임표]\n", ""},
		{"압박골절/가로돌기골절로 본다.", "압박골절/가로돌기골절로", "압박골절/가로돌기골절+로[부사격:로]", "\n압박골절 /[빗금] 가로돌기골절+로[부사격:로]\n", ""},
		{"여름~가을 사이에 핀다.", "여름~가을", "여름~가을", "\n여름 ~[물결표] 가을\n", ""},
		{"서울∼부산 간을 달린다.", "서울∼부산", "서울∼부산", "\n서울 ∼[물결표] 부산\n", ""},
		{"이것이 답이다―누구나 안다.", "답이다―누구나", "", "\n답+이[서술격:이]+다[맺음:다] ―[줄표] 누구+나[보조사:나]\n", ""},
		{"토지․건축물 값을 냈다.", "토지․건축물", "", "\n토지 ․[가운뎃점] 건축물\n", ""},
		{"벽보‧전단을 붙였다.", "벽보‧전단을", "벽보‧전단+을[목적격:을]", "\n벽보 ‧[가운뎃점] 전단+을[목적격:을]\n", ""},
		{"아라비안•나이트를 읽었다.", "아라비안•나이트를", "아라비안•나이트+를[목적격:를]", "\n아라비안 •[가운뎃점] 나이트+를[목적격:를]\n", ""},
		{"1월~3월에 한다.", "1월~3월에", "1[수]+월[단위:월]+~[범위] 3[수]+월[단위:월]+에[부사격:에]", "", "[물결표]"},
		{"입력/출력을 본다.", "입력/출력을", "입력/출력[대비:/]+을[목적격:을]", "", ""},
		{"ab/cd를 연다.", "", "", "ab/cd[경로]+를[목적격:를]", "[빗금]"},
		{"~/a/b를 연다.", "", "", "~/a/b[경로]+를[목적격:를]", "[빗금]"},
		{"http://a.com/b/c를 연다.", "", "", "http://a.com/b/c[주소]+를[목적격:를]", "[빗금]"},
		{"a~b를 쓴다.", "", "", "a~b+를[목적격:를]", "[물결표]"},
		{"수/ 없다.", "", "", "\n/\n", "[빗금]"},
		{"토지·건물을 샀다.", "", "", "\n토지 ·[가운뎃점] 건물+을[목적격:을]\n", ""},
		{"├──┼", "", "", "├──┼", "[줄표]"},
	} {
		got := streamText(r, c.in)
		if c.first != "" && stretchFirst(r, c.in, c.unit) != c.first {
			t.Errorf("%q: the first line of %q is not %q\n%s", c.in, c.unit, c.first, got)
		}
		if !strings.Contains(got, c.has) || (c.not != "" && strings.Contains(got, c.not)) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
