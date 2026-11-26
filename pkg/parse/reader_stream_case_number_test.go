//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCaseNumber — end to end with no word table: a court case number is one word of three pieces, the year and the serial number each a number and the mark named 사건 (2004[수]+다[사건:다]+25123[수], 97도2836 before a comma, a mark of two syllables in 2007가단25609, a mark of the prosecution in 67형53722), a unit and particles after it follow in the same word (2008회합9 with 호 and 로), a sign of a pair before it and after it is a contest of its own, and numbers that only look so stay numbers (17만2000, 2020년3월, 11의2)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamCaseNumber(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, first string }{
		{"대법원 2004다25123 판결", "2004[수]+다[사건:다]+25123[수]"},
		{"대법원 97도2836, 판결", "97[수]+도[사건:도]+2836[수]"},
		{"법원 2007가단25609 사건", "2007[수]+가단[사건:가단]+25609[수]"},
		{"검찰청 67형53722호 기록", "67[수]+형[사건:형]+53722[수]+호[단위:호]"},
		{"판결(2004다25123)에서", "2004[수]+다[사건:다]+25123[수]"},
		{"서울 2008회합9호로 신청", "2008[수]+회합[사건:회합]+9[수]+호[단위:호]+로[부사격:로]"},
		{"인구 17만2000", "17[수]+만[큰 수:만] 2000[수]"},
		{"2020년3월 시작", "2020[수]+년[단위:년] 3[수]+월[단위:월]"},
		{"별표 11의2", "11[수]+의[관형격:의] 2[수]"},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(r, c.sentence), "\n") + "\n"
		if !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
	}
}
