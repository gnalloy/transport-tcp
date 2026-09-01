//go:build linux

package tcp

import (
	"testing"

	"golang.org/x/sys/unix"
)

func TestRawNonblockingReadWrite(t *testing.T) {
	pair := newNonblockingSocketPair(t)
	payload := []byte("gnalloy")

	n, err := writeNonblocking(pair[0], payload)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(payload) {
		t.Fatalf("write bytes=%d, want %d", n, len(payload))
	}
	dst := make([]byte, len(payload))
	n, err = readNonblocking(pair[1], dst)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(payload) || string(dst) != string(payload) {
		t.Fatalf("read bytes=%d payload=%q, want %d/%q", n, dst, len(payload), payload)
	}
}

func TestRawNonblockingReadReturnsAgain(t *testing.T) {
	pair := newNonblockingSocketPair(t)
	dst := make([]byte, 1)

	n, err := readNonblocking(pair[0], dst)
	if n != 0 || !isAgain(err) {
		t.Fatalf("read bytes=%d err=%v, want 0/EAGAIN", n, err)
	}
}

func TestRawNonblockingWritev(t *testing.T) {
	pair := newNonblockingSocketPair(t)
	parts := [][]byte{nil, []byte("gn"), {}, []byte("alloy")}

	n, err := writevNonblocking(pair[0], parts)
	if err != nil {
		t.Fatal(err)
	}
	if n != len("gnalloy") {
		t.Fatalf("writev bytes=%d, want %d", n, len("gnalloy"))
	}
	dst := make([]byte, len("gnalloy"))
	n, err = readNonblocking(pair[1], dst)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(dst) || string(dst) != "gnalloy" {
		t.Fatalf("read bytes=%d payload=%q", n, dst)
	}
}

func TestRawNonblockingWritevLimitsSingleBatch(t *testing.T) {
	pair := newNonblockingSocketPair(t)
	parts := make([][]byte, maxRawWritevBuffers+1)
	for i := range parts {
		parts[i] = []byte{byte(i)}
	}

	n, err := writevNonblocking(pair[0], parts)
	if err != nil {
		t.Fatal(err)
	}
	if n != maxRawWritevBuffers {
		t.Fatalf("writev bytes=%d, want %d", n, maxRawWritevBuffers)
	}
	dst := make([]byte, maxRawWritevBuffers)
	n, err = readNonblocking(pair[1], dst)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(dst) || dst[len(dst)-1] != byte(maxRawWritevBuffers-1) {
		t.Fatalf("read bytes=%d last=%d", n, dst[len(dst)-1])
	}
}

func TestRawNonblockingEmptyBuffers(t *testing.T) {
	if n, err := readNonblocking(-1, nil); n != 0 || err != nil {
		t.Fatalf("empty read bytes=%d err=%v", n, err)
	}
	if n, err := writeNonblocking(-1, nil); n != 0 || err != nil {
		t.Fatalf("empty write bytes=%d err=%v", n, err)
	}
	if n, err := writevNonblocking(-1, [][]byte{nil, {}}); n != 0 || err != nil {
		t.Fatalf("empty writev bytes=%d err=%v", n, err)
	}
}

func newNonblockingSocketPair(t *testing.T) [2]int {
	t.Helper()
	pair, err := unix.Socketpair(unix.AF_UNIX, unix.SOCK_STREAM|unix.SOCK_NONBLOCK|unix.SOCK_CLOEXEC, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = unix.Close(pair[0])
		_ = unix.Close(pair[1])
	})
	return pair
}
