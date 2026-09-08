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

	io.WriteString(stdin, content)
	stdin.Close()

	return cmd.Wait()
}
