//ff:func feature=parse type=reader control=sequence
//ff:what plateEnd — where the shape of a motor vehicle's registration number ends that begins at a place of the unit (0 when none begins there): the number of the kind of vehicle (plateHeadLens digits, or plateRegionLens digits after the name of a region: plateRegion), a letter of plateLetters and a serial number of plateSerialLen digits that no digit and no Latin letter runs on from; not when the letter is a greater number (60조5000), nor when it is a piece read after a number and such a piece is read right after the serial but for 호, the unit a plate is numbered with (100주5000원: a measure written with no space)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) plateEnd(from int) int {
	region, lens := sc.plateRegion(from), plateHeadLens
	if region > 0 {
		lens = plateRegionLens
	}
	head := digitsFrom(sc.s, from+region)
	at := from + region + head
	letter := string(sc.at(at))
	end := at + 1 + digitsFrom(sc.s, at+1)
	if !slices.Contains(lens, head) || !slices.Contains(plateLetters, letter) || end-at-1 != plateSerialLen || rule.RunRune(sc.at(end)) {
		return 0
	}
	if sc.readsAfter(letter, placeBig) || (sc.readsAfter(letter, 0) && sc.afterAt(end) && string(sc.at(end)) != plateUnit) {
		return 0
	}
	return end
}
