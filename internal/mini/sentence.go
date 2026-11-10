//ff:type feature=mini type=model
//ff:what sentence — the facts of the whole sentence that a word's rules read: the stems of the last word's tail readings, and how many earlier words carry 이/가
package mini

type sentence struct {
	stems      []string
	nominative int
}
