// Package pkgmgr wraps the external package management tools (pacman, paru,
// flatpak) that spkg delegates to.
package pkgmgr

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
)

// CommandExists reports whether name is available on PATH.
func CommandExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

// runLines runs name with args and returns its stdout split into
// non-empty, trimmed lines. Errors (including "not found") are treated as
// no output, mirroring the original script's use of "2>/dev/null".
func runLines(name string, args ...string) []string {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()

	var lines []string
	for _, line := range strings.Split(out.String(), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

// runOutput runs name with args and returns raw stdout as a string.
func runOutput(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	_ = cmd.Run()
	return out.String()
}

// runSilent runs name with args, discarding all output, and reports
// whether it exited successfully. Used for cheap "is X installed" checks.
func runSilent(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	return cmd.Run() == nil
}

// runAttached runs name with args, connecting the child to the current
// process's stdin/stdout/stderr so interactive prompts and progress output
// reach the terminal. It reports whether the command exited successfully.
func runAttached(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}

// fieldValue extracts the value of a "Key : value" style line, as produced
// by tools like `pacman -Si`. It looks for the first line whose first
// whitespace-separated field equals key, and returns everything after the
// first colon, trimmed.
func fieldValue(output, key string) string {
	for _, line := range strings.Split(output, "\n") {
		trimmed := strings.TrimSpace(line)
		fields := strings.Fields(trimmed)
		if len(fields) == 0 || fields[0] != key {
			continue
		}
		idx := strings.Index(trimmed, ":")
		if idx == -1 {
			continue
		}
		return strings.TrimSpace(trimmed[idx+1:])
	}
	return ""
}
