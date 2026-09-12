package main

import (
	"fmt"
	"os"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdSearch(args []string) {
	filter, rest := repoFilter(args)
	if len(rest) == 0 {
		fmt.Println("Usage: spkg search <app> [--pacman|--aur|--flatpak]")
		os.Exit(1)
	}
	app := rest[0]

	fmt.Println()
	fmt.Printf("==> Exact results for '%s':\n", app)
	fmt.Println()

	foundExact := false
	if wantsPacman(filter) {
		for _, pkg := range pkgmgr.SearchPacmanExact(app) {
			printResult(pkg, pkgmgr.VersionPacman(pkg), pkgmgr.Pacman)
			foundExact = true
		}
	}
	if wantsAUR(filter) && requireParu() {
		for _, pkg := range pkgmgr.SearchAURExact(app) {
			printResult(pkg, pkgmgr.VersionAUR(pkg), pkgmgr.AUR)
			foundExact = true
		}
	}
	if wantsFlatpak(filter) && requireFlatpak() {
		for _, pkg := range pkgmgr.SearchFlatpakExact(app) {
			printResult(pkg, pkgmgr.VersionFlatpak(pkg), pkgmgr.Flatpak)
			foundExact = true
		}
	}
	if !foundExact {
		fmt.Println(" No exact results found.")
	}

	fmt.Println()
	fmt.Println("==> Closely related results:")
	fmt.Println()

	foundRelated := false
	if wantsPacman(filter) {
		for _, pkg := range pkgmgr.SearchPacmanPrefix(app) {
			if pkg == app {
				continue
			}
			printResult(pkg, pkgmgr.VersionPacman(pkg), pkgmgr.Pacman)
			foundRelated = true
		}
	}
	if wantsAUR(filter) && requireParu() {
		for _, pkg := range pkgmgr.SearchAURPrefix(app) {
			if pkg == app {
				continue
			}
			printResult(pkg, pkgmgr.VersionAUR(pkg), pkgmgr.AUR)
			foundRelated = true
		}
	}
	if wantsFlatpak(filter) && requireFlatpak() {
		for _, pkg := range pkgmgr.SearchFlatpakPrefix(app) {
			if pkg == app {
				continue
			}
			printResult(pkg, pkgmgr.VersionFlatpak(pkg), pkgmgr.Flatpak)
			foundRelated = true
		}
	}
	if !foundRelated {
		fmt.Println(" No closely related results found.")
	}

	fmt.Println()
	choice := prompt("==> Show all results? This may return many results and slow down search. [y/n]: ")
	if choice != "y" {
		return
	}

	fmt.Println()
	fmt.Printf("==> All results for '%s':\n", app)
	fmt.Println()

	if wantsPacman(filter) {
		for _, pkg := range pkgmgr.SearchPacmanAll(app) {
			if pkg == app {
				continue
			}
			printResult(pkg, pkgmgr.VersionPacman(pkg), pkgmgr.Pacman)
		}
	}
	if wantsAUR(filter) && requireParu() {
		for _, pkg := range pkgmgr.SearchAURAll(app) {
			if pkg == app {
				continue
			}
			printResult(pkg, pkgmgr.VersionAUR(pkg), pkgmgr.AUR)
		}
	}
	if wantsFlatpak(filter) && requireFlatpak() {
		for _, pkg := range pkgmgr.SearchFlatpakAll(app) {
			printResult(pkg, pkgmgr.VersionFlatpak(pkg), pkgmgr.Flatpak)
		}
	}
}
