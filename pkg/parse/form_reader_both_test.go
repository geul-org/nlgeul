//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFormReaderBoth — a sign of the tables of special letters that a hand-written table also names has both readings, the hand-written table's first: the sign table's name and then the shape (，, ＋, ～), a sign of a pair and then the shape (（, ）), a unit of the unit table standing alone and then the shape (㎡, ％); a sign no hand-written table names has the shape first and then the reading of the letter it writes (？, ！, ：); a unit sign right after its number stays a piece of the number's word (30㎡)
package parse

import (
	"strings"
	"testing"
)

func TestFormReaderBoth(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string][]string{
		"，":   {"，[쉼표]", "，[전각:기호]"},
		"～":   {"～[물결표]", "～[전각:기호]"},
		"（곰）": {"（[열기:1]", "（[전각:기호]"},
		"곰）":  {"）[닫기:1]", "）[전각:기호]"},
		"㎡":   {"㎡[단위:제곱미터]", "㎡[네모:기호+로마자]"},
		"％":   {"％[단위:퍼센트]", "％[전각:기호]"},
		"？":   {"？[전각:기호]", "？[물음표]"},
		"곰！":  {"！[전각:기호]", "！[느낌표]"},
		"곰：":  {"：[전각:기호]", "：[쌍점]"},
	} {
		got := streamText(r, text)
		first, second := strings.Index(got, want[0]), strings.Index(got, want[1])
		if first < 0 || second < 0 || first > second {
			t.Errorf("%s: want %v in that order in\n%s", text, want, got)
		}
	}
	if got := streamText(r, "30㎡"); !strings.Contains(got, "30[수:십진]+㎡[단위:제곱미터]") || strings.Contains(got, "[네모:") {
		t.Errorf("30㎡: %s", got)
	}
}
