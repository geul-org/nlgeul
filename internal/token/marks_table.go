package token

// The readings of the full stop that join the runs on its two sides
// into one name, and those that end what stands before it. Each is a
// graph of its own (dotGraph). A joining reading's graph is named with
// the sign (파일명:.); a closing reading is named as it is printed, with
// no sign after it (마침표).
var joinNames = []string{"파일명", "도메인", "소수점", "버전"}
var stopNames = []string{"마침표", "줄임표"}

// extensions are the file-name extensions the reading 파일명 knows, and
// topDomains the top-level domains the reading 도메인 knows. They are
// closed tables of common ones, not complete.
var extensions = []string{
	"py", "go", "md", "json", "tsv", "csv", "txt", "yaml", "yml", "toml", "js", "ts", "tsx", "jsx", "html", "css",
	"sh", "c", "h", "cpp", "rs", "java", "kt", "rb", "sql", "xml", "pdf", "png", "jpg", "jpeg", "gif", "svg", "webp",
	"zip", "gz", "tar", "log", "lock", "mod", "sum", "env", "ini", "conf", "ipynb", "docx", "xlsx", "pptx", "hurl",
}
var topDomains = []string{"com", "org", "net", "io", "dev", "ai", "app", "co", "kr", "jp", "uk", "us", "de", "fr", "cn", "edu", "gov", "info", "me", "tv", "xyz", "py"}

// linkSigns are signs that bind two words written with no space into
// one stem: a slash sets two things side by side (입력/출력), a hyphen
// joins two words that belong together (원-달러). Each stands only
// with a letter on both sides, and the stem only when every word in it
// is known (bound). Between two runs of Latin letters and digits the
// same signs join the runs into one name before the unit is read
// (nameSigns).
var linkSigns = map[rune]string{'/': "대비:/", '-': "붙임:-"}
var linkOrder = []rune{'/', '-'}

// run names the graph of a run of Latin letters and digits that no
// table lists. forceRun is the force of the rebuttal "it is in no key
// table" against it: lower than a new Hangul word's, because the
// change of script at its edges is itself a sign of a boundary. It is
// a placeholder, not a measured value.
const (
	runName  = "열:영숫자"
	forceRun = 0.25
)
