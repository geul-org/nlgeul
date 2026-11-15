//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nearOf — a candidate of one word as the rules of a sentence read it: its stem as printed and the stem's name (both "" when its first part is no stem), the name of its last part when that is not the stem, whether one of its parts is a closing (predKinds), and the names of the parts that are not the stem (joined with "+")
package token

import (
	"slices"
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func nearOf(br Branch) rule.Near {
	var n rule.Near
	if len(br.Words) == 0 || len(br.Words[0].Parts) == 0 {
		return n
	}
	parts := br.Words[0].Parts
	if parts[0].Stem {
		n.Stem, n.Kind = parts[0].Text, parts[0].Name.String()
	}
	if last := parts[len(parts)-1]; !last.Stem {
		n.Last = last.Name.String()
	}
	var marks []string
	for _, p := range parts {
		n.Pred = n.Pred || (!p.Stem && slices.Contains(predKinds, p.Name.Kind))
		if !p.Stem {
			marks = append(marks, p.Name.String())
		}
	}
	n.Marks = strings.Join(marks, "+")
	return n
}
