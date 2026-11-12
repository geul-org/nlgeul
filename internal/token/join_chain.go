//ff:func feature=token type=reader control=iteration dimension=1
//ff:what joinChain — judge each sign of a chain of runs and note the atoms: runs stay joined across a sign that a joining reading stands for; an atom made of runs of digits is a number, and so is a run that is 0x and hexadecimal digits (named 수:16진법)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) joinChain(chain [][2]int) {
	labels, signs := make([]string, len(chain)), make([]rune, len(chain)-1)
	for i, r := range chain {
		labels[i] = string(sc.s[r[0]:r[1]])
		if i < len(signs) {
			signs[i] = sc.s[r[1]]
		}
	}
	start, names, num := chain[0][0], []string(nil), true
	for k, r := range chain {
		num = num && sideOf(labels[k])&rule.SideDigit != 0
		var joined []markRead
		if k < len(signs) {
			joined = sc.judgeMark(sc.ix.signs[signs[k]], signFacts(labels, signs, k), string(signs[k]))
		}
		if len(joined) > 0 {
			names = union(names, markNames(shownNames(joined)))
			continue
		}
		if len(names) == 0 && start == r[0] && hexRun(labels[k]) {
			sc.closeAtom(start, r[1], []string{hexName}, true)
		} else {
			sc.closeAtom(start, r[1], names, num)
		}
		start, names, num = r[1]+1, nil, true
	}
}
