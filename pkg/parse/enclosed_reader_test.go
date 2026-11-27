//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestEnclosedReader — a circled or parenthesized number is read as a number with a particle, a unit or an order after it, as a digit is, a parenthesized Hangul letter is a sign named 괄호문자:한글 (one piece for each letter) and no number, and circled Hangul is left a symbol with no name
package parse

import (
	"strings"
	"testing"
)

func TestEnclosedReader(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for text, want := range map[string]string{
		"①":    "①[원문자:수]",
		"⑴":    "⑴[괄호문자:수]",
		"㉑":    "㉑[원문자:수]",
		"㊿":    "㊿[원문자:수]",
		"⑦의":   "⑦[원문자:수]+의[관형격:의]",
		"①은":   "①[원문자:수]+은[보조사:은]",
		"②번으로": "②[원문자:수]+번[단위:번]+으로[부사격:으로]",
		"③째":   "③[원문자:수]+째[차례:째]",
		"①이다":  "①[원문자:수]+이[서술격:이]+다[맺음:다]",
		"사실⑤":  " ⑤[원문자:수]",
		"㈎":    "㈎[괄호문자:한글]",
		"㈜":    "㈜[괄호문자:한글]",
		"㈎㈏":   "㈏[괄호문자:한글]",
	} {
		if got := streamText(r, text); !strings.Contains(got, want) {
			t.Errorf("%s: want %s in\n%s", text, want, got)
		}
	}
	for _, text := range []string{"ⓐ", "㉠", "㈎", "㈎㈏"} {
		if got := streamText(r, text); strings.Contains(got, "[수") {
			t.Errorf("%s is no number: %s", text, got)
		}
	}
}
