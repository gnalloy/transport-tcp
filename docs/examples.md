# Examples

[简体中文](examples.zh-CN.md) | [Docs Index](README.md)

## Example 1: Add the Module to an Application

```bash
mkdir gnalloy-app && cd gnalloy-app
go mod init example.com/gnalloy-app
go get gnalloy.org/transport-tcp@dev
go doc gnalloy.org/transport-tcp
```

## Example 2: Inspect Current Packages

The current source tree exposes these package import paths:
- `gnalloy.org/transport-tcp`

Use `go doc` against the package that matches the behavior you need:

```bash
go doc gnalloy.org/transport-tcp
```

Selected current exported entry points:
- `var ErrInvalidAddress = errors.New("gnalloy/transport/tcp: invalid tcp address") ...`
- `type AllocatorFactory func(loop *transport.EventLoop) (buffer.Allocator, error)`
- `type Config struct{ ... }`
- `type Server struct{ ... }`
- `type Transport struct{ ... }`

## Example 3: Use Executable Tests as Behavioral Examples

Repository tests are executable examples of supported behavior. Start with the selected names below, then read the matching `_test.go` files for complete setup and assertions. See [Testing and Performance](testing.md) for the complete discovered list.

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -run Test -count=1
```

Selected current test, benchmark, fuzz, and example entry points:
- `BenchmarkLengthFieldTCPRoundTrip`
- `BenchmarkNativeTCPEchoRoundTrip`
- `TestBenchmarkPayloadUsesRequestedSize`
- `TestBindFailureClosesPartiallyRegisteredReusePortListeners`
- `TestLengthFieldTCPHandlesSplitAndStickyFrames`
- `TestMmapAllocatorFactoryFallback`
- `TestNativeTCPEchoLifecycle`
- `TestNativeTCPLongConnectionsStability`
- `TestNativeTCPRepeatedConnectClose`
- `TestNormalizeConfigCarriesWriteBufferWatermark`
- `TestNormalizeConfigFillsNettyStyleSocketDefaults`
- `TestReusePortBindsOneListenerPerBoss`
- `TestSocketOptionsApplyChildChannelOptions`
- `TestSocketOptionsApplyClientConnectTimeout`
- `TestSocketOptionsApplyListenChannelOptions`
- `TestTCPAllocatorFactoryIsBoundPerWorker`
- `TestTCPDialerEcho`
- `TestTCPFileRegionWritesNativeFileToSocket`

## Example 4: Cross-Module Assembly

Direct Gnalloy dependencies for this module:
- `gnalloy.org/gnalloy`
- `gnalloy.org/transport-zerocopy`

Assembly guidance:
- Use this transport to own the concrete I/O endpoint and connect it to Gnalloy Channel and EventLoop contracts.
- Protocol parsing stays in codec modules and policy stays in handler modules.
- Platform-specific capability checks should happen during application startup and integration tests.

## Example 5: Pressure-Test Harness

For sustained load, wire this module into a scenario under `gnalloy.org/benchmarks` or a runnable client under `gnalloy.org/examples` when the module participates in network traffic. Record host, OS, CPU, Go version, protocol, payload, concurrency, warmup, repetitions, throughput, and p99 latency in the report.
