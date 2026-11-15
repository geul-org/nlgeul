//ff:func feature=token type=rule control=iteration dimension=1
//ff:what addHa — the closings written with the ㅎ of a stem's last 하 fused into their first consonant (haPairs, clause 40: 치, 케, 타, 토록, 코자, 컨대), each a copy of the plain closing's entry under the written form, kept by its last letter (haEnds) and not under the index's keys: read only with a word table, turned down by a particle right after it (wide), with no count of its own; a plain closing the index does not have is passed over
package token

// The copy keeps the plain closing's name, class and ways of attaching.
// Its count is 0: its preference is what its syllables would cost as
// stem syllables, not the plain closing's (pref). The plain closing is
// found under the last letter of its name, which is its form's.
func (ix *index) addHa() {
	ix.haEnds = map[rune][]*entry{}
	for _, pr := range haPairs {
		name := []rune(pr[0])
		plain := ix.named(name[len(name)-1], pr[0])
		if plain.g == nil {
			continue
		}
		e := *plain
		e.form, e.wide, e.count = pr[1], true, 0
		ix.edges += markerGraph(&e, nil, ix.lexed)
		ix.graphs++
		k := entryKey(&e)
		ix.haEnds[k] = append(ix.haEnds[k], &e)
	}
}
