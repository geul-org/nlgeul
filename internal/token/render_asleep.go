//ff:func feature=token type=reader control=iteration dimension=1
//ff:what renderAsleep — the lines of a unit's readings that are asleep: a head line with the unit and how many there are, then each reading whole with its verdict and preference; nothing when none is asleep
package token

import "fmt"

func renderAsleep(text string, asleep []*path) []string {
	if len(asleep) == 0 {
		return nil
	}
	out := []string{fmt.Sprintf("%s\t잠듦(%d)", text, len(asleep))}
	for _, p := range asleep {
		out = append(out, fmt.Sprintf("  %s\t성립 %.2f 선호 %.2f", p.sig, p.verdict, p.pref))
	}
	return out
}
