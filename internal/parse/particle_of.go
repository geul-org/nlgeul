//ff:func feature=parse type=reader control=sequence
//ff:what particleOf — one row of the particle table read from a line: the form in Hangul syllables, what must stand before it, its place in a chain (0 to 5), the cases it marks (- for none) and its observed count
package parse

import "github.com/geul-org/nlgeul/data"

func particleOf(row data.Row) (particle, error) {
	lead, form, err := FormOf(row, 0)
	if err != nil {
		return particle{}, err
	}
	if lead != leadNone || form == "" {
		return particle{}, row.Errf(0, "조사의 꼴은 한글 음절로만 적는다")
	}
	after, err := row.Pick(1, AfterNames)
	if err != nil {
		return particle{}, err
	}
	slot, err := row.Count(2)
	if err != nil {
		return particle{}, err
	}
	if slot > maxSlot {
		return particle{}, row.Errf(2, "0에서 %d까지의 수가 아니다", maxSlot)
	}
	kase, err := row.Flags(3, CaseNames)
	if err != nil {
		return particle{}, err
	}
	count, err := row.Count(4)
	if err != nil {
		return particle{}, err
	}
	return particle{form, after, slot, kase, count}, nil
}
