//ff:func feature=token type=reader control=iteration dimension=1
//ff:what caseNames — the candidates of a particle, one for each case it marks: 주격, 보격, 목적격, 관형격, 부사격, 호격; 보조사 when it marks none, 접미 for the suffix-like ones
package token

import "github.com/geul-org/nlgeul/pkg/rule"

var caseOrder = []int{rule.CaseSubj, rule.CaseComp, rule.CaseObj, rule.CaseAdn, rule.CaseAdv, rule.CaseVoc}
var caseLabel = map[int]string{rule.CaseSubj: "주격", rule.CaseComp: "보격", rule.CaseObj: "목적격", rule.CaseAdn: "관형격", rule.CaseAdv: "부사격", rule.CaseVoc: "호격"}

func caseNames(kase, slot int) []string {
	var out []string
	for _, c := range caseOrder {
		if kase&c != 0 {
			out = append(out, caseLabel[c])
		}
	}
	if out == nil && slot == 0 {
		return []string{"접미"}
	}
	if out == nil {
		return []string{"보조사"}
	}
	return out
}
