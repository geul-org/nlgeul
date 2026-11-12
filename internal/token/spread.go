//ff:func feature=token type=reader control=iteration dimension=1
//ff:what spread — the bundles of a unit whose head is a sign of a pair, with each reading that names the head given once for each way the sign is read: for each depth it may open or close at (([열기:1], ([열기:2]), and with the head as its letter alone when the sign is also read so; a reading that does not name the head (the whole unit as a new word) stays as it is; the bundles of any other unit are given back as they are
package token

import "strconv"

// The readings that differ only in how the head is read have the same
// verdict and preference.
func spread(bundles []Bundle, u unit) []Bundle {
	if u.name != pairName || len(bundles) == 0 {
		return bundles
	}
	var out []Branch
	for _, br := range bundles[0].Branches {
		if !headed(br) {
			out = append(out, br)
			continue
		}
		for _, rd := range u.reads {
			out = append(out, withHead(br, Name{Kind: rd.name, Form: strconv.Itoa(rd.depth)}, rd.depth))
		}
		if u.bare {
			out = append(out, withHead(br, Name{}, 0))
		}
	}
	bundles[0].Branches = out
	return bundles
}
