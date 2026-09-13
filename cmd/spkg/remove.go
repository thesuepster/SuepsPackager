package main

import (
	"fmt"
	"os"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdRemove(args []string) {
	requireSudo("remove", args)

	if len(args) == 0 {
		fmt.Println("Usage: sudo spkg remove <app> [app2] [app3]...")
		os.Exit(1)
	}

	for _, app := range args {
		fmt.Println()
		fmt.Printf("==> Searching for installed package '%s'...\n", app)

		switch {
		case pkgmgr.IsInstalledPacman(app):
			pkgmgr.RemovePacman(app)
		case pkgmgr.IsInstalledFlatpak(app):
			pkgmgr.RemoveFlatpak(app)
		default:
			fmt.Printf(" Package '%s' not found in pacman or Flatpak, skipping.\n", app)
			fmt.Printf(" If it was installed via AUR, try: paru -R %s\n", app)
		}
	}
}
