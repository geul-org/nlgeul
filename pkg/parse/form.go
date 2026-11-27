//ff:type feature=parse type=model
//ff:what form — one line of the tables of special letters (special-etc.tsv, special-forms.tsv): a letter that writes other letters in a shape of its own (① is 1 in a circle), with that shape (꼴: 원), what the letters it writes are (내용: 수), those letters (값: 1) and the number it stands for ("" when it stands for none)
package parse

type form struct {
	letter  rune
	shape   string
	content string
	value   string
	number  string
}
