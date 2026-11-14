//ff:func feature=token type=reader control=iteration dimension=2
//ff:what TestReaderLinesRadix — 0x and hexadecimal digits is the one candidate 수:16진법 (with a word table too), and what a number takes stands after it (면, 를); a number of 0 and 1 only is a contest of 수 and 수:2진법 (before a unit too); another digit, a sign, a number joined by a sign, 0x12G4, 0X1F are not
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesRadix(t *testing.T) {
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
		{"0x1a2F", []string{"0x1a2F[수:16진법]"}, []string{"경합", "[열:영숫자]", "[수]"}},
		{"0x1010", []string{"0x1010[수:16진법]"}, []string{"경합", "2진법]\n", "[수]"}},
		{"0x1234면", []string{"  0x1234[수:16진법]+면[맺음:면]\t성립 1.00"}, []string{"[열:영숫자]"}},
		{"0x1234를", []string{"  0x1234[수:16진법]+를[목적격:를]\t성립 1.00"}, nil},
		{"0x12G4", []string{"0x12G4[열:영숫자]"}, []string{"16진법"}},
		{"0x", []string{"0x[열:영숫자]"}, []string{"16진법"}},
		{"0X1F", []string{"0X1F[열:영숫자]"}, []string{"16진법"}},
		{"255", []string{"255[수]"}, []string{"경합", "2진법"}},
		{"102", []string{"102[수]"}, []string{"경합", "2진법"}},
		{"1010", []string{"1010\t경합(수 / 수:2진법)", "  1010[수]\t성립 1.00", "  1010[수:2진법]\t성립 1.00"}, nil},
		{"0", []string{"0\t경합(수 / 수:2진법)"}, nil},
		{"1", []string{"1\t경합(수 / 수:2진법)"}, nil},
		{"10개", []string{"  10[수]+개[단위:개]\t성립 1.00", "  10[수:2진법]+개[단위:개]\t성립 1.00"}, nil},
		{"100%", []string{"  100[수]+%[단위:퍼센트]\t", "  100[수:2진법]+%[단위:퍼센트]\t"}, nil},
		{"1,000", []string{"1,000[수]"}, []string{"경합", "2진법"}},
		{"1.01", []string{"1.01\t경합(소수점 / 버전)", "  1.01[소수점]\t", "  1.01[버전]\t"}, []string{"2진법", "|"}},
		{"10:10", []string{"10:10\t경합(시각 / 비율)", "  10:10[시각]\t", "  10:10[비율]\t"}, []string{"2진법", "|"}},
		{"-11", []string{"-11\t경합(부호 / 수)", "  -11[부호]\t", "  -11[수]\t"}, []string{"2진법", "|"}},
	}
	for _, r := range []*Reader{bare, words} {
		for _, c := range cases {
			checkLines(t, c.sentence, strings.Join(r.Lines(c.sentence, false), "\n")+"\n", c.has, c.lacks)
		}
	}
}
