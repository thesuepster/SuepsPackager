package pkgmgr

// Source identifies which backend a package came from.
type Source string

const (
	Pacman  Source = "pacman"
	AUR     Source = "aur"
	Flatpak Source = "flatpak"
)

// Result is a single package hit from a search, carrying enough
// information to display and later install it.
type Result struct {
	Name    string
	Version string
	Source  Source
}
