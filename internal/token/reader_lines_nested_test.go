//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestReaderLinesNested — with no word table, brackets and quotation marks nest and are cut where they stand, from the first letter on: 함수(인자(안쪽))를 gives 열기:1, 열기:2, 닫기:2 and a closing sign with 를 on it as 닫기:1+목적격 beside the whole 새 낱말; a JSON object gives the braces, the quotation marks at depth 2 and the 쌍점 and 쉼표 between; a particle written onto the closing brace reads 로 after 닫기:1
package token

import (
	"strings"
	"testing"
)

func TestReaderLinesNested(t *testing.T) {
	r, err := New(Options{})
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		in   string
		want []string
	}{
		{"함수(인자(안쪽))를 호출", []string{"([열기:1]", "인자", "([열기:2]", "안쪽", ")[닫기:2]", ")[닫기:1]", "를\t경합(목적격:를 / 새 낱말)", "  +를[목적격:를]\t성립 1.00", "  를\t성립 0.33", "호출"}},
		{"{\"이름\": \"작은 고양이\", \"수\": [1, 2]}", []string{"{[열기:1]", "\"[열기:2]", "이름", "\"[닫기:2]", ":[쌍점]", "\"[열기:2]", "작은", "고양이", "\"[닫기:2]", ",[쉼표]", "\"[열기:2]", "수[의존명사:수]", "\"[닫기:2]", ":[쌍점]", "[[열기:2]", "1\t경합(수 / 수:2진법)", "  1[수]\t성립 1.00", "  1[수:2진법]\t성립 1.00", ",[쉼표]", "2[수]", "][닫기:2]", "}[닫기:1]"}},
		{"설정은 {\"mode\": \"fast\"}로 해", []string{"설정은", "{[열기:1]", "\"[열기:2]", "mode[열:영숫자]", "\"[닫기:2]", ":[쌍점]", "\"[열기:2]", "fast[열:영숫자]", "\"[닫기:2]", "}[닫기:1]", "로\t경합(부사격:로 / 새 낱말)", "  +로[부사격:로]\t성립 1.00", "해"}},
		{"「책」을 “말”이라고", []string{"「[열기:1]", "책", "」[닫기:1]", "을\t경합(목적격:을 / 새 낱말)", "“[열기:1]", "말", "”[닫기:1]", "이라고\t경합(부사격:이라고 /", "  +이라고[부사격:이라고]"}},
		{"(가", []string{"([열기:1]", "가"}},
	} {
		lines := r.Lines(c.in, false)
		if miss := missingLine(lines, c.want); miss != "" {
			t.Errorf("%q: no line starts with %q in order, in\n%s", c.in, miss, strings.Join(lines, "\n"))
		}
	}
}
