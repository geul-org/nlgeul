//ff:func feature=token type=reader control=sequence
//ff:what pairOf — one sign of a pair read from a line of its table: the sign, the opening sign of its kind, whether it opens and closes, whether it is unsure, and what keeps it from being such a sign
package token

import "github.com/geul-org/nlgeul/data"

func pairOf(row data.Row) (rune, pairRow, error) {
	sign, err := row.Rune(0)
	if err != nil {
		return 0, pairRow{}, err
	}
	kind, err := row.Rune(1)
	if err != nil {
		return 0, pairRow{}, err
	}
	role, err := row.Flags(2, pairRoles)
	if err != nil {
		return 0, pairRow{}, err
	}
	if role == 0 {
		return 0, pairRow{}, row.Errf(2, "열기와 닫기 가운데 하나는 있어야 한다")
	}
	bars, err := row.Flags(3, pairBars)
	if err != nil {
		return 0, pairRow{}, err
	}
	return sign, pairRow{kind: kind, opens: role&pairOpens != 0, shuts: role&pairShuts != 0, unsure: bars&pairUnsure != 0, bars: bars &^ pairUnsure}, nil
}
