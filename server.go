package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// Download dependencies using "go mod download" command
	cmd := exec.Command("go", "mod", "download")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	// Start the server
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler is the exported function that Vercel will invoke
func Handler(w http.ResponseWriter, r *http.Request) {
	// Run your main.go file using "go run main.go" command
	cmd := exec.Command("go", "run", "main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(w, "Go project is running on Vercel!")
}
