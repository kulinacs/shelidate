#!/bin/bash
# https://gitlab.winehq.org/wine/wine/-/wikis/Debian-Ubuntu

# add 32-bit architecture
dpkg --add-architecture i386

# download and add the repository key:
mkdir -pm755 /etc/apt/keyrings
wget -O - https://dl.winehq.org/wine-builds/winehq.key | gpg --dearmor -o /etc/apt/keyrings/winehq-archive.key -

# add the bookworm sources
# this is hardcoded and could be made configurable or dynamically discoverable in the future
wget -NP /etc/apt/sources.list.d/ https://dl.winehq.org/wine-builds/debian/dists/bookworm/winehq-bookworm.sources

# update to include the new sources
apt-get update

# install wine
apt-get install -y --install-recommends winehq-stable

# install mingw
apt-get install -y gcc-mingw-w64-x86-64 gcc-mingw-w64-i686

# clean sources to shrink the layer
rm -rf /var/cache/apt/archives /var/lib/apt/lists/*

