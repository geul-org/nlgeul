//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamCaseNumber — end to end with no word table: a court case number is one word with no name (2004다25123, 97도2836 before a comma), a unit and particles after it are read after it (2008회합9+호+로), and numbers that only look so stay numbers (17만2000, 2020년3월, 11의2)
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
		{"대법원 2004다25123 판결", "2004다25123"},
		{"대법원 97도2836, 판결", "97도2836"},
		{"서울 2008회합9호로 신청", "2008회합9+호[단위:호]+로[부사격:로]"},
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
