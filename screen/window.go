package screen

import (
	"github.com/xinterm/terminal"
)

// Window is the interface for a window
type Window interface {
	Write()
	Input()
	Event()
	Draw()
}

// TerminalWindow contains terminal
type TerminalWindow struct {
	term *terminal.Terminal
}
