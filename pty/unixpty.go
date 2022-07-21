package pty

import (
	"errors"
	"os"
	"os/exec"

	unixpty "github.com/creack/pty"
)

// UnixPTY deals with unix pty
type UnixPTY struct {
	ptmx *os.File
}

// Start a unix pty
func (pty *UnixPTY) Start(c *exec.Cmd, rows, cols int) error {
	var err error
	pty.ptmx, err = unixpty.StartWithSize(c, &unixpty.Winsize{
		Rows: uint16(rows),
		Cols: uint16(cols),
	})
	return err
}

// Close a unix pty
func (pty *UnixPTY) Close() error {
	if pty.ptmx == nil {
		return errors.New("Unix PTY not opened")
	}
	err := pty.ptmx.Close()
	pty.ptmx = nil
	return err
}

// SetSize set the pty size
func (pty *UnixPTY) SetSize(rows, cols int) error {
	if pty.ptmx == nil {
		return errors.New("Unix PTY not opened")
	}
	return unixpty.Setsize(pty.ptmx, &unixpty.Winsize{
		Rows: uint16(rows),
		Cols: uint16(cols),
	})
}

// GetSize set the pty size
func (pty *UnixPTY) GetSize() (rows, cols int, err error) {
	if pty.ptmx == nil {
		return 0, 0, errors.New("Unix PTY not opened")
	}
	return unixpty.Getsize(pty.ptmx)
}

// Read implements io.Reader
func (pty *UnixPTY) Read(p []byte) (n int, err error) {
	if pty.ptmx == nil {
		return 0, errors.New("Unix PTY not opened")
	}
	return pty.ptmx.Read(p)
}

// Write implements io.Reader
func (pty *UnixPTY) Write(p []byte) (n int, err error) {
	if pty.ptmx == nil {
		return 0, errors.New("Unix PTY not opened")
	}
	return pty.ptmx.Write(p)
}
