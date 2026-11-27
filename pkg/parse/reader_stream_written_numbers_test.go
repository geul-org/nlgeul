//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestReaderStreamWrittenNumbers — numbers that are no run of digits, end to end with a word table: each case gives candidates of a made-up sentence in the order they are read, and what must not be read; a number its word tells (a unit, digits or a bracket written onto it) comes first, and where nothing in the word tells a number, the number or the unit written apart is a candidate behind the word of the word table
package parse

import "testing"

func TestReaderStreamWrittenNumbers(t *testing.T) {
	r := streamReader(t, "words-written-numbers.tsv")
	cases := []struct {
		sentence string
		order    []string
		lacks    []string
	}{
		{"칠팔년을", []string{"칠[수:한글십진]+팔[수:한글십진]+년[단위:년]+을[목적격:을]", "칠팔년을"}, nil},
		{"이삼십명", []string{"이[수:한글십진]+삼십[수:한글십진]+명[단위:명]", "이삼십명"}, nil},
		{"육칠 명", []string{"육칠[열쇠:체언]", "육[수:한글십진]+칠[수:한글십진]", "명[열쇠:체언]", "명[단위:명]"}, nil},
		{"사십 년", []string{"사십[열쇠:체언]", "사십[수:한글십진]", "년[열쇠:체언]", "년[단위:년]"}, nil},
		{"오 년을", []string{"오[열쇠:체언]", "오[수:한글십진]", "년[열쇠:체언]+을[목적격:을]"}, []string{"년[단위:년]+을"}},
		{"사십을", []string{"사십[열쇠:체언]+을[목적격:을]"}, []string{"사십[수:한글십진]+을"}},
		{"자", []string{"자[열쇠:홀로]", "자[단위:자]"}, nil},
		{"자", []string{"자[열쇠:체언]", "자[단위:자]"}, nil},
		{"천500억", []string{"천[큰 수:천]", "천[열쇠:체언]", "천[수:한글십진]", "500[수:십진]+억[큰 수:억]"}, nil},
		{"천6백여", []string{"천[큰 수:천]", "천[열쇠:체언]", "6[수:십진]+백[큰 수:백]+여[어림:여]"}, nil},
		{"천", []string{"천[열쇠:체언]", "천[수:한글십진]"}, []string{"큰 수:천"}},
		{"만19세", []string{"만[열쇠:관형사]", "19[수:십진]+세[단위:세]"}, []string{"큰 수:만"}},
		{"二三일", []string{"二三[수:한자십진]+일[단위:일]", "二三일"}, nil},
		{"十원은", []string{"十[수:한자십진]+원[단위:원]+은[보조사:은]", "十원은"}, nil},
		{"三인", []string{"三[수:한자십진]+인[단위:인]", "三인"}, nil},
		{"十", []string{"十[수:한자십진]", "十"}, nil},
		{"第一회", []string{"第一회"}, []string{"수:한자십진"}},
		{"Ⅱ", []string{"Ⅱ[수:로마십진]"}, []string{"열:영숫자"}},
		{"Ⅰ의", []string{"Ⅰ[수:로마십진]+의[관형격:의]", "Ⅰ의"}, nil},
		{"ⅰ호는", []string{"ⅰ[수:로마십진]+호[단위:호]+는[보조사:는]", "ⅰ호는"}, nil},
		{"제Ⅱ장", []string{"제Ⅱ[차례]+장[단위:장]", "제Ⅱ[수:로마십진]+장[단위:장]", "제Ⅱ장"}, nil},
		{"XIV세", []string{"XIV[수:로마십진]+세[단위:세]", "XIV[열:영숫자] 세[열쇠:체언]"}, nil},
		{"제II장", []string{"제II[차례]+장[단위:장]", "제II[수:로마십진]+장[단위:장]", "제[열쇠:체언] II[열쇠:체언] 장"}, nil},
		{"크로뮴(III)", []string{"크로뮴[열쇠:체언]", "([열기:1]", "III[수:로마십진]", "III[열:영숫자]", ")[닫기:1]"}, nil},
		{"크로뮴(II)", []string{"II[수:로마십진]", "II[열쇠:체언]"}, nil},
		{"크로뮴 (II)", []string{"II[열쇠:체언]", "II[수:로마십진]"}, nil},
		{"II", []string{"II[열쇠:체언]", "II[수:로마십진]"}, nil},
		{"XIV", []string{"XIV[열:영숫자]", "XIV[수:로마십진]"}, nil},
		{"V", []string{"V[열:영숫자]", "V[수:로마십진]"}, nil},
		{"V세", []string{"V[열:영숫자]", "V[수:로마십진]"}, []string{"V[수:로마십진]+세"}},
		{"XX동", []string{"XX동", "XX[열:영숫자] 동", "XX[수:로마십진] 동"}, []string{"XX[수:로마십진]+"}},
		{"CD", []string{"CD[열:영숫자]"}, []string{"수:로마십진"}},
	}
	for _, c := range cases {
		lines := streamLines(r, c.sentence)
		if !linesInOrder(lines, c.order) {
			t.Errorf("%q: not in the order %v\n%s", c.sentence, c.order, streamText(r, c.sentence))
		}
		checkLines(t, c.sentence, streamText(r, c.sentence), nil, c.lacks)
	}
}
