---
sidebar_position: 3
---

# Installation Guide

GASP can be installed on any major operating system. Choose your platform below for specific installation instructions.

## Linux Installation

```bash
curl -L "https://github.com/bytewer-lab/gasp/releases/latest/download/gasp-linux" -o gasp
chmod +x gasp
sudo mv gasp /usr/local/bin/
```

## macOS Installation

```bash
curl -L "https://github.com/bytewer-lab/gasp/releases/latest/download/gasp-darwin" -o gasp
chmod +x gasp
sudo mv gasp /usr/local/bin/
```

## Windows Installation

1. Download `gasp-windows.exe` from the [latest release](https://github.com/bytewer-lab/gasp/releases/latest)
2. Rename it to `gasp.exe`
3. Move it to a directory in your PATH:
   - Option 1: Move to `C:\Windows\System32\`
   - Option 2: Create a new directory, add it to PATH, and move the executable there

## Verifying Installation

After installation, verify GASP is working correctly:

```bash
gasp --version
```

You should see the current version of GASP displayed.