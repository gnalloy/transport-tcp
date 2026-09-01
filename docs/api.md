# API Reference

[简体中文](api.zh-CN.md) | [Docs Index](README.md)

This inventory is generated from `go doc -short` for the packages in this repository. It is a quick public-surface map; source files and tests remain the authority for exact semantics.

## Packages

### `gnalloy.org/transport-tcp`

Package name: `tcp`

```text
var ErrInvalidAddress = errors.New("gnalloy/transport/tcp: invalid tcp address") ...
type AllocatorFactory func(loop *transport.EventLoop) (buffer.Allocator, error)
    func NewMmapAllocatorFactory(cfg buffer.MmapAllocatorConfig, fallbackToHeap bool) AllocatorFactory
type Config struct{ ... }
    func DefaultConfig() Config
type Server struct{ ... }
type Transport struct{ ... }
    func NewTransport(cfg Config) *Transport
```
