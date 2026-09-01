//go:build darwin || freebsd || netbsd || openbsd || dragonfly

package tcp

import "golang.org/x/sys/unix"

func readNonblocking(fd int, dst []byte) (int, error) {
	return unix.Read(fd, dst)
}

func writeNonblocking(fd int, src []byte) (int, error) {
	return unix.Write(fd, src)
}

func writevNonblocking(fd int, src [][]byte) (int, error) {
	return unix.Writev(fd, src)
}
