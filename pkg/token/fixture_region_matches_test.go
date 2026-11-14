//ff:func feature=token type=helper control=iteration dimension=1
//ff:what regionMatches — the regions a stream has sent against the contests Reader.Contests gives: as many, each region's letters, unit and Glued those of its contest, and the node that is the whole region holds the candidates of one word the contest holds (letters, verdict, preference); the names of what does not match
package token

func regionMatches(st *Stream, contests []Contest) []string {
	if len(st.sent) != len(contests) {
		return []string{"count"}
	}
	var bad []string
	for i, rg := range st.sent {
		whole, ok := wholeValues(rg)
		first := rg.nodes[0].Contest
		if !ok || !sameStrings(whole, branchValues(contests[i])) {
			bad = append(bad, "whole "+contests[i].Text)
		}
		if first.Unit != contests[i].Unit || first.Glued != contests[i].Glued || rg.span.End-rg.span.Start != len([]rune(contests[i].Text)) {
			bad = append(bad, "head "+contests[i].Text)
		}
	}
	return bad
}
