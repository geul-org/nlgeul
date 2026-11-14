//ff:func feature=token type=helper control=iteration dimension=1
//ff:what draw — the boxes of the way drawn at the width and stacked top to bottom; last is the line of the title of the last box
package token

func (w *boxWay) draw(width int) {
	w.lines, w.last = nil, 0
	for i, b := range w.boxes {
		if i == len(w.boxes)-1 {
			w.last = len(w.lines)
		}
		w.lines = append(w.lines, b.lines(width)...)
	}
}
