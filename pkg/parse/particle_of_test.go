//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestParticleOf — a line of the particle table is read into the particle: form, what must stand before it, place in a chain, cases and count; a form that is not whole Hangul syllables, a name outside the sets, a place above 5 and a count that is not a number are errors that name the column
package parse

import "testing"

func TestParticleOf(t *testing.T) {
	p, err := particleOf(rowFor("이", "받침 뒤", "3", "주격|보격", "55"))
	want := particle{"이", afterBatchim, 3, caseSubj | caseComp, 55}
	if err != nil || p != want {
		t.Fatal(p, err)
	}
	p, err = particleOf(rowFor("들", "-", "0", "-", "1"))
	if err != nil || p != (particle{"들", afterAny, 0, 0, 1}) {
		t.Fatal(p, err)
	}
	cases := []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-", "3", "-", "1"}, []string{"1번째 칸(꼴)", "한글 음절"}},
		{[]string{"ㄴ다", "-", "3", "-", "1"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"ㄴ", "-", "3", "-", "1"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"", "-", "3", "-", "1"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"이", "앞", "3", "-", "1"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름"}},
		{[]string{"이", "-", "셋", "-", "1"}, []string{"3번째 칸(갈래)", "0 이상의 수가 아니다"}},
		{[]string{"이", "-", "6", "-", "1"}, []string{"3번째 칸(갈래)", "0에서 5까지"}},
		{[]string{"이", "-", "3", "처격", "1"}, []string{"4번째 칸(조건)", "관형격, 목적격"}},
		{[]string{"이", "-", "3", "-", "많이"}, []string{"5번째 칸(횟수)", "0 이상의 수가 아니다"}},
	}
	for _, c := range cases {
		_, err := particleOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
