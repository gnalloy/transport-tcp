# Overview

[简体中文](overview.zh-CN.md) | [Docs Index](README.md)

## Purpose

Native TCP stream transport for Gnalloy, including ServerBootstrap and Dialer integration.

This module owns an I/O boundary. It creates or adapts Gnalloy Channels for a concrete transport while protocol parsing, business handlers, TLS policy, and observability remain in other modules.

## Repository Identity

- Module path: `gnalloy.org/transport-tcp`
- GitHub repository: `github.com/gnalloy/transport-tcp`
- Default branch: `dev`
- License: Apache-2.0

## Package Map
- `gnalloy.org/transport-tcp` (`tcp`)

## Direct Gnalloy Dependencies

- `gnalloy.org/gnalloy`
- `gnalloy.org/transport-zerocopy`

## Direct Dependents in the Current Repository Set

- `gnalloy.org/examples`

## Architecture Position

Gnalloy keeps the core small and dependency-light. This repository is a replaceable module around one responsibility, connected through explicit Go packages instead of runtime discovery.

## Compatibility

The public import path is `gnalloy.org/transport-tcp`. Until the first stable tag is published, use `@dev` or an explicit pseudo-version selected by your dependency policy.
