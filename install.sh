#!/bin/bash

set -e

if [[ $EUID -ne 0 ]]; then
    echo "Please run as root: sudo ./install.sh"
    exit 1
fi

if ! command -v go &>/dev/null; then
    echo ""
    echo "The Go toolchain isn't installed, and spkg needs it to build."
    echo ""
    echo "Here's exactly what this script will run:"
    echo "  1. pacman -S go                 (installs the Go toolchain)"
    echo "  2. Downloads the spkg source, builds it, and installs the"
    echo "     resulting binary to /usr/local/bin/spkg"
    echo ""
    read -rp "Proceed? [y/N]: " choice < /dev/tty
    if [[ "$choice" != "y" && "$choice" != "Y" ]]; then
        echo "Aborted. Install Go yourself (sudo pacman -S go) and re-run this script when ready."
        exit 1
    fi

    echo ""
    echo "==> Installing Go..."
    pacman -S go

    if ! command -v go &>/dev/null; then
        echo "ERROR: go still isn't on PATH after installation. Aborting."
        exit 1
    fi
fi

INSTALL_DIR="/usr/local/bin"
REF="${SPKG_REF:-main}"
SRC_URL="https://codeload.github.com/thesuepster/SuepsPackager/tar.gz/refs/heads/${REF}"

TMP_DIR=$(mktemp -d)

cleanup() {
    rm -rf "$TMP_DIR"
}

trap cleanup EXIT

echo "Downloading spkg source (${REF})..."

curl -fsSL "$SRC_URL" -o "$TMP_DIR/source.tar.gz"

echo "Extracting..."

tar -xzf "$TMP_DIR/source.tar.gz" -C "$TMP_DIR"

SRC_DIR=$(find "$TMP_DIR" -mindepth 1 -maxdepth 1 -type d)

echo "Building spkg..."

(cd "$SRC_DIR" && go build -o "$TMP_DIR/spkg" ./cmd/spkg)

echo "Installing spkg..."

install -m 755 "$TMP_DIR/spkg" "$INSTALL_DIR/spkg"

echo "Done! Run 'spkg doctor' to verify your setup works."
echo "Try: spkg search <app>"
