//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what caseNumberBranch — a candidate with each word whose stem is a court case number sent up as the number's pieces (caseNumberWord); a candidate with no such word is given back as it is
package parse

func caseNumberBranch(br Branch) Branch {
	out := br
	out.Words = make([]Eojeol, len(br.Words))
	for k, w := range br.Words {
		out.Words[k] = caseNumberWord(w)
	}
	return out
}
