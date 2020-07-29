# Wombat - cross platform gRPC client

<img src="screenshot.png" width="100%" />

**WARNING** - alpha software; this is in active development and is not fully-featured yet.

## Features

- Automatic parsing of proto definitions to render services and input messages
- `.proto` file discovery
- Selection of multiple services and methods
- Basic configuration of TLS, including disabling TLS (plain text)
- Input generation for all scalar types
- Input generation for nested messages
- Input generation for enums, including nested
- Input generation for repeated fields
- Support for adding RPC metadata
- Execute unary requests
- Execute server streaming requests
- View response messages
- View RPC Header and Trailer
- View full RPC statistics
- MacOS build

### Features still working on:

- [ ] Error messages (will siliently fail on error)
- [ ] Windows build
- [ ] Linux build
- [ ] Support for client streaming
- [ ] Support for bidirectional streaming
- [ ] Support for `oneof` fields
- [ ] Unsetting nested messages
- [ ] Reflection API support
- [ ] Multiple Workspaces
- [ ] Multiple Request's within a Workspace

## Download

Visit the [Releases](releases) page for the latest downloads. 
