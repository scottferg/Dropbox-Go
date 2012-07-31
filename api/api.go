package api

const (
	RootSandbox = "sandbox"
	RootDropbox = "dropbox"
)

type Uri struct {
	Root string
	Path string
}
