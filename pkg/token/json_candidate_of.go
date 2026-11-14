//ff:func feature=token type=helper control=iteration dimension=2
//ff:what jsonCandidateOf — a candidate (a branch) as it is written in JSON: the parts of its words in order, its verdict, its preference and its lambda
package token

func jsonCandidateOf(br Branch) JSONCandidate {
	out := JSONCandidate{Parts: []JSONPart{}, Verdict: br.Verdict, Pref: br.Pref, Lambda: jsonFloat(lambdaOf(br))}
	for _, w := range br.Words {
		for _, p := range w.Parts {
			out.Parts = append(out.Parts, jsonPartOf(p))
		}
	}
	return out
}
