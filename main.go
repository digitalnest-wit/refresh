package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	config := map[string]interface{}{
		"editor.tabSize":             4,
		"workbench.colorTheme":       "Bearded Theme Coffee",
		"workbench.iconTheme":        "material-icon-theme",
		"workbench.productIconTheme": "icons-carbon",
		"[css]": map[string]interface{}{
			"editor.defaultFormatter": "esbenp.prettier-vscode",
			"editor.formatOnPaste":    true,
			"editor.formatOnSave":     true,
		},
		"[html]": map[string]interface{}{
			"editor.defaultFormatter": "esbenp.prettier-vscode",
			"editor.formatOnPaste":    true,
			"editor.formatOnSave":     true,
			"editor.tabSize":          2,
		},
		"liveServer.settings.donotShowInfoMsg": true,
		"explorer.confirmDelete":               false,
		"explorer.confirmDragAndDrop":          false,
	}

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
