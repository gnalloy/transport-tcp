# 案例

[English](examples.md) | [文档索引](README.zh-CN.md)

## 案例 1：将模块加入应用

```bash
mkdir gnalloy-app && cd gnalloy-app
go mod init example.com/gnalloy-app
go get gnalloy.org/transport-tcp@dev
go doc gnalloy.org/transport-tcp
```

## 案例 2：查看当前包

当前源码树暴露这些 package 导入路径：
- `gnalloy.org/transport-tcp`

按需要的行为对对应 package 执行 `go doc`：

```bash
go doc gnalloy.org/transport-tcp
```

精选当前导出入口：
- `var ErrInvalidAddress = errors.New("gnalloy/transport/tcp: invalid tcp address") ...`
- `type AllocatorFactory func(loop *transport.EventLoop) (buffer.Allocator, error)`
- `type Config struct{ ... }`
- `type Server struct{ ... }`
- `type Transport struct{ ... }`

## 案例 3：将可执行测试作为行为示例

仓库测试是受支持行为的可执行示例。先从下面的精选名称开始，再阅读对应 `_test.go` 文件中的完整 setup 和断言。完整发现列表见 [测试与性能](testing.zh-CN.md)。

```bash
GOWORK=off GOTOOLCHAIN=local go test ./... -run Test -count=1
```

精选当前 test、benchmark、fuzz 与 example 入口：
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

## 案例 4：跨模块装配

本模块的直接 Gnalloy 依赖：
- `gnalloy.org/gnalloy`
- `gnalloy.org/transport-zerocopy`

装配说明：
- transport 负责具体 I/O endpoint，并连接 Gnalloy Channel 与 EventLoop 契约。
- 协议解析留给 codec 模块，策略能力留给 handler 模块。
- 平台能力检查应放在应用启动和集成测试阶段完成。

## 案例 5：压测 Harness

持续负载测试时，如果该模块参与网络流量路径，将它接入 `gnalloy.org/benchmarks` 的场景，或接入 `gnalloy.org/examples` 的可运行客户端。报告中记录 host、OS、CPU、Go version、protocol、payload、concurrency、warmup、repetitions、throughput 和 p99 latency。
