//ff:type feature=parse type=model
//ff:what ruling — the outcome of judging one word: each constituent's highest verdict, and what stands for marking the word: a particle reading, and the longest tail reading with its length
package parse

type ruling struct {
	verdicts map[string]float64
	chain    bool
	tail     string
	size     int
}
