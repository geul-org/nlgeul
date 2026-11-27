//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what signRun — is a word a run of signs alone with no name: one part, a stem no table names, whose letters are all punctuation marks or symbols (signRune); in a text such a run stands for a name kept back (○○, □□□)
package parse

func signRun(parts []Part) bool {
	if len(parts) != 1 || !parts[0].Stem || !bareOrKept(parts[0]) || parts[0].Text == "" {
		return false
	}
	for _, r := range parts[0].Text {
		if !signRune(r) {
			return false
		}
	}
	return true
}
