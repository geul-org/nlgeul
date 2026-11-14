//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what writeTokenJSON — every sentence as one line of JSON (JSON Lines, UTF-8, the Korean letters as they are), numbered from 0
package cli

import (
	"encoding/json"
	"io"

	"github.com/geul-org/nlgeul/pkg/token"
)

func writeTokenJSON(out io.Writer, r *token.Reader, sentences []string) error {
	enc := json.NewEncoder(out)
	enc.SetEscapeHTML(false)
	for i, s := range sentences {
		if err := enc.Encode(r.JSONSentence(s, i)); err != nil {
			return err
		}
	}
	return nil
}
