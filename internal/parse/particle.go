//ff:type feature=parse type=model
//ff:what particle — one particle form, what must precede it, its place in a chain, the case it marks, and its observed count in particles.md
package parse

type particle struct {
	form  string
	after int
	slot  int
	kase  int
	count int
}
