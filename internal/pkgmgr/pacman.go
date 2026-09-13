package pkgmgr

import "strings"

// IsInstalledPacman reports whether pkg is installed via pacman.
func IsInstalledPacman(pkg string) bool {
	return runSilent("pacman", "-Qq", pkg)
}

// SearchPacmanExact returns packages whose name exactly matches app.
func SearchPacmanExact(app string) []string {
	return runLines("pacman", "-Ssq", "^"+app+"$")
}

// SearchPacmanPrefix returns packages whose name starts with app.
func SearchPacmanPrefix(app string) []string {
	return runLines("pacman", "-Ssq", "^"+app)
}

// SearchPacmanAll returns packages matching app anywhere in name/description.
func SearchPacmanAll(app string) []string {
	return runLines("pacman", "-Ssq", app)
}

// VersionPacman returns the repository version of pkg, or "" if unknown.
func VersionPacman(pkg string) string {
	out := runOutput("pacman", "-Si", pkg)
	return fieldValue(out, "Version")
}

// InfoPacman returns Name, Version, Repository and Description for pkg, as
// reported by `pacman -Si`. ok is false if pacman has no info on pkg.
func InfoPacman(pkg string) (name, version, repo, desc string, ok bool) {
	out := runOutput("pacman", "-Si", pkg)
	if strings.TrimSpace(out) == "" {
		return "", "", "", "", false
	}
	return fieldValue(out, "Name"), fieldValue(out, "Version"), fieldValue(out, "Repository"), fieldValue(out, "Description"), true
}

// InstallPacman installs pkg via pacman, attaching to the current terminal.
func InstallPacman(pkg string) bool {
	return runAttached("pacman", "-S", pkg)
}

// RemovePacman removes pkg via pacman, attaching to the current terminal.
func RemovePacman(pkg string) bool {
	return runAttached("pacman", "-R", pkg)
}

// UpdatePacman updates pkg via pacman, attaching to the current terminal.
func UpdatePacman(pkg string) bool {
	return runAttached("pacman", "-S", pkg)
}

// SyncAllPacman runs a full pacman system upgrade.
func SyncAllPacman() bool {
	return runAttached("pacman", "-Syu")
}
