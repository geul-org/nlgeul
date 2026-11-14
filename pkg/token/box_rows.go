//ff:func feature=token type=helper control=iteration dimension=1
//ff:what boxRows — the rows inside a node's box: each candidate of its contest shown (the first head, 0 for all) as its letters, "성립" with its verdict and "선호" with its preference, the letters padded to the widest and the preferences set right; a last row "(k개 생략됨)" counts those left out
package token

import "fmt"

func boxRows(n Node, head int) []string {
	shown := n.Contest.Branches
	if head > 0 && len(shown) > head {
		shown = shown[:head]
	}
	texts := make([]string, len(shown))
	prefs := make([]string, len(shown))
	for i, b := range shown {
		texts[i] = b.String()
		prefs[i] = fmt.Sprintf("%.2f", b.Pref)
	}
	cw, pw := maxCells(texts), maxCells(prefs)
	rows := make([]string, 0, len(shown)+1)
	for i, b := range shown {
		rows = append(rows, fmt.Sprintf("%s  성립 %.2f 선호 %s", padCells(texts[i], cw), b.Verdict, padCellsLeft(prefs[i], pw)))
	}
	if left := len(n.Contest.Branches) - len(shown); left > 0 {
		rows = append(rows, fmt.Sprintf("(%d개 생략됨)", left))
	}
	return rows
}
