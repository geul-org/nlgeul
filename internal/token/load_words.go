//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadWords — read a word table: tab-separated key, part of speech, irregular type, count and (for a misspelt key) the right form on each line, lines that begin with # skipped; a key's candidates are folded into one entry, and a line with a right form goes to the misspelt forms instead
package token

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The count of a key is its greatest. The sum of every line's count
// comes back for preference. A line with a right form is not a word of
// the table: it is kept apart, by its written form, with the kind of
// the line (kinds of lines with the same two forms are joined), and
// adds to neither the entries nor the sum.
func loadWords(path string) (map[string]word, map[string][]typo, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, 0, fmt.Errorf("낱말 표를 열 수 없다: %w", err)
	}
	defer f.Close()
	lex, typos, total, line := map[string]word{}, map[string][]typo{}, 0.0, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line++
		if strings.HasPrefix(sc.Text(), "#") || sc.Text() == "" {
			continue
		}
		key, kind, irr, count, right, err := parseWord(sc.Text())
		if err != nil {
			return nil, nil, 0, fmt.Errorf("낱말 표 %d번째 줄: %w", line, err)
		}
		if right != "" {
			typos[key] = typoJoined(typos[key], right, kind)
			continue
		}
		lex[key] = folded(lex[key], kind, irr, count)
		total += float64(count)
	}
	return lex, typos, total, sc.Err()
}
