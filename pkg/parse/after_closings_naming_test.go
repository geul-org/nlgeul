//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestAfterClosingsNaming — end to end with the word table of testdata for the naming closings 음 and ㅁ (rows of the words in the sentences): a particle chain right after 음 or ㅁ is read in the same word and comes first (있+음+을, 끝나+(았)+음+이, 사람+이+(ㅁ)+을, 규정하+(ㅁ)+을, 참석하+(ㅁ)+으로써, 되+(ㅁ)+을, 오+(았)+음+에+도, 만들+(ㅁ)+을, 아니+(ㅁ)+을, 받+았+음+은); a noun of the table whose letters are a stem with 음 or ㅁ too (그림, 웃음, 걸음, 삶, 꿈, 봄, 잠, 다툼) stays first with the naming closing's reading among the candidates, and 마음, 처음, 먹기를, 함께, 있음 by itself and 음 with no stem before it stay as they are; where the naming closing's reading comes first (감을) the noun is still a candidate; no dropped copula stem stands before ㅁ (학굠을, 함께, 함에)
package parse

import (
	"slices"
	"strings"
	"testing"
)

func TestAfterClosingsNaming(t *testing.T) {
	r := streamReader(t, "words-naming.tsv")
	cases := []struct {
		text string
		best []string
		also string
	}{
		{"그 사실이 있음을 안다", []string{"있[보조 용언:있]+음[이름:음]+을[목적격:을]", "있[열쇠:동사]+음[이름:음]+을[목적격:을]", "있[열쇠:형용사]+음[이름:음]+을[목적격:을]"}, ""},
		{"일이 끝났음이 분명하다", []string{"끝나[열쇠:동사]+(았)[선어말:ㅆ]+음[이름:음]+이[주격:이]", "끝나[열쇠:동사]+(았)[선어말:ㅆ]+음[이름:음]+이[보격:이]"}, ""},
		{"그가 사람임을 밝혔다", []string{"사람[열쇠:체언]+이[서술격:이]+(ㅁ)[이름:ㅁ]+을[목적격:을]"}, ""},
		{"질서를 규정함을 목적으로 한다", []string{"규정하[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"}, ""},
		{"회의에 참석함으로써 마쳤다", []string{"참석하[열쇠:동사]+(ㅁ)[이름:ㅁ]+으로써[부사격:으로써]"}, ""},
		{"계약이 무효가 됨을 알았다", []string{"되[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"}, ""},
		{"비가 왔음에도 갔다", []string{"오[열쇠:동사]+(았)[선어말:ㅆ]+음[이름:음]+에[부사격:에]+도[보조사:도]"}, ""},
		{"집을 만듦을 보았다", []string{"만들[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"}, ""},
		{"그것이 아님을 알았다", []string{"아니[열쇠:형용사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"}, ""},
		{"돈을 받았음은 다툼이 없다", []string{"받[열쇠:동사]+았[선어말:았]+음[이름:음]+은[보조사:은]"}, ""},
		{"돈을 받았음은 다툼이 없다", []string{"다툼[열쇠:체언]+이[주격:이]", "다툼[열쇠:체언]+이[보격:이]"}, "다투[열쇠:동사]+(ㅁ)[이름:ㅁ]+이[주격:이]"},
		{"그림을 그렸다", []string{"그림[열쇠:체언]+을[목적격:을]"}, "그리[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"},
		{"웃음을 지었다", []string{"웃음[열쇠:체언]+을[목적격:을]"}, "웃[열쇠:동사]+음[이름:음]+을[목적격:을]"},
		{"걸음을 옮겼다", []string{"걸음[열쇠:체언]+을[목적격:을]"}, "걷[열쇠:동사]+음[이름:음]+을[목적격:을]"},
		{"삶을 살았다", []string{"삶[열쇠:체언]+을[목적격:을]"}, "살[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"},
		{"꿈을 꾸었다", []string{"꿈[열쇠:체언]+을[목적격:을]"}, "꾸[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"},
		{"봄에 만났다", []string{"봄[열쇠:체언]+에[부사격:에]"}, "보[열쇠:동사]+(ㅁ)[이름:ㅁ]+에[부사격:에]"},
		{"잠을 잤다", []string{"잠[열쇠:체언]+을[목적격:을]"}, "자[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]"},
		{"마음을 먹었다", []string{"마음[열쇠:체언]+을[목적격:을]"}, ""},
		{"처음에 만났다", []string{"처음[열쇠:체언]+에[부사격:에]"}, ""},
		{"먹기를 바란다", []string{"먹[열쇠:동사]+기[이름:기]+를[목적격:를]"}, ""},
		{"함께 갔다", []string{"함께[열쇠:부사]"}, ""},
		{"있음.", []string{"있[보조 용언:있]+음[이름:음]", "있[열쇠:동사]+음[이름:음]", "있[열쇠:형용사]+음[이름:음]"}, ""},
		{"음을 낸다", []string{"음[열쇠:체언]+을[목적격:을]"}, ""},
	}
	for _, c := range cases {
		nodes := flatNodes(r, c.text, true)
		if !slices.ContainsFunc(nodes, func(n string) bool {
			return slices.ContainsFunc(c.best, func(w string) bool { return strings.HasSuffix(n, "} "+w) })
		}) {
			t.Errorf("%s: none of %v first\n%v", c.text, c.best, nodes)
		}
		if lines := streamText(r, c.text); c.also != "" && !strings.Contains(lines, "\n"+c.also+"\n") {
			t.Errorf("%s: no %s\n%s", c.text, c.also, lines)
		}
	}
	if lines := streamText(r, "감을"); !strings.Contains(lines, "\n감[열쇠:체언]+을[목적격:을]\n") || !strings.Contains(lines, "\n가[열쇠:동사]+(ㅁ)[이름:ㅁ]+을[목적격:을]\n") {
		t.Fatal("감을: the noun next to 가+(ㅁ)+을\n" + lines)
	}
	for _, in := range []string{"학굠을", "함께", "함에"} {
		if lines := streamText(r, in); strings.Contains(lines, droppedName+"]+(ㅁ)") {
			t.Errorf("%s: a dropped copula stem before ㅁ\n%s", in, lines)
		}
	}
	if lines := streamText(r, "학굠을"); strings.Contains(lines, "[이름:ㅁ]") {
		t.Fatal("학굠을: ㅁ after a noun\n" + lines)
	}
}
