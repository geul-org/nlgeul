//ff:func feature=parse type=reader control=sequence
//ff:what ClosedWords — the closed sets of words the tables list, by kind, for other packages
package parse

func ClosedWords() map[string][]string {
	return map[string][]string{
		"의존명사":  boundNouns,
		"보조 용언": helperStems,
		"관형사":   determiners,
		"접속 부사": conjunctions,
	}
}
