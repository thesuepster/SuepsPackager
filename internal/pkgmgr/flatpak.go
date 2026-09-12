package pkgmgr

import "strings"

// HasFlatpak reports whether the flatpak tool is installed.
func HasFlatpak() bool {
	return CommandExists("flatpak")
}

// IsInstalledFlatpak reports whether appID is installed via flatpak.
func IsInstalledFlatpak(appID string) bool {
	for _, line := range runLines("flatpak", "list", "--columns=application") {
		if line == appID {
			return true
		}
	}
	return false
}

// skipHeader drops the column-header row flatpak prints as the first line
// of tabular output.
func skipHeader(lines []string) []string {
	if len(lines) == 0 {
		return nil
	}
	return lines[1:]
}

// SearchFlatpakExact returns flatpak application IDs that exactly match
// app, case-insensitively.
func SearchFlatpakExact(app string) []string {
	var out []string
	for _, line := range skipHeader(runLines("flatpak", "search", "--columns=application", app)) {
		if line == "Application ID" {
			continue
		}
		if strings.EqualFold(line, app) {
			out = append(out, line)
		}
	}
	return out
}

// SearchFlatpakPrefix returns flatpak application IDs that start with app,
// case-insensitively.
func SearchFlatpakPrefix(app string) []string {
	var out []string
	prefix := strings.ToLower(app)
	for _, line := range skipHeader(runLines("flatpak", "search", "--columns=application", app)) {
		if line == "Application ID" {
			continue
		}
		if strings.HasPrefix(strings.ToLower(line), prefix) {
			out = append(out, line)
		}
	}
	return out
}

// SearchFlatpakAll returns every flatpak application ID matching app.
func SearchFlatpakAll(app string) []string {
	var out []string
	for _, line := range skipHeader(runLines("flatpak", "search", "--columns=application", app)) {
		if line == "Application ID" {
			continue
		}
		out = append(out, line)
	}
	return out
}

// VersionFlatpak returns the first search result's version for app, or ""
// if unknown.
func VersionFlatpak(app string) string {
	lines := skipHeader(runLines("flatpak", "search", "--columns=version", app))
	if len(lines) == 0 {
		return ""
	}
	return lines[0]
}

// InfoVersionFlatpak returns the installed version of appID via
// `flatpak info`, or "" if unknown.
func InfoVersionFlatpak(appID string) string {
	out := runOutput("flatpak", "info", appID)
	for _, line := range strings.Split(out, "\n") {
		fields := strings.Fields(line)
		if len(fields) >= 2 && fields[0] == "Version:" {
			return fields[1]
		}
	}
	return ""
}

// InstallFlatpak installs appID via flatpak, running as asUser and
// attaching to the current terminal.
func InstallFlatpak(appID, asUser string) bool {
	return runAttached("sudo", "-u", asUser, "flatpak", "install", appID)
}

// RemoveFlatpak uninstalls appID via flatpak, attaching to the current
// terminal.
func RemoveFlatpak(appID string) bool {
	return runAttached("flatpak", "uninstall", appID)
}

// UpdateFlatpak updates appID via flatpak, attaching to the current
// terminal.
func UpdateFlatpak(appID string) bool {
	return runAttached("flatpak", "update", appID)
}

// SyncAllFlatpak updates every installed flatpak.
func SyncAllFlatpak() bool {
	return runAttached("flatpak", "update")
}
