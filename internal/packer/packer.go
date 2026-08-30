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

	fmt.Printf("Package name: ")
	np.Name = r.Text()

	fmt.Printf("%s version: ", np.Name)
	np.Version = r.Text()

	np.Date = time.Now().Format("2006-01-02")

	fmt.Printf("%s author: ", np.Name)
	np.Author = r.Text()

	fmt.Printf("%s description: ", np.Name)
	np.Description = r.Text()

	fmt.Printf("%s Contributors: ", np.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		np.Contributors = append(np.Contributors, line)
	}

	//TODO: file generation and creation
	fmt.Printf("%s system file name: ", np.Name)
	np.FileName = r.Text()

	fmt.Printf("%s download URL: ", np.Name)
	np.DownloadURL = r.Text()

	//TODO: automatic hash generation
	fmt.Printf("%s sha256: ", np.Name)
	np.SHA256 = r.Text()

	fmt.Printf("%s install path: ", np.Name)
	np.InstallPath = r.Text()

	fmt.Printf("%s dependencies: ", np.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		np.DependsOn = append(np.DependsOn, line)
	}

	fmt.Printf("%s file contents: ", np.Name)
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		np.Contents = append(np.Contents, line)
	}

	fmt.Printf("Changelog: ")
	for {
		fmt.Print("> ")
		line := r.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		np.Changelog = append(np.Changelog, line)
	}

	return nil
}
