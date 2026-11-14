//ff:func feature=token type=helper control=iteration dimension=1
//ff:what spacingSet — the spacing rules to turn on as a set, or an error naming a rule that SpacingArticles does not list
package token

import (
	"fmt"
	"slices"
	"strings"
)

func spacingSet(names []string) (map[string]bool, error) {
	out := map[string]bool{}
	for _, n := range names {
		if !slices.Contains(SpacingArticles, n) {
			return nil, fmt.Errorf("띄어쓰기 조항 이름이 아니다: %q (쓸 수 있는 이름: %s)", n, strings.Join(SpacingArticles, ", "))
		}
		out[n] = true
	}
	return out, nil
}
