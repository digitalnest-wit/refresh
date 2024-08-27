package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

type jsonMap map[string]interface{}

var (
	// list of vscode extensions to install
	vscodeExtensions = []string{
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
	// Updated vscode user config to use
	config = jsonMap{
		"editor.tabSize":             4,
		"workbench.colorTheme":       "Bearded Theme Coffee",
		"workbench.iconTheme":        "material-icon-theme",
		"workbench.productIconTheme": "icons-carbon",
		"[css]": jsonMap{
			"editor.defaultFormatter": "esbenp.prettier-vscode",
			"editor.formatOnPaste":    true,
			"editor.formatOnSave":     true,
		},
		"[html]": jsonMap{
			"editor.defaultFormatter": "esbenp.prettier-vscode",
			"editor.formatOnPaste":    true,
			"editor.formatOnSave":     true,
			"editor.tabSize":          2,
		},
		"liveServer.settings.donotShowInfoMsg": true,
		"explorer.confirmDelete":               false,
		"explorer.confirmDragAndDrop":          false,
	}
)

func main() {
	// Install homebrew
	cmd := exec.Command("/bin/bash", "-c", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	mustInstall("brew", cmd)

	// Install vscode
	cmd = exec.Command("brew", "install", "--cask", "visual-studio-code")
	mustInstall("code", cmd)

	// Install each vscode extenion in the list
	log.Println("installing extensions")
	for _, ext := range vscodeExtensions {
		log.Printf("    %s\n", ext)
		cmd := exec.Command("code", "--install-extension", ext)
		if err := cmd.Run(); err != nil {
			log.Printf("failed to install extension %q.\n", ext)
			log.Println(err)
			os.Exit(1)
		}
	}
	log.Println("done. all extensions installed.")

	// Update vscode user config
	mustUpdateVscodeConfig(config)
	log.Println("vscode user settings updated.")

	// Install github
	cmd = exec.Command("brew", "install", "--cask", "github")
	mustInstall("github", cmd)

	// Install php
	cmd = exec.Command("brew", "install", "php")
	mustInstall("php", cmd)

	log.Println("all done!")
}

func mustUpdateVscodeConfig(config map[string]interface{}) {
	// Get the current logged-in user
	cmd := exec.Command("whoami")
	out, err := cmd.Output()
	if err != nil {
		log.Println("failed to get logged in user.")
		os.Exit(1)
	}
	user := strings.Trim(string(out), "\n\r")

	// Open the vscode user settings file
	settingsPath := "/Users/" + user + "/Library/Application Support/Code/User/settings.json"
	usf, err := os.Open(settingsPath)
	if err != nil {
		log.Println("failed to locate vscode user settings file.")
		os.Exit(1)
	}
	defer usf.Close()

	// Create a temporary file to store changes
	tf, err := os.Create("tmp.uset.json")
	if err != nil {
		log.Println("failed to create a temporary file.")
		os.Exit(1)
	}
	defer func() {
		tf.Close()
		os.Remove(tf.Name()) // Clean up temp file
	}()

	// Copy the contents of the existing settings file into the temporary file
	_, err = io.Copy(tf, usf)
	if err != nil {
		log.Println("failed to copy existing settings.")
		os.Exit(1)
	}

	// Close and reopen temp file for reading and writing
	tf.Close()
	tf, err = os.OpenFile(tf.Name(), os.O_RDWR, 0644)
	if err != nil {
		log.Println("failed to reopen temporary file for updates.")
		os.Exit(1)
	}

	// Read the existing user settings
	settings := make(map[string]any)
	err = json.NewDecoder(tf).Decode(&settings)
	if err != nil {
		log.Println("failed to parse user settings file.", err)
		os.Exit(1)
	}

	// Merge the new config into the existing settings
	for key, value := range config {
		settings[key] = value
	}

	// Move file pointer back to the start of the file
	tf.Seek(0, 0)

	// Write the updated settings back to the temp file
	tf.Truncate(0) // Clear the file before writing new content
	err = json.NewEncoder(tf).Encode(settings)
	if err != nil {
		log.Println("failed to write updated settings.", err)
		os.Exit(1)
	}

	// Replace the original file with the updated temporary file
	err = os.Rename(tf.Name(), settingsPath)
	if err != nil {
		log.Println("failed to save changes to the vscode user settings file.", err)
		os.Exit(1)
	}
}

func commandExists(name string) bool {
	cmd := exec.Command("command", "-v", name, "&> /dev/null")
	return cmd.Run() == nil
}

func mustInstall(s string, cmd *exec.Cmd) {
	if commandExists(s) {
		log.Printf("%q already installed.\n", s)
		return
	}

	log.Printf("installing %q.. ", s)
	_, err := cmd.Output()
	if err != nil {
		log.Printf("\nfailed to install %q.\n", s)
		os.Exit(1)
	}
	log.Println("done.")
}
