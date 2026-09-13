package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdInstall(args []string) {
	requireSudo("install", args)
	filter, rest := repoFilter(args)
	if len(rest) == 0 {
		fmt.Println("Usage: sudo spkg install <app> [app2]... [--pacman|--aur|--flatpak]")
		os.Exit(1)
	}
	sudoUser := requireSudoUser()

	for _, app := range rest {
		fmt.Println()
		fmt.Printf("==> Searching for '%s'...\n", app)

		results := searchExact(app, filter)
		if len(results) == 0 {
			results = searchAll(app, filter)
		}

		switch total := len(results); {
		case total == 0:
			fmt.Printf("No results found for '%s', skipping.\n", app)

		case total == 1:
			fmt.Printf("Found: %s [%s]\n", results[0].Name, results[0].Source)
			fmt.Println("Installing...")
			doInstall(results[0], sudoUser)

		default:
			fmt.Println("Multiple results found:")
			fmt.Println()
			for i, r := range results {
				fmt.Printf(" [%d] %-35s [%s]\n", i+1, r.Name, r.Source)
			}
			choice := prompt("Enter number to install (or 0 to skip): ")
			if !isDigitsOnly(choice) {
				fmt.Printf("Invalid input, skipping '%s'.\n", app)
				continue
			}
			idx, _ := strconv.Atoi(choice)
			if idx == 0 {
				fmt.Printf("Skipping '%s'.\n", app)
				continue
			}
			if idx > total {
				fmt.Printf("Invalid selection, skipping '%s'.\n", app)
				continue
			}
			doInstall(results[idx-1], sudoUser)
		}
	}
}

func searchExact(app, filter string) []pkgmgr.Result {
	var results []pkgmgr.Result
	if wantsPacman(filter) {
		for _, pkg := range pkgmgr.SearchPacmanExact(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.Pacman})
		}
	}
	if wantsAUR(filter) && requireParu() {
		for _, pkg := range pkgmgr.SearchAURExact(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.AUR})
		}
	}
	if wantsFlatpak(filter) && requireFlatpak() {
		for _, pkg := range pkgmgr.SearchFlatpakExact(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.Flatpak})
		}
	}
	return results
}

func searchAll(app, filter string) []pkgmgr.Result {
	var results []pkgmgr.Result
	if wantsPacman(filter) {
		for _, pkg := range pkgmgr.SearchPacmanAll(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.Pacman})
		}
	}
	if wantsAUR(filter) && requireParu() {
		for _, pkg := range pkgmgr.SearchAURAll(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.AUR})
		}
	}
	if wantsFlatpak(filter) && requireFlatpak() {
		for _, pkg := range pkgmgr.SearchFlatpakAll(app) {
			results = append(results, pkgmgr.Result{Name: pkg, Source: pkgmgr.Flatpak})
		}
	}
	return results
}

func doInstall(r pkgmgr.Result, sudoUser string) {
	switch r.Source {
	case pkgmgr.Pacman:
		pkgmgr.InstallPacman(r.Name)
	case pkgmgr.AUR:
		pkgmgr.InstallAUR(r.Name, sudoUser)
	case pkgmgr.Flatpak:
		pkgmgr.InstallFlatpak(r.Name, sudoUser)
	}
}

func isDigitsOnly(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
