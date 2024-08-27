# refresh ✨

A smol binary that installs the tools needed to get started with programming with HTML, CSS, JavaScript, and PHP. A handful of important VS Code extensions are installed as well.

> This project is set to target macOS machines.

## What happens?

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
4. Updates the VS Code user settings profile `settings.json`.
5. Installs GitHub Desktop (via Homebrew)
6. Installs PHP (via Homebrew)

## Installation

From the command line:

```sh
curl -L -o refresh https://github.com/villaleo/refresh/releases/download/v1.1.1/refresh_universal
chmod +x ./refresh_universal
./refresh
```

Or, checkout the releases [here](https://github.com/villaleo/refresh/releases).

There is a universal binary `refresh_universal` for your convenience. You may also the `refresh_amd64` for Intel Macs and `refresh_arm64` for Silicon Macs.
