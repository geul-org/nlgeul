//ff:func feature=mini type=reader control=iteration dimension=1
//ff:what slotsMark — the slots read before a closing, each a piece: "+었+겠+더"
package mini

func slotsMark(slots []rune) string {
	out := ""
	for _, r := range slots {
		out += "+" + string(r)
	}
	return out
}
