//go:build linux

package tcp

import (
	"runtime"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

const maxRawWritevBuffers = 64

// 已注册到 readiness poller 的 fd 必须非阻塞，raw syscall 可避免调度器切换 P。
func readNonblocking(fd int, dst []byte) (int, error) {
	if len(dst) == 0 {
		return 0, nil
	}
	for {
		n, _, errno := syscall.RawSyscall(
			unix.SYS_READ,
			uintptr(fd),
			uintptr(unsafe.Pointer(unsafe.SliceData(dst))),
			uintptr(len(dst)),
		)
		runtime.KeepAlive(dst)
		if errno == unix.EINTR {
			continue
		}
		if errno != 0 {
			return 0, errno
		}
		return int(n), nil
	}
}

func writeNonblocking(fd int, src []byte) (int, error) {
	if len(src) == 0 {
		return 0, nil
	}
	for {
		n, _, errno := syscall.RawSyscall(
			unix.SYS_WRITE,
			uintptr(fd),
			uintptr(unsafe.Pointer(unsafe.SliceData(src))),
			uintptr(len(src)),
		)
		runtime.KeepAlive(src)
		if errno == unix.EINTR {
			continue
		}
		if errno != 0 {
			return 0, errno
		}
		return int(n), nil
	}
}

func writevNonblocking(fd int, src [][]byte) (int, error) {
	var iovecs [maxRawWritevBuffers]unix.Iovec
	count := 0
	for _, part := range src {
		if len(part) == 0 {
			continue
		}
		iovecs[count].Base = unsafe.SliceData(part)
		iovecs[count].SetLen(len(part))
		count++
		if count == len(iovecs) {
			break
		}
	}
	if count == 0 {
		return 0, nil
	}
	for {
		n, _, errno := syscall.RawSyscall(
			unix.SYS_WRITEV,
			uintptr(fd),
			uintptr(unsafe.Pointer(&iovecs[0])),
			uintptr(count),
		)
		runtime.KeepAlive(src)
		if errno == unix.EINTR {
			continue
		}
		if errno != 0 {
			return 0, errno
		}
		return int(n), nil
	}
}
