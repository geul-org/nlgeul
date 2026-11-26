//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestSpokenOf — a line of the table of markers written as they are spoken is read into the letters it is written with (whole syllables, under a final consonant, or after a fused 어/아), the name of the standard marker and the copula condition; a form that is not Hangul, a form with no syllable, a standard marker that is not an index name and a condition the table does not know are errors that name the column
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSpokenOf(t *testing.T) {
	for _, r := range []struct {
		row  []string
		want spoken
	}{
		{[]string{"구", "맺음:고", "-"}, spoken{"구", 0, "맺음:고", false}},
		{[]string{"ㄴ다구", "맺음:ㄴ다고", "-"}, spoken{"다구", rule.JongN, "맺음:ㄴ다고", false}},
		{[]string{"어/아두", "맺음:어도(합침)", "-"}, spoken{"두", rule.LeadVowel, "맺음:어도(합침)", false}},
		{[]string{"요", "맺음:오", "서술격만"}, spoken{"요", 0, "맺음:오", true}},
	} {
		if got, err := spokenOf(rowFor(r.row...)); err != nil || got != r.want {
			t.Errorf("%v: %+v, %v", r.row, got, err)
		}
	}
	for _, r := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "맺음:고", "-"}, []string{"1번째 칸"}},
		{[]string{"ㄴ", "맺음:ㄴ", "-"}, []string{"1번째 칸", "적힌 꼴"}},
		{[]string{"구", "고", "-"}, []string{"2번째 칸", "색인의 이름"}},
		{[]string{"구", "맺음:고", "요 안 됨"}, []string{"3번째 칸", "정해진 이름이 아니다"}},
	} {
		_, err := spokenOf(rowFor(r.row...))
		errHas(t, err, r.parts...)
	}
}
