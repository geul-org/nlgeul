//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what plateRegion — how many letters at a place of the unit are the name of a region that an older registration plate begins with (plateRegions: 서울, 경기), 0 when none stands there
package parse

func (sc *scan) plateRegion(from int) int {
	for _, region := range plateRegions {
		r := []rune(region)
		if from+len(r) <= len(sc.s) && string(sc.s[from:from+len(r)]) == region {
			return len(r)
		}
	}
	return 0
}
