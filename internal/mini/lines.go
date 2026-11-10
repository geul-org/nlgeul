//ff:func feature=mini type=rule control=sequence
//ff:what Lines — one line "word:constituent" for each space-separated word of the sentence, "+" before each particle, judged by the mini graph
package mini

// Lines trusts three conditions of the input: the predicate comes
// last, the spacing is right, and every noun carries its particle.
func Lines(sentence string) ([]string, error) {
	g, _ := newGraph()
	return classify(g, sentence)
}
