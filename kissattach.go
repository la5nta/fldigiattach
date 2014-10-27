package main

import (
	"fmt"
	krpty "github.com/kr/pty"
	"os"
	"os/exec"
	"time"
)

func KissAttach(port string, mtu int) (pty *os.File, err error) {
	pty, tty, err := krpty.Open()
	if err != nil {
		return nil, err
	}

	time.Sleep(100 * time.Millisecond)

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

	// Wait for kissattach to open tty
	time.Sleep(2 * time.Second)

	err = tty.Close()
	return pty, err
}
