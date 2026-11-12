//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what FormOf — a form as a table writes it, read into its lead and its syllables: ㄴ다 is 다 under the final consonant ㄴ, 어/아서 is 서 after 어/아/여, 면 is whole syllables; what follows the lead must be Hangul syllables, and a form with no lead must not be empty
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/data"
)

func FormOf(row data.Row, col int) (int, string, error) {
	lead, form := leadNone, row.Cells[col]
	for name, v := range leadNames {
		if strings.HasPrefix(form, name) {
			lead, form = v, strings.TrimPrefix(form, name)
		}
	}
	for _, r := range form {
		if r < '가' || r > '힣' {
			return 0, "", row.Errf(col, "꼴은 한글 음절로 적고, 받침으로 붙는 꼴은 맨 앞에 ㄴ·ㄹ·ㅁ·ㅂ·ㅆ을, 어/아로 시작하는 꼴은 맨 앞에 %s를 적는다", vowelLead)
		}
	}
	return lead, form, nil
}
