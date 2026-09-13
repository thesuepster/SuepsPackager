package main

import (
	"fmt"

	"github.com/thesuepster/SuepsPackager/internal/pkgmgr"
)

func cmdDoctor() {
	fmt.Println()
	fmt.Println("==> spkg system check:")
	fmt.Println()

	for _, tool := range []string{"pacman", "paru", "flatpak", "curl", "gpg", "sha256sum"} {
		if pkgmgr.CommandExists(tool) {
			fmt.Printf(" ✓ %-15s available\n", tool)
		} else {
			fmt.Printf(" ✗ %-15s missing\n", tool)
		}
	}
	fmt.Println()
}
