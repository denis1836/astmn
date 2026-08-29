package packer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"astmn/internal/manifest"
)

var newPackage manifest.Manifest

func CreateNewPackage() error {
	r := bufio.NewScanner(os.Stdin)
	if !r.Scan() {
		return fmt.Errorf("failed to create a bufio stdin reader")
	}

	fmt.Printf("Package name: ")
	newPackage.Name = r.Text()

	fmt.Printf("%s version: ", newPackage.Name)
	newPackage.Version = r.Text()

	newPackage.Date = time.Now().Format("2006-01-02")

	fmt.Printf("%s author: ", newPackage.Name)
	newPackage.Author = r.Text()

	fmt.Printf("%s description: ", newPackage.Name)
	newPackage.Description = r.Text()

	fmt.Printf("%s Contributors: ", newPackage.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		newPackage.Contributors = append(newPackage.Contributors, line)
	}

	//TODO: file generation and creation
	fmt.Printf("%s system file name: ", newPackage.Name)
	newPackage.FileName = r.Text()

	fmt.Printf("%s download URL: ", newPackage.Name)
	newPackage.DownloadURL = r.Text()

	//TODO: automatic hash generation
	fmt.Printf("%s sha256: ", newPackage.Name)
	newPackage.SHA256 = r.Text()

	fmt.Printf("%s install path: ", newPackage.Name)
	newPackage.InstallPath = r.Text()

	fmt.Printf("%s dependencies: ", newPackage.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		newPackage.DependsOn = append(newPackage.DependsOn, line)
	}

	fmt.Printf("%s file contents: ", newPackage.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		newPackage.Contents = append(newPackage.Contents, line)
	}

	fmt.Printf("Changelog: ")
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		newPackage.Changelog = append(newPackage.Changelog, line)
	}

	return nil
}
