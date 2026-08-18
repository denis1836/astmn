package extractor

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// magic bytes
var (
	zipMagic      = []byte{0x50, 0x4B, 0x03, 0x04}
	zipEmptyMagic = []byte{0x50, 0x4B, 0x05, 0x06}
	sevenZipMagic = []byte{0x37, 0x7A, 0xBC, 0xAF, 0x27, 0x1C}
	gzipMagic     = []byte{0x1F, 0x8B, 0x08}
	tarMagic      = []byte{0x75, 0x73, 0x74, 0x61, 0x72}
	rarMagic      = []byte{0x52, 0x61, 0x72, 0x21, 0x1A, 0x07}
)

func DetectArhiveType(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file for type detection: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read file header: %v", err)
	}

	header := buf[:n]

	switch {
	case bytes.HasPrefix(header, zipMagic) || bytes.HasPrefix(header, zipEmptyMagic):
		return "zip", nil
	case bytes.HasPrefix(header, sevenZipMagic):
		return "7z", nil
	case bytes.HasPrefix(header, gzipMagic):
		return "gzip", nil
	case len(header) >= 262 && bytes.Equal(header[257:262], tarMagic):
		return "tar", nil
	case bytes.HasPrefix(header, rarMagic):
		return "rar", nil
	default:
		return "unknown", nil
	}
}
