//ff:func feature=parse type=reader control=sequence
//ff:what TestParticleMarkers — every particle row comes out with its fields
package parse

import "testing"

func TestParticleMarkers(t *testing.T) {
	got := ParticleMarkers()
	if len(got) != len(particles) || got[0].Form != "이" || got[0].After != afterBatchim || got[0].Kase != caseSubj|caseComp || got[0].Count != 5530000 {
		t.Fatal(len(got), got[0])
	}
}
