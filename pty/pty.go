package pty

import (
	"os/exec"
)

// PTY is the interface for pseudoterminal
type PTY interface {
	Start(c *exec.Cmd, rows, cols int) error
	Close() error
	SetSize(rows, cols int) error
	GetSize() (rows, cols int, err error)
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}
