# transport-tcp

Native TCP stream transport for Gnalloy, including ServerBootstrap and Dialer integration.

This repository is part of the Gnalloy modular networking stack. The default branch is `dev`; no release tag is created during bootstrap.

## Install

```bash
go get gnalloy.org/transport-tcp@dev
```

## Module Boundary

- Module path: `gnalloy.org/transport-tcp`
- Responsibility: Native TCP stream transport for Gnalloy, including ServerBootstrap and Dialer integration
- Core dependency: `gnalloy.org/gnalloy` when this module uses Gnalloy buffers, channels, event loops, or bootstrap contracts.

## Gnalloy Dependencies

- `gnalloy.org/gnalloy`
- `gnalloy.org/transport-zerocopy`

## Development

```bash
go test ./... -count=1
go vet ./...
go test ./... -run '^$' -bench . -benchmem -benchtime=100ms -count=1
```

For multi-repository development, use the workspace at `G:\opensource\gnalloy\go.work`. For standalone verification, set `GOWORK=off`.

## License

Apache-2.0.
