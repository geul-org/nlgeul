//ff:type feature=rule type=model
//ff:what Sentence — the facts of the whole sentence that a word's rules read: the stems of the last word's tail readings, and how many earlier words carry 이/가
package rule

type Sentence struct {
	Stems      []string
	Nominative int
}
