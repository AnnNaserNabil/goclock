# Building the Go Clock Package

This document explains how this package was built from scratch, including the steps to create the Arch Linux package.

## Prerequisites

- Go (1.16 or higher)
- Fyne toolkit
- `base-devel` package group (for building Arch Linux packages)
- `git` (for version control)
- `upx` (for binary compression, optional but recommended)

## Building the Application

1. **Install Dependencies**
   ```bash
   # Install Go
   sudo pacman -S go
   
   # Install Fyne
   go get fyne.io/fyne/v2
   
   # Install build dependencies
   sudo pacman -S base-devel upx
   ```

2. **Build the Application**
   ```bash
   # Clone the repository
   git clone https://github.com/AnnNaserNabil/goclock.git
   cd goclock
   
   # Build the binary
   go build -o goclock
   
   # Test the application
   ./goclock
   ```

## Creating the Arch Linux Package

1. **Create the PKGBUILD**
   The `PKGBUILD` file contains all the instructions needed to build and package the application.

2. **Build the Package**
   ```bash
   # In the project root directory
   makepkg -si
   ```
   This will:
   - Download the source code
   - Build the application
   - Create a package file (`.pkg.tar.zst`)
   - Install the package (because of the `-i` flag)

## Package Contents

The package installs the following:
- Binary: `/usr/bin/goclock`
- Desktop file: `/usr/share/applications/goclock.desktop`
- Icon: `/usr/share/pixmaps/goclock.png`

## Verifying the Installation

After installation, you can:
1. Run the application from the terminal:
   ```bash
   goclock
   ```
2. Or find it in your application menu as "Go Clock"

## Updating the Package

When making changes to the application:
1. Update the version number in `PKGBUILD`
2. Update the checksums in `PKGBUILD` using:
   ```bash
   updpkgsums
   ```
3. Build and install the updated package:
   ```bash
   makepkg -si
   ```

## Troubleshooting

- If you get permission errors when running `makepkg`, ensure you're not running as root
- If the application doesn't start, check the logs with:
  ```bash
  journalctl -xe
  ```
- For build issues, ensure all dependencies are installed and your Go environment is properly set up
