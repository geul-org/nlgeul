//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNounSuffixOf — a line of the table of suffixes read after a nominal is read into a particle named 접미:<form>, in the place of a suffix in chains, with what must stand before it; the table held in the program gives 률 after a final consonant other than ㄴ; a form that is not whole Hangul syllables or is empty, and a name outside the set, are errors that name the column
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestNounSuffixOf(t *testing.T) {
	e, err := nounSuffixOf(rowFor("률", "ㄴ 아닌 받침 뒤"))
	if err != nil || e.name != "접미:률" || e.form != "률" || e.class != classParticle || e.after != rule.AfterNotN || e.after != parse.AfterNames["ㄴ 아닌 받침 뒤"] || e.slot != rule.LinkSuffix || e.count != 0 || e.gate != nil || e.lead != 0 {
		t.Fatal(e, err)
	}
	if len(nounSuffixes) != 1 || nounSuffixes[0].name != "접미:률" || nounSuffixes[0].after != rule.AfterNotN {
		t.Fatal("the table held in the program", nounSuffixes)
	}
	for _, c := range []struct {
		row   []string
		parts []string
	}{
		{[]string{"a", "-"}, []string{"1번째 칸(꼴)", "한글 음절"}},
		{[]string{"ㄴ률", "-"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"", "-"}, []string{"1번째 칸(꼴)", "한글 음절로만"}},
		{[]string{"률", "ㄴ 뒤"}, []string{"2번째 칸(앞말)", "쓸 수 있는 이름", "ㄴ 아닌 받침 뒤"}},
	} {
		_, err := nounSuffixOf(rowFor(c.row...))
		errHas(t, err, c.parts...)
	}
}
