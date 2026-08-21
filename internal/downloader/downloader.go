package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"astmn/internal/log"
)

var (
	gdrivePathRegex  = regexp.MustCompile(`/file/d/([a-zA-Z0-9_-]+)`)
	gdriveQueryRegex = regexp.MustCompile(`[?&]id=([a-zA-Z0-9_-]+)`)
)

func ExtractGDriveID(rawURL string) string {
	if matches := gdrivePathRegex.FindStringSubmatch(rawURL); len(matches) > 1 {
		return matches[1]
	}
	if matches := gdriveQueryRegex.FindStringSubmatch(rawURL); len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func DownloadFile(rawURL, destPath string) error {
	finalURL := rawURL

	if strings.Contains(rawURL, "drive.google.com") || strings.Contains(rawURL, "drive.usercontent.google.com") {
		fileID := ExtractGDriveID(rawURL)
		if fileID != "" {
			log.Infof("google drive url detected (id: %s)", fileID)
			finalURL = fmt.Sprintf("https://drive.usercontent.google.com/download?id=%s&export=download&authuser=0&confirm=t", fileID)
		}
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return fmt.Errorf("failed to create a cookie jar: %w", err)
	}

	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create a request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64)")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("download request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned bad statusL: %s", resp.Status)
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	out, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write content to a file: %w", err)
	}

	return nil
}
