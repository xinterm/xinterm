package screen

// Screen is the final display
type Screen struct {
	windows      []Window
	activeWindow int
}

// New creates a new screen
func New() *Screen {
	screen := &Screen{}

	screen.windows = []Window{TerminalWindow{}}
	screen.activeWindow = 0

	return screen
}

func (s *Screen) CreateWindow() {

}
