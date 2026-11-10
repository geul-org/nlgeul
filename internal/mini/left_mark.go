//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what leftMark — the text before 하다·되다·시키다, divided when it reads as a quotation (가셨다고 → 가+시+었+다+고), else as it is
package mini

import "slices"

func leftMark(left []rune) string {
	best := tailFact{marked: string(left)}
	for _, t := range collectTails(left, len(left), false) {
		if slices.Contains(quoteClosings, t.closing) && t.size > best.size {
			best = t
		}
	}
	return best.marked
}
