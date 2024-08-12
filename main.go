package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Verify Homebrew is installed.
	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	mustInstall("brew", cmd)

	// Verify VS Code is installed.
	cmd = exec.Command("brew", "install", "--cask", "visual-studio-code")
	mustInstall("code", cmd)

	// VS Code extensions to install
	vscodeExtensions := []string{
		"beardedbear.beardedtheme",
		"bmewburn.vscode-intelephense-client",
		"davidanson.vscode-markdownlint",
		"ecmel.vscode-html-css",
		"esbenp.prettier-vscode",
		"neilbrayfield.php-docblocker",
		"pkief.material-icon-theme",
		"ritwickdey.liveserver",
		"wayou.vscode-todo-highlight",
		"xabikos.javascriptsnippets",
		"yzhang.markdown-all-in-one",
	}

	// Install VS Code extensions.
	fmt.Println("checking extensions..")
	for _, ext := range vscodeExtensions {
		fmt.Printf("\t%s.. ", ext)
		cmd := exec.Command("code", "--install-extension", ext)
		if err := cmd.Run(); err != nil {
			fmt.Printf("\nerror: failed to install extension %q.", ext)
			fmt.Println("aborted.")
			os.Exit(1)
		}
		fmt.Println("ok")
	}
	fmt.Println("done. all extensions installed.")

	// Verify GitHub Desktop is installed.
	cmd = exec.Command("brew", "install", "--cask", "github")
	mustInstall("github", cmd)

	// Verify PHP is installed.
	cmd = exec.Command("brew", "install", "php")
	mustInstall("php", cmd)

	fmt.Println("finished. all processes done.")
}

// commandExists checks if the command name exists on the computer.
func commandExists(name string) bool {
	cmd := exec.Command("command", "-v", name, "&> /dev/null")
	return cmd.Run() == nil
}

// mustInstall checks if command s exists. If s exists, then a message is
// printed to stdout. Otherwise cmd is executed, the output is printed to
// stdout, and the process is killed.
func mustInstall(s string, cmd *exec.Cmd) {
	if commandExists(s) {
		fmt.Printf("%q already installed.\n", s)
		return
	}

	fmt.Printf("%q not installed.\n", s)
	fmt.Printf("installing %q.. ", s)
	out, err := cmd.Output()
	fmt.Printf("%s\n", out)
	if err != nil {
		fmt.Printf("\nerror: failed to install %q.\n", s)
		fmt.Println("aborted.")
		os.Exit(1)
	}
	fmt.Println("done.")
}
