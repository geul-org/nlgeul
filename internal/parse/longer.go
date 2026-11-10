//ff:func feature=parse type=helper control=sequence
//ff:what longer — a yes/no rule value for "a rival reading reads more of the word's end than this one"
package parse

func longer(mine, rival int) float64 {
	return applies(rival > mine)
}
