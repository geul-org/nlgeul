//ff:func feature=token type=reader control=sequence
//ff:what typoClosingOf — one closing written wrong read from a line of its table: the letters it is written with (a form of whole syllables, or one under a final consonant) and the name of the right closing
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/pkg/parse"
)

func typoClosingOf(row data.Row) (typoClosing, error) {
	lead, form, err := parse.FormOf(row, 0)
	if err != nil {
		return typoClosing{}, err
	}
	if lead < 0 || form == "" {
		return typoClosing{}, row.Errf(0, "틀린 꼴은 한글 음절이나 받침 뒤에 한글 음절로 적는다")
	}
	if !strings.Contains(row.Cells[1], ":") {
		return typoClosing{}, row.Errf(1, "바른 맺음은 색인의 이름(맺음:ㄹ게, 이름:음)으로 적는다")
	}
	return typoClosing{form, lead, row.Cells[1]}, nil
}
