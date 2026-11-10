//ff:func feature=token type=reader control=iteration dimension=1
//ff:what loadWords — read a word table: tab-separated key, part of speech, irregular type and count on each line, lines that begin with # skipped; a key's readings are folded into one entry
package token

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// The count of a key is its greatest. The sum of every line's count
// comes back for preference.
func loadWords(path string) (map[string]word, float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, fmt.Errorf("낱말 표를 열 수 없다: %w", err)
	}
	defer f.Close()
	lex, total, line := map[string]word{}, 0.0, 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line++
		if strings.HasPrefix(sc.Text(), "#") || sc.Text() == "" {
			continue
		}
		key, w, err := parseWord(sc.Text())
		if err != nil {
			return nil, 0, fmt.Errorf("낱말 표 %d번째 줄: %w", line, err)
		}
		old := lex[key]
		lex[key] = word{old.kinds | w.kinds, max(old.count, w.count)}
		total += float64(w.count)
	}
	return lex, total, sc.Err()
}
