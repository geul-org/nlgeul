//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestClosingNamesOf — the linking closings of the closing table hold -고 (연결|종결 too: 지), -면, the fused 어 and -어야 and none of the modifying, naming or final-only ones (꾸밈:는, 이름:기, 맺음:다); the final ones hold 맺음:다 and not -면
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestClosingNamesOf(t *testing.T) {
	links := closingNamesOf(rule.KindLink)
	for _, name := range []string{"맺음:고", "맺음:지", "맺음:면", "맺음:으면", "맺음:어(합침)", "맺음:어야"} {
		if !slices.Contains(links, name) {
			t.Errorf("%s is a linking closing", name)
		}
	}
	for _, name := range []string{"꾸밈:는", "이름:기", "맺음:다"} {
		if slices.Contains(links, name) {
			t.Errorf("%s is no linking closing", name)
		}
	}
	finals := closingNamesOf(rule.KindFinal)
	if !slices.Contains(finals, "맺음:다") || slices.Contains(finals, "맺음:면") {
		t.Error("the final closings", len(finals))
	}
}
