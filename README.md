# refresh ✨

A smol binary that installs the tools needed to get started with programming with HTML, CSS, JavaScript, and PHP. A handful of important VS Code extensions are installed as well.

> The binaries on this project are set to target macOS machines.

## What's included

The binary does the following.

1. Installs Homebrew (required to install everything below 👇🏻).
2. Installs VS Code (via Homebrew)
3. Installs the following VS Code extensions:
   - `beardedbear.beardedtheme` (color theme)
   - `bmewburn.vscode-intelephense-client` (php intellisense)
   - `davidanson.vscode-markdownlint` (markdown linting)
   - `ecmel.vscode-html-css` (html and css support)
   - `esbenp.prettier-vscode` (code formatting)
   - `neilbrayfield.php-docblocker` (php documentation)
   - `pkief.material-icon-theme` (icon theme)
   - `ritwickdey.liveserver` (hot reloading)
   - `wayou.vscode-todo-highlight` (highlight keywords)
   - `xabikos.javascriptsnippets` (snippets)
   - `yzhang.markdown-all-in-one` (more markdown support)
4. Installs GitHub Desktop (via Homebrew)
5. Installs PHP (via Homebrew)

## Installation

```sh
curl -L -o refresh https://github.com/villaleo/refresh/releases/download/1.0.0/refresh
chmod +x ./refresh
./refresh
```

Checkout the releases [here](https://github.com/villaleo/refresh/releases). Or click on the latest release on the right side bar 👉🏻.

 All the binaries can be found in the `bin/` directory.

There is a universal binary `refresh` for your convenience. You may also the `refresh_amd64` for Intel Macs and `refresh_arm64` for Silicon Macs.
