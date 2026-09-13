package main

import (
	"fmt"
	"os"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdUpdate(args []string) {
	requireSudo("update", args)

	if len(args) == 0 {
		fmt.Println("Usage: sudo spkg update <app>|all [app2] [app3]...")
		os.Exit(1)
	}

	if args[0] == "all" {
		if requireParu() {
			fmt.Println("==> Updating all pacman & AUR packages...")
			pkgmgr.SyncAllAUR()
		} else {
			fmt.Println("==> Updating pacman packages only (paru unavailable)...")
			pkgmgr.SyncAllPacman()
		}
		fmt.Println()
		if requireFlatpak() {
			fmt.Println("==> Updating all Flatpak packages...")
			pkgmgr.SyncAllFlatpak()
		}
		return
	}

	for _, app := range args {
		fmt.Println()
		fmt.Printf("==> Updating '%s'...\n", app)

		switch {
		case pkgmgr.IsInstalledPacman(app):
			if requireParu() {
				pkgmgr.UpdateAUR(app)
			} else {
				pkgmgr.UpdatePacman(app)
			}
		case pkgmgr.IsInstalledFlatpak(app):
			pkgmgr.UpdateFlatpak(app)
		default:
			fmt.Printf(" Package '%s' not found, skipping.\n", app)
		}
	}
}
