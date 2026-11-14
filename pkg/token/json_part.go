//ff:type feature=token type=model
//ff:what JSONPart — one part of a candidate as nlgeul token --format json writes it: its letters, its name (Kind:Form, with the form it is written in when that is wrong; null when no table gives one), whether a table knows it (false for a stem no table knows, a new word), whether it leaves no trace in the letters (restored), what the letters do not show (hidden), and whether it is the stem
package token

type JSONPart struct {
	Text     string  `json:"text"`
	Name     *string `json:"name"`
	Known    bool    `json:"known"`
	Restored bool    `json:"restored"`
	Hidden   string  `json:"hidden"`
	Stem     bool    `json:"stem"`
}
