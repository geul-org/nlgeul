//ff:func feature=token type=reader control=iteration dimension=1
//ff:what spread — the contests of a unit whose head is a sign of a pair, with each candidate that names the head given once for each way the sign is read: for each depth it may open or close at (([열기:1], ([열기:2]), and with the head as its letter alone when the sign is also read so; a candidate that does not name the head (the whole unit as a new word) stays as it is; the contests of any other unit are given back as they are
package token

import "strconv"

// The candidates that differ only in how the head is read have the same
// verdict and preference.
func spread(contests []Contest, u unit) []Contest {
	if u.name != pairName || len(contests) == 0 {
		return contests
	}
	var out []Branch
	for _, br := range contests[0].Branches {
		if !headed(br) {
			out = append(out, br)
			continue
		}
		for _, rd := range u.candidates {
			out = append(out, withHead(br, Name{Kind: rd.name, Form: strconv.Itoa(rd.depth)}, rd.depth))
		}
		if u.bare {
			out = append(out, withHead(br, Name{}, 0))
		}
	}
	contests[0].Branches = out
	return contests
}
