//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamPlate — end to end with no word table: a motor vehicle's registration number is one word of pieces, the kind and the serial number each a number and the letter named 차등록 (206[수:십진]+너[차등록:너]+9499[수:십진]), the unit 호 and particles after it follow in the same word, the name of a region before it is the word's stem, a sign of a pair before it and after it is a contest of its own; letters that are a court case number too are read as that first and as a registration number after it (97도2836), and numbers that only look so stay numbers (60조5000억)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamPlate(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct{ sentence, first string }{
		{"차량 206너9499 한 대", "206[수:십진]+너[차등록:너]+9499[수:십진]"},
		{"번호 306너8412가 찍힌", "306[수:십진]+너[차등록:너]+8412[수:십진]+가[주격:가]"},
		{"소유의 서울34가5678호 차량", "서울[열쇠:체언]+34[수:십진]+가[차등록:가]+5678[수:십진]+호[단위:호]"},
		{"차량(306너8412)을", "306[수:십진]+너[차등록:너]+8412[수:십진]"},
		{"대법원 97도2836 판결", "97[수:십진]+도[사건:도]+2836[수:십진]"},
		{"예산 60조5000억", "60[수:십진]+조[큰 수:조] 5000[수:십진]+억[큰 수:억]"},
	}
	for _, c := range cases {
		first := "\n" + strings.Join(bestLines(r, c.sentence), "\n") + "\n"
		if !strings.Contains(first, "\n"+c.first+"\n") {
			t.Errorf("%s: %q is not first in\n%s", c.sentence, c.first, first)
		}
	}
	if box := strings.Join(r.BoxLines("대법원 97도2836 판결", false, 0), "\n"); !strings.Contains(box, "97[수:십진]+도[차등록:도]+2836[수:십진]") {
		t.Error("97도2836 is not read as a registration number too", box)
	}
}
