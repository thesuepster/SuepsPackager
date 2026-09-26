# Suep's Packager (spkg)

A package manager wrapper for Arch-based Linux systems, written in Go. spkg searches and installs packages from Pacman, the AUR (via Paru), and Flatpak — all from one command.

---

## Installation

Run this in a terminal to install spkg on your Arch-based system:

    curl -fsSL https://raw.githubusercontent.com/thesuepster/SuepsPackager/main/install.sh | sudo bash

The installer downloads the spkg source, builds it with `go build`, and installs the resulting binary to `/usr/local/bin/`.

**Dependencies:** Building requires the Go toolchain. Make sure you also have `paru` and `flatpak` installed for full functionality:

    sudo pacman -S go flatpak

`paru` itself isn't in the official Arch repos, so how you get it depends on what you're running:

- **Base Arch:** paru is AUR-only, so bootstrap it manually:

      sudo pacman -S --needed base-devel git
      git clone https://aur.archlinux.org/paru.git
      cd paru
      makepkg -si

- **Some Arch-based distros (e.g. CachyOS):** paru ships in the distro's own main repo, so a plain `sudo pacman -S paru` is enough.

Run `spkg doctor` after installing to check your setup.

### Building from source manually

    git clone https://github.com/thesuepster/SuepsPackager.git
    cd SuepsPackager
    go build -o spkg ./cmd/spkg
    sudo install -m 755 spkg /usr/local/bin/spkg

---

## Commands

| Command | Description |
|---|---|
| `sudo spkg install <app>` | Install one or more packages |
| `sudo spkg remove <app>` | Remove one or more packages |
| `sudo spkg update <app>\|all` | Update a specific package or everything |
| `spkg search <app>` | Search for a package |
| `spkg info <app>` | Show package details |
| `spkg doctor` | Check system dependencies |

---

## Flags

Add these to `install` or `search` to target a specific repository:

| Flag | Description |
|---|---|
| `--pacman` | Only search/install from Pacman |
| `--aur` | Only search/install from the AUR |
| `--flatpak` | Only search/install from Flatpak |

---

## Examples

```bash
# Search for brave
spkg search brave

# Search AUR only
spkg search brave --aur

# Install a single package
sudo spkg install brave-bin

# Install multiple packages at once
sudo spkg install brave-bin kitty fish

# Install from a specific repo
sudo spkg install brave-bin --aur

# Update everything
sudo spkg update all

# Remove multiple packages
sudo spkg remove brave-bin kitty

# Get info on a package
spkg info firefox

# Check your setup
spkg doctor
```

---

## Notes

- `search`, `info`, and `doctor` do not require sudo
- `install`, `remove`, and `update` require sudo
- spkg must be run with `sudo`, not from a root shell directly
- Building from source means you're compiling exactly what's in this repo — read it before you build it
- GPG signing coming in a future release

---

## License

MIT License — see [LICENSE](LICENSE) for details.
