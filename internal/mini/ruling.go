//ff:type feature=mini type=model
//ff:what ruling — the outcome of judging one word: each constituent's highest verdict, and what stands for marking the word: a particle reading, and the longest tail reading with its length
package mini

type ruling struct {
	verdicts map[string]float64
	chain    bool
	tail     string
	size     int
}
