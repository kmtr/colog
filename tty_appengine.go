// +build appengine

package colog

// Not applicable in appengine
// define constant to avoid compilation error
const ioctlReadTermios = 0x0

var isTerminal = isTerminalFunc
var terminalWidth = terminalWidthFunc

// isTerminalFunc returns true if the given file descriptor is a terminal.
func isTerminalFunc(fd int) bool {
	return false
}

// terminalWidthFunc returns the width in characters of the terminal.
func terminalWidthFunc(fd int) (width int) {
	return -1
}
