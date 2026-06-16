##READ BEFORE INSTALLING:
IF YOU WOULD LIKE TO TRY THIS PACKAGER OUT DO NOT DOWNLOAD THE ZIP FILE PROVIDED BY THE RELEASES
ENTER THIS COMMAND ON YOUR ARCH BASED SYSTEM AS IT WILL INSTANTLY DOWNLOAD AND MOVE THE FILE TO WHERE IT NEEDS TO BE

curl -fsSL https://raw.githubusercontent.com/thesuepster/SuepsPackager/main/install.sh | sudo bash


#The SPKG installer is a package installer that can install any app or tool as long as it's in the Pacman, AUR, and Flatpak repositories.
This only works on Arch-based linux and must have Flatpak and Paru installed on system for the full experience.
You can do use SPKG to install Paru and Flatpak if you don't already have it.
Here are the command layouts, insert whatever's necessary: sudo spkg <spkg feature> <app>
I will give example commands for each feature using the Brave Browser as an example.

Dependencies:
Pacman - comes with Arch/Arch based distro's
Paru - needed to access the AUR
Flatpak - needed for accessing Flatpaks

Features:

1. Install, installs whatever app you enter off of three different repositories. BE SPECIFIC! Command: sudo spkg install brave-bin
2. Search, searches Pacman, AUR, and Flatpak in that order for whatever you searched up. Doesn't have to be insanely specific. Command: sudo spkg search Brave
3. Update, finds whatever packages needs updating and updates them. Command: sudo spkg update brave-bin OR sudo spkg update all
4. Remove, removes packages. BE SPECIFIC! Command: sudo spkg remove brave-bin

Feedback is greatly appreciated.
