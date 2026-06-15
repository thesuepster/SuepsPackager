#!/bin/bash

if [[ $EUID -ne 0 ]]; then
    echo "Please run as root: sudo ./install.sh"
    exit 1
fi

echo "Installing spkg..."
curl -fsSL https://raw.githubusercontent.com/thesuepster/SuepsPackager/main/spkg -o /usr/local/bin/spkg
chmod +x /usr/local/bin/spkg
echo "Done! Try: sudo spkg search <app>"
