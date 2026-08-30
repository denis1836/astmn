package packer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"astmn/internal/manifest"
)

func CreateNewPackage(np *manifest.Manifest) error {
	r := bufio.NewScanner(os.Stdin)
	if !r.Scan() {
		return fmt.Errorf("failed to create a bufio stdin reader")
	}

	np.Name = readLine(r, "Package name: ")
	np.Version = readLine(r, fmt.Sprintf("%s version: ", np.Name))
	np.Date = time.Now().Format("2006-01-02")
	np.Author = readLine(r, fmt.Sprintf("%s author: ", np.Name))
	np.Description = readLine(r, fmt.Sprintf("%s description: ", np.Name))

	fmt.Printf("%s Contributors: ", np.Name)
	for {
		line := readLine(r, "> ")
		if line == "" {
			break
		}
		np.Contributors = append(np.Contributors, line)
	}

	//TODO: file generation and creation
	np.FileName = readLine(r, fmt.Sprintf("%s system file name: ", np.Name))
	np.DownloadURL = readLine(r, fmt.Sprintf("%s download URL: ", np.Name))
	//TODO: automatic hash generation
	np.SHA256 = readLine(r, fmt.Sprintf("%s sha256: ", np.Name))
	np.InstallPath = readLine(r, fmt.Sprintf("%s install path: ", np.Name))

	fmt.Printf("%s dependencies: ", np.Name)
	for {
		line := readLine(r, "> ")
		if line == "" {
			break
		}
		np.DependsOn = append(np.DependsOn, line)
	}

	fmt.Printf("%s file contents: ", np.Name)
	for {
		line := readLine(r, "> ")
		if line == "" {
			break
		}
		np.Contents = append(np.Contents, line)
	}

	fmt.Printf("Changelog: ")
	for {
		line := readLine(r, "> ")
		if line == "" {
			break
		}
		np.Changelog = append(np.Changelog, line)
	}

	return nil
}

func readLine(r *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	if !r.Scan() {
		return ""
	}
	return strings.TrimSpace(r.Text())
}
