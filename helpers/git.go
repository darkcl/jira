package helpers

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

// GetTicketNumberFromGit - Get current jira ticket number from git branch
func GetTicketNumberFromGit(dir string) string {
	var out bytes.Buffer
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	cmd.Dir = dir
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	branch := out.String()

	fmt.Printf("branch: %s\n", branch)

	ticketNumber := GetTicketNumberFromString(branch)

	fmt.Printf("ticket number: %s\n", ticketNumber)

	return ticketNumber
}

// GetTicketNumberFromString - Get ticket number from string
func GetTicketNumberFromString(input string) string {
	r, _ := regexp.Compile("[A-Z]+-[0-9]+")

	ticketNumber := r.FindString(input)

	return ticketNumber
}
