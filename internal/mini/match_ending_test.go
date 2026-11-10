//ff:func feature=mini type=reader control=sequence
//ff:what TestMatchEnding — each way of joining reaches its reader: 먹다, 봐, 간다
package mini

import "testing"

func TestMatchEnding(t *testing.T) {
	if got, ok := matchEnding([]rune("먹다"), 2, ending{form: "다"}, false); !ok || got.closing != "다" {
		t.Fatal("plain")
	}
	if got, ok := matchEnding([]rune("봐"), 1, ending{lead: leadA}, false); !ok || got.closing != "어" {
		t.Fatal("a")
	}
	if got, ok := matchEnding([]rune("간다"), 2, ending{form: "다", lead: jongN}, false); !ok || got.closing != "ㄴ다" {
		t.Fatal("jamo")
	}
}
