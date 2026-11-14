//ff:func feature=parse type=reader control=sequence
//ff:what MarkerOf — a closing's row read from four columns of a table's line, beginning at the column given: its form, what must stand before it, its kinds and its ways of attaching; a form with a lead takes no condition on what stands before it
package parse

import "github.com/geul-org/nlgeul/data"

func MarkerOf(row data.Row, at int) (Marker, error) {
	lead, form, err := FormOf(row, at)
	if err != nil {
		return Marker{}, err
	}
	after, err := row.Pick(at+1, AfterNames)
	if err != nil {
		return Marker{}, err
	}
	if lead != leadNone && after != afterAny {
		return Marker{}, row.Errf(at+1, "받침으로 붙는 꼴과 어/아 꼴은 이 칸에 -를 적는다")
	}
	kind, err := row.Flags(at+2, KindNames)
	if err != nil {
		return Marker{}, err
	}
	if kind == 0 {
		return Marker{}, row.Errf(at+2, "갈래가 하나는 있어야 한다")
	}
	flags, err := row.Flags(at+3, FlagNames)
	if err != nil {
		return Marker{}, err
	}
	return Marker{
		Form: form, Lead: lead, After: after, Kind: kind,
		Eu: flags&fEu != 0, NoYo: flags&fNoYo != 0, NoIda: flags&fNoIda != 0,
		NeedIda: flags&fNeedIda != 0, SelfIda: flags&fSelfIda != 0,
	}, nil
}
