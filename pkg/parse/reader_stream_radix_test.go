//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what TestReaderStreamRadix — 0x and hexadecimal digits is the one candidate 수:16진법 (with a word table too), and what a number takes stands after it (the copula 이 with 면, 를), a closing not right on its digit (0x1234면); a number of 0 and 1 only is a contest of 수 and 수:2진법 (before a unit too); another digit, a sign, a number joined by a sign, 0x12G4, 0X1F are not
package parse

import "testing"

func TestReaderStreamRadix(t *testing.T) {
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
		one      bool
		has      []string
		lacks    []string
	}{
		{"0x1a2F", true, []string{"0x1a2F[수:16진법]"}, []string{"[열:영숫자]", "[수:십진]"}},
		{"0x1010", true, []string{"0x1010[수:16진법]"}, []string{"2진법]\t", "[수:십진]"}},
		{"0x1234이면", false, []string{"\n0x1234[수:16진법]+이[서술격:이]+면[맺음:면]\t성립 1.00"}, []string{"[열:영숫자]"}},
		{"0x1234면", false, nil, []string{"\n0x1234[수:16진법]+면[맺음:면]"}},
		{"0x1234를", false, []string{"\n0x1234[수:16진법]+를[목적격:를]\t성립 1.00"}, nil},
		{"0x12G4", false, []string{"0x12G4[열:영숫자]"}, []string{"16진법"}},
		{"0x", false, []string{"0x[열:영숫자]"}, []string{"16진법"}},
		{"0X1F", false, []string{"0X1F[열:영숫자]"}, []string{"16진법"}},
		{"255", true, []string{"255[수:십진]"}, []string{"2진법"}},
		{"102", true, []string{"102[수:십진]"}, []string{"2진법"}},
		{"1010", false, []string{"\n1010[수:십진]\t성립 1.00", "\n1010[수:2진법]\t성립 1.00"}, nil},
		{"0", false, []string{"\n0[수:십진]\t", "\n0[수:2진법]\t"}, nil},
		{"1", false, []string{"\n1[수:십진]\t", "\n1[수:2진법]\t"}, nil},
		{"10개", false, []string{"\n10[수:십진]+개[단위:개]\t성립 1.00", "\n10[수:2진법]+개[단위:개]\t성립 1.00"}, nil},
		{"100%", false, []string{"\n100[수:십진]+%[단위:퍼센트]\t", "\n100[수:2진법]+%[단위:퍼센트]\t"}, nil},
		{"1,000", true, []string{"1,000[수:십진]"}, []string{"2진법"}},
		{"1.01", false, []string{"\n1.01[소수점]\t", "\n1.01[버전]\t"}, []string{"2진법", "|"}},
		{"10:10", false, []string{"\n10:10[시각]\t", "\n10:10[비율]\t"}, []string{"2진법", "|"}},
		{"-11", false, []string{"\n-11[부호]\t", "\n-11[수:십진]\t"}, []string{"2진법", "|"}},
	}
	for _, r := range []*Reader{bare, words} {
		for _, c := range cases {
			checkLines(t, c.sentence, scoredText(r, c.sentence), c.has, c.lacks)
			if c.one && len(streamLines(r, c.sentence)) != 1 {
				t.Errorf("%s: more than the one candidate %v", c.sentence, streamLines(r, c.sentence))
			}
		}
	}
}
