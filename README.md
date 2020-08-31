<p align="center">
<img src="assets/gh/title.png" width="400px" alt="Wombat" />
<h3 align="center">Cross platform gRPC client</h3>
</p>

<p align="center">
<a href="https://github.com/rogchap/wombat/releases"><img src="https://img.shields.io/github/v/release/rogchap/wombat?include_prereleases&style=flat-square" alt="Github release"></a>
<img alt="Go version" src="https://img.shields.io/github/go-mod/go-version/rogchap/wombat?style=flat-square">
</p>

<img src="assets/gh/screenshot.png" width="100%" alt="screenshot" />

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
- Execute client streaming requests
- Execute bidirectional requests
- Cancel requests in-flight
- Send EOF for client streaming and bi-directional streaming
- View response messages
- View RPC Header and Trailer
- View full RPC statistics
- MacOS build
- Linux build (inc AppImage)
- Reflection API to determine RPC schema
- Error messages on failed gRPC connections

### Features still working on:

- [ ] Windows build
- [ ] Support for `oneof` fields
- [ ] Multiple Workspaces
- [ ] Multiple Request's within a Workspace

## Download

Visit the [Releases](releases) page for the latest downloads. 

## Install

### MacOS

If you get this error message: `"Wombat.app" can't be opened because the identity of the developer cannot be
confirmed.`, Install by **Right Click/Option Click** -> **Open** -> **Open**. You'll only need to do this on first
install.

### Linux

If you are using the `AppImage`, just download and set the file as executable:

```
$ chmod +x Wombat.AppImage
```
