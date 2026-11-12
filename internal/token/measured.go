//ff:func feature=token type=reader control=iteration dimension=1
//ff:what measured — the bundles of a sign of a pair alone, when it stands right after a digit and is a sign of the unit table (3", 5'), with a branch added for each line of the table that has it: the sign as a unit (단위:인치), sure, with the preference of a unit; after the readings as a sign of a pair and the sign as its letter alone
package token

func (ix *index) measured(bundles []Bundle, u unit) []Bundle {
	if !u.measure || len(bundles) == 0 {
		return bundles
	}
	sign := string([]rune(u.text)[:u.head])
	for _, row := range unitRows {
		if row.Cells[0] != sign {
			continue
		}
		p := Part{Text: sign, Names: namesOf(unitLabel + row.Cells[1])}
		bundles[0].Branches = append(bundles[0].Branches, Branch{Words: []Eojeol{{Parts: []Part{p}, End: u.head}}, Verdict: 1, Pref: ix.afterPrefs[false]})
	}
	return bundles
}
