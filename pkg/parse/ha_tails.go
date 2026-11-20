//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what haTails — the closings written with the ㅎ of a stem's last 하 fused into them (haEnds: 원치, 가능케) that end at j with a letter before them, judged by the 하 they took in (the syllable they are judged by is the stem's, as irrTails' are) and, unless 요 follows them (more), by what stands right after: each stands as a tail that asks for a stem whose 하 is so dropped (irrHa), its piece the letters as they stand under the plain closing's name
package parse

// Such a closing takes no slot before it. more tells that 요 follows.
func (sc *scan) haTails(j int, more bool) []tail {
	if j < 1 {
		return nil
	}
	var out []tail
	for _, e := range sc.ix.haEnds[sc.s[j-1]] {
		start, p, ok := sc.place(e, j, more)
		if !ok || start < 1 {
			continue
		}
		p.Prev = '하'
		if !more {
			p.Right = sc.right(j)
		}
		if v := sc.judge(e, start, j, p); v > 0 {
			out = append(out, tail{start: start, pieces: []piece{sc.entryPiece(e, string(sc.s[start:j]), v)}, head: headHa, irr: irrHa, noYo: e.noYo, close: e})
		}
	}
	return out
}
