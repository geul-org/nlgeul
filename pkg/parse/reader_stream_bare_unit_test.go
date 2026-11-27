//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamBareUnit — end to end, with and without a word table: a unit sign with no number before it and 당 written onto it is one word, the unit as its stem and 당 as the suffix (㎡당, in a bracket too, with a particle after it); after a number the unit is read as before (3㎡당); a unit sign with no 당 stays a sign
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamBareUnit(t *testing.T) {
	with, err := New(Options{Words: "testdata/words-suffix.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	without, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence, first string
		bare            bool
	}{
		{"㎡당 비율이 높다", "㎡[단위:제곱미터]+당[접미:당]", true},
		{"(㎡당 비율)", "㎡[단위:제곱미터]+당[접미:당]", true},
		{"㎡당으로 높다", "㎡[단위:제곱미터]+당[접미:당]+으로[부사격:으로]", true},
		{"3㎡당 비율", "3[수:십진]+㎡[단위:제곱미터]+당[접미:당]", false},
		{"㎡ 비율", "", false},
	}
	for _, c := range cases {
		all := streamText(with, c.sentence) + streamText(without, c.sentence)
		if c.first != "" && strings.Count(all, c.first) < 2 {
			t.Errorf("%s: %q is not read with and without a word table in\n%s", c.sentence, c.first, all)
		}
		if strings.Contains(all, "\n㎡[단위:제곱미터]+당[접미:당]") != c.bare {
			t.Errorf("%s: a unit with no number read %v, want %v\n%s", c.sentence, !c.bare, c.bare, all)
		}
	}
}
