//ff:func feature=parse type=helper control=sequence
//ff:what suffixedRun — is the word before a tail a run of Latin letters and digits that no table lists (runName: OTS용, ABC사): such a run names a thing as a nominal does, so the gate of a suffix passes it as one
package parse

func suffixedRun(e step) bool {
	return e.word.Parts[0].Name.Kind == labelOf(runName)
}
