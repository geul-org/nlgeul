//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestReaderStreamUnits — a unit is named by the standard Hangul name of the unit table (킬로미터, 기가바이트, 섭씨), whatever the letters; a letter on two lines (m) is a contest of two candidates; a number and a Latin unit read as one run is named 열:영숫자; 도 after a number is a unit and a particle, both; the lines of a unit in the table are read right after a number, and no unit is read after a space
package parse

import "testing"

func TestReaderStreamUnits(t *testing.T) {
	bare, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	words, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		sentence string
		has      []string
		lacks    []string
	}{
		{"7m", []string{"\n7[수:십진]+m[단위:미터]\t성립 1.00", "\n7[수:십진]+m[단위:분]\t성립 1.00", "\n7m[열:영숫자]\t성립 0.60"}, nil},
		{"64GB", []string{"\n64[수:십진]+GB[단위:기가바이트]\t성립 1.00", "\n64GB[열:영숫자]\t성립 0.60"}, []string{"단위:GB"}},
		{"45도", []string{"\n45[수:십진]+도[보조사:도]\t성립 1.00", "\n45[수:십진]+도[단위:도]\t성립 1.00"}, nil},
		{"9㎢", []string{"\n9[수:십진]+㎢[단위:제곱킬로미터]\t성립 1.00"}, nil},
		{"9km²", []string{"\n9[수:십진]+km²[단위:제곱킬로미터]\t성립 1.00"}, nil},
		{"70℃", []string{"\n70[수:십진]+℃[단위:섭씨]\t성립 1.00"}, nil},
		{"12%", []string{"\n12[수:십진]+%[단위:퍼센트]\t성립 1.00"}, nil},
		{"20％", []string{"\n20[수:십진]+％[단위:퍼센트]\t성립 1.00"}, nil},
		{"5ｍ", []string{"\n5[수:십진]+ｍ[단위:미터]\t성립 1.00", "\n5[수:십진]+ｍ[단위:분]\t성립 1.00"}, nil},
		{"200ｇ", []string{"\n200[수:십진]+ｇ[단위:그램]\t성립 1.00"}, nil},
		{"5ｍｍ", []string{"5ｍｍ[열:영숫자]"}, []string{"단위:미터", "단위:밀리미터"}},
		{"2∼3세트", []string{"\n2∼3[범위]+세트[단위:세트]\t성립 1.00"}, nil},
		{"1～2명", []string{"\n1～2[범위]+명[단위:명]\t성립 1.00"}, nil},
		{"10〜20분", []string{"\n10〜20[범위]+분[단위:분]\t성립 1.00"}, nil},
		{"3–5일", []string{"\n3–5[범위]+일[단위:일]\t성립 1.00", "\n3–5[빼기]+일[단위:일]\t성립 1.00"}, nil},
		{"1월∼3월", []string{"\n1[수:십진]+월[단위:월]+∼[범위] 3[수:십진]+월[단위:월]\t"}, nil},
		{"2~3세트", []string{"\n2~3[범위]+세트[단위:세트]\t성립 1.00"}, nil},
		{"12 kg", []string{"kg[열:영숫자]"}, []string{"단위:킬로그램"}},
		{"85m2", []string{"\n85[수:십진]+m2[단위:제곱미터]\t성립 1.00", "\n85m2[열:영숫자]\t성립 0.60"}, nil},
		{"3,456m2", []string{"\n3,456[수:십진]+m2[단위:제곱미터]\t성립 1.00"}, []string{"쉼표"}},
	}
	for _, r := range []*Reader{bare, words} {
		for _, c := range cases {
			checkLines(t, c.sentence, scoredText(r, c.sentence), c.has, c.lacks)
		}
	}
}
