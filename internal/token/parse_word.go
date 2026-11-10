//ff:func feature=token type=reader control=sequence
//ff:what parseWord — one line of a word table: the key, its kind and its count
package token

import (
	"errors"
	"strconv"
	"strings"
)

func parseWord(line string) (string, int, int, error) {
	cols := strings.Split(line, "\t")
	if len(cols) != 4 || cols[0] == "" {
		return "", 0, 0, errors.New("칸 넷(검색 키, 갈래, 불규칙 유형, 횟수)이 탭으로 나뉘어 있어야 한다")
	}
	count, err := strconv.Atoi(cols[3])
	if err != nil || count < 0 {
		return "", 0, 0, errors.New("횟수가 수가 아니다")
	}
	kind, ok := partKinds[cols[1]]
	if !ok {
		kind = kindAlone
	}
	return cols[0], kind, count, nil
}
