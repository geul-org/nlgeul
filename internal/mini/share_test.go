//ff:func feature=mini type=helper control=sequence
//ff:what TestShare — the rival's part of the two counts on [-1, 1]; no value when either count is unknown
package mini

import "testing"

func TestShare(t *testing.T) {
	if share(1, 3) != 0.5 || share(5, 5) != 0 || share(0, 3) != -1 || share(3, 0) != -1 {
		t.Fatal("share")
	}
}
