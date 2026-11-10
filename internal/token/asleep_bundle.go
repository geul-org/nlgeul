//ff:func feature=token type=reader control=iteration dimension=1
//ff:what asleepBundle — a unit's readings that are asleep as one bundle: the unit's letters, and each reading whole as a branch marked asleep, with the reading's own verdict and preference; none when no reading is asleep
package token

func (sc *scan) asleepBundle(asleep []*path) []Bundle {
	if len(asleep) == 0 {
		return nil
	}
	b := Bundle{Text: string(sc.s)}
	for _, p := range asleep {
		br := stretch(p, 0, len(sc.s))
		br.Verdict, br.Pref, br.Asleep = p.verdict, p.pref, true
		b.Branches = append(b.Branches, br)
	}
	return []Bundle{b}
}
