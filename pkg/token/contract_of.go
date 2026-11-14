//ff:func feature=token type=reader control=sequence
//ff:what contractOf — one fused form read from a line of its table: the form, the word it stands for, the particle under it as it is printed and as it is named (its entry is found by the last letter of the name), whether it is read only right after a modifying closing, and the letters its word is written with when they are a wrong way to write it
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/data"
)

func contractOf(row data.Row) (contract, error) {
	if _, err := row.Pick(0, contractSets); err != nil {
		return contract{}, err
	}
	name := []rune(row.Cells[4])
	if !strings.Contains(row.Cells[4], ":") || name[len(name)-1] == ':' {
		return contract{}, row.Errf(4, "조사 이름은 색인의 이름(주격:이, 보조사:는)으로 적는다")
	}
	place := row.Opt(5)
	if place != "" && place != adnPlace {
		return contract{}, row.Errf(5, "\"%s\"이나 -를 적는다", adnPlace)
	}
	return contract{row.Cells[1], row.Cells[2], row.Cells[3], row.Cells[4], name[len(name)-1], place != "", row.Opt(6)}, nil
}
