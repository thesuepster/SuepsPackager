package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

var stdinReader = bufio.NewReader(os.Stdin)

// requireSudo exits unless the process is running with root privileges,
// telling the user how to invoke the current subcommand under sudo.
func requireSudo(action string, args []string) {
	if os.Geteuid() != 0 {
		fmt.Printf("This command requires sudo. Try: sudo spkg %s %s\n", action, strings.Join(args, " "))
		os.Exit(1)
	}
}

// requireSudoUser exits unless spkg was invoked via `sudo` (rather than
// from an already-root shell), returning the original unprivileged user.
func requireSudoUser() string {
	user := os.Getenv("SUDO_USER")
	if user == "" {
		fmt.Println("Please run using sudo, not a root shell.")
		os.Exit(1)
	}
	return user
}

// requireParu reports whether paru is available, printing a notice if not.
func requireParu() bool {
	if !pkgmgr.HasParu() {
		fmt.Println(" AUR support unavailable: paru not installed.")
		return false
	}
	return true
}

// requireFlatpak reports whether flatpak is available, printing a notice
// if not.
func requireFlatpak() bool {
	if !pkgmgr.HasFlatpak() {
		fmt.Println(" Flatpak support unavailable: flatpak not installed.")
		return false
	}
	return true
}

// repoFilter pulls a --pacman/--aur/--flatpak flag out of args, returning
// the filter (empty string if none given) and the remaining arguments.
func repoFilter(args []string) (filter string, rest []string) {
	for _, arg := range args {
		switch arg {
		case "--pacman":
			filter = "pacman"
		case "--aur":
			filter = "aur"
		case "--flatpak":
			filter = "flatpak"
		default:
			rest = append(rest, arg)
		}
	}
	return filter, rest
}

func wantsPacman(filter string) bool  { return filter == "" || filter == "pacman" }
func wantsAUR(filter string) bool     { return filter == "" || filter == "aur" }
func wantsFlatpak(filter string) bool { return filter == "" || filter == "flatpak" }

// printResult renders one search/install candidate line, marking it
// installed if applicable.
func printResult(name, version string, source pkgmgr.Source) {
	installed := pkgmgr.IsInstalledPacman(name) || pkgmgr.IsInstalledFlatpak(name)
	mark := " "
	if installed {
		mark = "✓"
	}
	fmt.Printf(" [%s]  %-35s %-20s [%s]\n", mark, name, "v"+version, source)
}

// prompt asks the user a question and returns their trimmed response.
func prompt(question string) string {
	fmt.Print(question)
	line, _ := stdinReader.ReadString('\n')
	return strings.TrimSpace(line)
}
