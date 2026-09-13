// Command spkg is a package manager wrapper for Arch-based Linux systems.
// It searches and installs packages from Pacman, the AUR (via Paru), and
// Flatpak from a single command.
package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println()
	fmt.Println("Usage: spkg <command> [options]")
	fmt.Println()
	fmt.Println(" Commands:")
	fmt.Println("   install <app> [app2]...   Install one or more packages")
	fmt.Println("   remove  <app> [app2]...   Remove one or more packages")
	fmt.Println("   update  <app>|all         Update a package or everything")
	fmt.Println("   search  <app>             Search for a package")
	fmt.Println("   info    <app>             Show package details")
	fmt.Println("   doctor                    Check system dependencies")
	fmt.Println()
	fmt.Println(" Flags:")
	fmt.Println("   --pacman   Only search/install from pacman")
	fmt.Println("   --aur      Only search/install from AUR")
	fmt.Println("   --flatpak  Only search/install from Flatpak")
	fmt.Println()
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	action := os.Args[1]
	rest := os.Args[2:]

	switch action {
	case "install":
		cmdInstall(rest)
	case "remove":
		cmdRemove(rest)
	case "update":
		cmdUpdate(rest)
	case "search":
		cmdSearch(rest)
	case "info":
		cmdInfo(rest)
	case "doctor":
		cmdDoctor()
	default:
		usage()
		os.Exit(1)
	}
}
