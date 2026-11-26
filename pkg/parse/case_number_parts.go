//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what caseNumberParts — the contests of a unit with each word whose stem is a court case number (one part named caseKind) given the number's three pieces in its place, and each word whose stem is a motor vehicle's registration number (one part named plateKind) given that number's pieces (caseNumberBranch); the order of the branches and what they hold are kept, and nothing is changed in the contests given
package parse

func caseNumberParts(contests []Contest) []Contest {
	out := make([]Contest, len(contests))
	for i, c := range contests {
		out[i] = c
		out[i].Branches = make([]Branch, len(c.Branches))
		for k, br := range c.Branches {
			out[i].Branches[k] = caseNumberBranch(br)
		}
	}
	return out
}
