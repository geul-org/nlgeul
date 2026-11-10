//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what slotsMark — the slots read before a closing, each a piece: "+었+겠+더"
package parse

func slotsMark(slots []rune) string {
	out := ""
	for _, r := range slots {
		out += "+" + string(r)
	}
	return out
}
