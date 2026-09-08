package viewer

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func DisplayInPager(content string) error {
	pager := os.Getenv("PAGER")
	if pager == "" {
		pager = "less"
	}

	cmd := exec.Command(pager, "-R")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Print(content)
		return nil
	}

	if err := cmd.Start(); err != nil {
		fmt.Print(content)
		return err
	}

	_, writeErr := io.WriteString(stdin, content)
	_ = stdin.Close()

	if writeErr != nil {
		return fmt.Errorf("failed to write content to pager: %w", writeErr)
	}

	return cmd.Wait()
}
