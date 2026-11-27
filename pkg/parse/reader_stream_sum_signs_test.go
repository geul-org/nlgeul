//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamSumSigns — the signs of sums and of units: a plus written onto a number's unit sign is 더하기 and 연결 (90%+), and cuts the sentence there (90%+10%); the full-width plus ＋ is read as the plus is; a tilde between letters right before a digit is a word of its own (벌금~300, 3.∼5) while the range sign after a number and its unit, the range between numbers, a path's head and a name stay as they were; an equals sign with a Hangul syllable on a side is also read as a word of its own with no name (값=3) and a=1 stays one name; a backslash at a unit's edge is \[기호], not one next to another sign (C:\, \\) nor before a Latin letter (\n); the dollar sign after a number is its unit (300$, and $40 is money as before); km/h after a number is a unit; a ratio joins any number of numbers (5:3:2:1) and a time stays first (10:30:15)
package parse

import (
	"strings"
	"testing"
)

func TestReaderStreamSumSigns(t *testing.T) {
	r, err := New(Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in, has, not string
	}{
		{"정확도 90%+ 속도를 본다.", "\n90%\n+[더하기]\n+[연결]\n", ""},
		{"90%+10%를 더한다.", "\n90%\n+[더하기]\n+[연결]\n10[수:십진]+%[단위:퍼센트]+를[목적격:를]\n", ""},
		{"곰＋솔을 본다.", "\n＋[더하기]\n＋[연결]\n", "곰＋솔"},
		{"가 ＋ 나", "\n＋[더하기]\n＋[연결]\n", ""},
		{"C++ 을 쓴다.", "\n++\n", "[더하기]"},
		{"a+=1 을 쓴다.", "a+=1", "[더하기]"},
		{"+3 이다.", "\n+3[부호]\n", "[더하기]"},
		{"벌금~300만 원을 낸다.", "\n벌금 ~[물결표] 300[수:십진]+만[큰 수:만]\n", ""},
		{"기간은 3.∼5. 이다.", " ∼[물결표] 5[수:십진]\n", ""},
		{"1월~3월에 한다.", "\n1[수:십진]+월[단위:월]+~[범위] 3[수:십진]+월[단위:월]+에[부사격:에]\n", "[물결표]"},
		{"3~4명이 온다.", "\n3~4[범위]+명[단위:명]+이[주격:이]\n", "[물결표]"},
		{"~/a/b를 연다.", "~/a/b[경로]+를[목적격:를]", "[물결표]"},
		{"a~b를 쓴다.", "a~b+를[목적격:를]", "[물결표]"},
		{"~3 이다.", "\n~\n", "[물결표]"},
		{"값=3 으로 둔다.", "\n값 =[기호:같음] 3[수:십진]\n", "[기호]"},
		{"값=3 으로 둔다.", "\n값=3\n", ""},
		{"a=1 로 둔다.", "\na=1\n", "a = 1"},
		{"x = 3 이다.", "\n=[기호:같음]\n", "[기호]"},
		{"가==나", "가==나", "가 == 나"},
		{"무조건 \\붙인다.", "\n\\[기호]\n", ""},
		{"\\ 하나를 쓴다.", "\n\\[기호]\n", ""},
		{"C:\\ 를 연다.", "\n\\\n", "[기호]"},
		{"\\\\ 를 쓴다.", "\n\\\\\n", "[기호]"},
		{"\\n 을 쓴다.", "\n\\\n", "[기호]"},
		{"미화 300$를 냈다.", "\n300[수:십진]+$[단위:달러]+를[목적격:를]\n", ""},
		{"값은 $40 이다.", "\n$40[돈]\n", "[단위:달러]"},
		{"시속 151km/h의 공이다.", "\n151[수:십진]+km/h[단위:km/h]+의[관형격:의]\n", ""},
		{"거리는 151km 이다.", "\n151[수:십진]+km[단위:킬로미터]\n", ""},
		{"비율은 5:3:2:1 이다.", "\n5:3:2:1[비율]\n", ""},
		{"10:30:15 에 온다.", "\n10:30:15[시각]\n10:30:15[비율]\n", ""},
		{"3:1 로 이겼다.", "\n3:1[비율]\n", "[시각]"},
	} {
		got := streamText(r, c.in)
		if !strings.Contains(got, c.has) || (c.not != "" && strings.Contains(got, c.not)) {
			t.Errorf("%q: want %q and no %q\n%s", c.in, c.has, c.not, got)
		}
	}
}
