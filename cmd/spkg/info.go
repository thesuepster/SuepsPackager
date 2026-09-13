package main

import (
	"fmt"
	"os"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdInfo(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: spkg info <app>")
		os.Exit(1)
	}
	app := args[0]
	fmt.Println()

	switch {
	case pkgmgr.IsInstalledPacman(app):
		name, version, repo, desc, _ := pkgmgr.InfoPacman(app)
		fmt.Printf(" Name:        %s\n", name)
		fmt.Printf(" Version:     %s\n", version)
		fmt.Printf(" Repository:  %s\n", repo)
		fmt.Printf(" Description: %s\n", desc)
		fmt.Println(" Installed:   Yes")

	case pkgmgr.IsInstalledFlatpak(app):
		version := pkgmgr.InfoVersionFlatpak(app)
		fmt.Printf(" Name:        %s\n", app)
		fmt.Printf(" Version:     %s\n", version)
		fmt.Println(" Repository:  flatpak")
		fmt.Println(" Installed:   Yes")

	default:
		if name, version, repo, desc, ok := pkgmgr.InfoPacman(app); ok {
			fmt.Printf(" Name:        %s\n", name)
			fmt.Printf(" Version:     %s\n", version)
			fmt.Printf(" Repository:  %s\n", repo)
			fmt.Printf(" Description: %s\n", desc)
			fmt.Println(" Installed:   No")
		} else if requireParu() {
			if name, version, desc, ok := pkgmgr.InfoAUR(app); ok {
				fmt.Printf(" Name:        %s\n", name)
				fmt.Printf(" Version:     %s\n", version)
				fmt.Println(" Repository:  aur")
				fmt.Printf(" Description: %s\n", desc)
				fmt.Println(" Installed:   No")
			} else {
				fmt.Printf(" No info found for '%s'.\n", app)
			}
		}
	}
	fmt.Println()
}
