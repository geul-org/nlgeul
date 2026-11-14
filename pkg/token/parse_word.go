//ff:func feature=token type=reader control=sequence
//ff:what parseWord — one line of a word table: the key, its kind, its irregular type (0 for none, or one this layer does not read), its count and the right form of a misspelt key (the fifth column, "" for a line of four columns or an empty fifth)
package token

import (
	"errors"
	"strconv"
	"strings"
)

func parseWord(line string) (string, int, int, int, string, error) {
	cols := strings.Split(line, "\t")
	if (len(cols) != 4 && len(cols) != 5) || cols[0] == "" {
		return "", 0, 0, 0, "", errors.New("칸 넷이나 다섯(검색 키, 갈래, 불규칙 유형, 횟수, 바른 꼴)이 탭으로 나뉘어 있어야 한다")
	}
	count, err := strconv.Atoi(cols[3])
	if err != nil || count < 0 {
		return "", 0, 0, 0, "", errors.New("횟수가 수가 아니다")
	}
	kind, ok := partKinds[cols[1]]
	if !ok {
		kind = kindAlone
	}
	right := ""
	if len(cols) == 5 {
		right = cols[4]
	}
	return cols[0], kind, typeIrr[cols[2]], count, right, nil
}
