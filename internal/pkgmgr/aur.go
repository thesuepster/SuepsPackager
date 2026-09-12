package pkgmgr

import "strings"

// HasParu reports whether the paru AUR helper is installed.
func HasParu() bool {
	return CommandExists("paru")
}

// filterParuNoise drops paru's "::" progress/status lines from search output.
func filterParuNoise(lines []string) []string {
	var out []string
	for _, line := range lines {
		if strings.HasPrefix(line, "::") {
			continue
		}
		out = append(out, line)
	}
	return out
}

// SearchAURExact returns AUR packages whose name exactly matches app.
func SearchAURExact(app string) []string {
	return filterParuNoise(runLines("paru", "-Ssq", "--aur", "^"+app+"$"))
}

// SearchAURPrefix returns AUR packages whose name starts with app.
func SearchAURPrefix(app string) []string {
	return filterParuNoise(runLines("paru", "-Ssq", "--aur", "^"+app))
}

// SearchAURAll returns AUR packages matching app anywhere in name/description.
func SearchAURAll(app string) []string {
	return filterParuNoise(runLines("paru", "-Ssq", "--aur", app))
}

// VersionAUR returns the AUR version of pkg, or "" if unknown.
func VersionAUR(pkg string) string {
	out := runOutput("paru", "-Si", pkg)
	return fieldValue(out, "Version")
}

// InfoAUR returns Name, Version and Description for pkg, as reported by
// `paru -Si`. ok is false if paru has no info on pkg.
func InfoAUR(pkg string) (name, version, desc string, ok bool) {
	out := runOutput("paru", "-Si", pkg)
	if strings.TrimSpace(out) == "" {
		return "", "", "", false
	}
	return fieldValue(out, "Name"), fieldValue(out, "Version"), fieldValue(out, "Description"), true
}

// InstallAUR installs pkg via paru, running as asUser (never root) and
// attaching to the current terminal.
func InstallAUR(pkg, asUser string) bool {
	return runAttached("sudo", "-u", asUser, "paru", "-S", pkg)
}

// UpdateAUR updates pkg via paru, attaching to the current terminal.
func UpdateAUR(pkg string) bool {
	return runAttached("paru", "-S", pkg)
}

// SyncAllAUR runs a full paru system upgrade (pacman + AUR).
func SyncAllAUR() bool {
	return runAttached("paru", "-Syu")
}
