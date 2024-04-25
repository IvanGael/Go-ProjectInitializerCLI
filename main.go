package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// go get -u github.com/AlecAivazis/survey/v2

func main() {
	var technology string

	// List of technologies
	technologies := []string{"React", "Angular", "Vue.js", "SvelteKit", "Next.js", "Flutter", "React Native", "Astro", "Ionic", "Node.js", "Flask", "Django"}

	// Prompt the user to select a technology
	prompt := &survey.Select{
		Message: "Choose the technology you want to initialize the project in:",
		Options: technologies,
	}
	survey.AskOne(prompt, &technology)

	// Initialize the project based on the selected technology
	initializeProject(technology)
}

func initializeProject(technology string) {
	fmt.Printf("Initializing %s project...\n", technology)

	var projectName string
	// Prompt the user to enter the project name
	prompt := &survey.Input{
		Message: "Enter the project name:",
	}
	survey.AskOne(prompt, &projectName)

	var path string
	// Prompt the user to enter the path
	prompt = &survey.Input{
		Message: "Enter the path to initialize the project (press Enter for current directory):",
	}
	survey.AskOne(prompt, &path)

	path = strings.TrimSpace(path)
	if path == "" {
		path = "."
	}

	// Create directory if it doesn't exist
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	switch technology {
	case "React":
		cmd := exec.Command("npx", "create-react-app", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Angular":
		cmd := exec.Command("ng", "new", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Vue.js":
		cmd := exec.Command("vue", "create", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "SvelteKit":
		cmd := exec.Command("npx", "create-svelte", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Next.js":
		cmd := exec.Command("npx", "create-next-app", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Flutter":
		cmd := exec.Command("flutter", "create", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "React Native":
		cmd := exec.Command("npx", "react-native", "init", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Astro":
		cmd := exec.Command("npx", "create-astro", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	case "Ionic":
		cmd := exec.Command("ionic", "start", projectName, "--type", "angular")
		cmd.Dir = path
		executeCommand(cmd)
	case "Node.js":
		cmd := exec.Command("npm", "init", "-y")
		cmd.Dir = path
		executeCommand(cmd)
	case "Flask":
		cmd := exec.Command("python", "-m", "flask", "run")
		cmd.Dir = path
		executeCommand(cmd)
	case "Django":
		cmd := exec.Command("django-admin", "startproject", projectName)
		cmd.Dir = path
		executeCommand(cmd)
	}

	fmt.Printf("Project %s initialized successfully in %s\n", projectName, path)
}

func executeCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
