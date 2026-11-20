//ff:func feature=parse type=helper control=sequence
//ff:what pop — take the first walk out of the frontier (it holds at least one)
package parse

func (f *frontier) pop() *walk {
	w := f.items[0]
	f.items = f.items[1:]
	return w
}
