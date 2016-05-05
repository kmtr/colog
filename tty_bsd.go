// +build darwin dragonfly freebsd netbsd openbsd
// +build !appengine

package colog

import "syscall"

const ioctlReadTermios = syscall.TIOCGETA
