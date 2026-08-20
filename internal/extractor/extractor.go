package extractor

import (
	"archive/zip"
	"astmn/internal/log"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	MaxUncompressedSize int64 = 20 * 1024 * 1024 * 1024
	MaxFileCount        int   = 100000
)

type ExtractedFileInfo struct {
	Name    string
	RelPath string
	Size    int64
	Hash    string
}

func ExtractArchive(archivePath, targetDir string) ([]ExtractedFileInfo, error) {
	archiveType, err := DetectArchiveType(archivePath)
	if err != nil {
		return nil, fmt.Errorf("failed to detect archive type: %w", err)
	}

	switch archiveType {
	case "zip":
		return extractZip(archivePath, targetDir)
	//TODO: other archive files extracting
	case "7z":
		return nil, fmt.Errorf("7zip is not implemented yet")
	case "gzip":
		return nil, fmt.Errorf("gzip is not implemented yet")
	case "tar":
		return nil, fmt.Errorf("tar is not implemented yet")
	case "rar":
		return nil, fmt.Errorf("rar is not implemented yet")
	default:
		return nil, fmt.Errorf("unsupported or corrupted archive format %s", archivePath)
	}
}

func extractZip(archivePath, targetDir string) ([]ExtractedFileInfo, error) {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip archive: %w", err)
	}
	defer r.Close()

	zipFileAmount := len(r.File)
	if zipFileAmount > MaxFileCount {
		return nil, fmt.Errorf("archive contains too many files (%d > %d)", zipFileAmount, MaxFileCount)
	}

	var totalSize int64
	for _, f := range r.File {
		totalSize += int64(f.UncompressedSize64)
	}
	if totalSize > MaxUncompressedSize {
		return nil, fmt.Errorf("the uncompressed archive size exceeds limit (%d bytes > %d limit)", totalSize, MaxUncompressedSize)
	}

	targetDirAbs, err := filepath.Abs(targetDir)
	if err != nil {
		return nil, fmt.Errorf("invalid target directory: %w", err)
	}

	var extractedFiles []ExtractedFileInfo
	for _, f := range r.File {
		destPath := filepath.Join(targetDirAbs, f.Name)
		cleanDestPath := filepath.Clean(destPath)

		if !strings.HasPrefix(cleanDestPath, targetDirAbs+string(os.PathSeparator)) && cleanDestPath != targetDirAbs {
			return nil, fmt.Errorf("illegal zip file path (Zip-Slip attempt): %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(cleanDestPath, os.ModePerm); err != nil {
				return nil, fmt.Errorf("failed to create directory %s: %w", cleanDestPath, err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(cleanDestPath), os.ModePerm); err != nil {
			return nil, fmt.Errorf("failed to create parent dir for %s: %w", cleanDestPath, err)
		}

		rc, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to open zip entry %s: %w", f.Name, err)
		}

		fileInfo, err := saveAndHashStream(rc, cleanDestPath)
		rc.Close()
		if err != nil {
			return nil, err
		}

		relPath, _ := filepath.Rel(targetDirAbs, cleanDestPath)
		fileInfo.RelPath = relPath

		extractedFiles = append(extractedFiles, fileInfo)
	}

	log.Infof("successfully extracted %d files to %s", len(extractedFiles), targetDir)
	return extractedFiles, nil
}

func saveAndHashStream(src io.Reader, destPath string) (ExtractedFileInfo, error) {
	out, err := os.Create(destPath)
	if err != nil {
		return ExtractedFileInfo{}, fmt.Errorf("failed to create target file %s: %w", destPath, err)
	}
	defer out.Close()

	hasher := sha256.New()
	limitReader := io.LimitReader(src, MaxUncompressedSize)
	writer := io.MultiWriter(out, hasher)

	written, err := io.Copy(writer, limitReader)
	if err != nil {
		return ExtractedFileInfo{}, fmt.Errorf("failed to write file %s: %w", destPath, err)
	}

	return ExtractedFileInfo{
		Name: filepath.Base(destPath),
		Size: written,
		Hash: hex.EncodeToString(hasher.Sum(nil)),
	}, nil
}
