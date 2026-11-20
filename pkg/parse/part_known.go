//ff:func feature=parse type=helper control=sequence
//ff:what Known — whether a table knows a part: false only for a stem no table knows (a new word, labelled 새 낱말); a marker, a sign and a stem of a table or a closed set are known. It is the known of nlgeul parse --format json
package parse

func (p Part) Known() bool {
	return p.Label() != newWord
}
