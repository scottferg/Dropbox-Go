package dropbox

const (
	RootSandbox = "sandbox"
	RootDropbox = "dropbox"
)

type Parameters struct {
	Rev            string
	Locale         string
	Overwrite      string
	ParentRev      string
	FileLimit      string
	Hash           string
	List           string
	Cursor         string
	IncludeDeleted string
	RevLimit       string
	ShortUrl       string
	Format         string
	Size           string
	Root           string
	ToPath         string
	FromPath       string
	FromCopyRef    string
	OAuthCallback string
}

type Uri struct {
	Root string
	Path string
}
