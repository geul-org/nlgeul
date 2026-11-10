//ff:func feature=cli type=command control=iteration dimension=1
//ff:what tokenInput — the sentences nlgeul token reads: the one argument, or with --stdin every line of standard input
package cli

import (
	"bufio"
	"errors"
	"io"
)

func tokenInput(args []string, stdin bool, in io.Reader) ([]string, error) {
	if !stdin && len(args) == 1 {
		return args, nil
	}
	if !stdin || len(args) != 0 {
		return nil, errors.New("문장 하나를 인자로 주거나, 인자 없이 --stdin으로 줄마다 문장을 준다")
	}
	var out []string
	sc := bufio.NewScanner(in)
	sc.Buffer(make([]byte, 1<<20), 1<<24)
	for sc.Scan() {
		out = append(out, sc.Text())
	}
	return out, sc.Err()
}
