module github.com/xinterm/xinterm

go 1.14

require (
	github.com/creack/pty v1.1.9
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/nsf/termbox-go v0.0.0-20200418040025-38ba6e5628f1
	github.com/sirupsen/logrus v1.5.0
	github.com/xinterm/terminal v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71
	golang.org/x/text v0.3.2
)

replace github.com/xinterm/terminal => ../terminal
