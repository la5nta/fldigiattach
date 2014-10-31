package main

import (
	"fmt"
	krpty "github.com/kr/pty"
	"os"
	"os/exec"
)

func KissAttach(port string, mtu int) (pty *os.File, err error) {
	pty, tty, err := krpty.Open()
	if err != nil {
		return nil, err
	}

	args := []string{tty.Name(), port, "-l"}
	if mtu > 0 {
		args = append(args, fmt.Sprintf("-m %d", mtu))
	}

	c := exec.Command("kissattach", args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err = c.Start()
	if err != nil {
		pty.Close()
		return nil, err
	}

	// Wait for kissattach to daemonize
	if err = c.Wait(); err != nil {
		pty.Close()
		return nil, err
	}

	return pty, tty.Close()
}
