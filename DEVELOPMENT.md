# Development

## Prerequisites 

- Working [Go](https://golang.org/) 1.15+ environment
- Working [Node.js](https://nodejs.org/en/) v12+ environment

## Getting started

Wombat is built with [Wails](https://wails.app/); to install the CLI:

```zsh
$ go get github.com/wailsapp/wails/cmd/wails
```

Wails has a "Bridge Mode" that allows you to debug with your browser of choice; to enable:

```zsh
$ wails serve
```

In a separate terminal you will need to start the frontend:

```zsh
$ cd frontend
$ npm run dev
```

The frontend will automatically re-build and reload on changes.

## Built-in gRPC server

During development it's common to want to test proto files and gRPC APIs. When Wombat lanuches in "Bridge Mode" it
automatically starts a gRPC server running on `localhost:5001`. You can also make changes to the
`/internal/server/foobar.proto` file to add test protos.

## Backend API

The Go backend has a singular API entry: `app.api{}`; All public methods defined on the `api` struct are exposed to the
frontend.

Wombat follows a (loose) [flux](https://facebook.github.io/flux/) like architecture. API calls are made to the backend
(actions) and the backend will then emit events (dispatcher) for any changes required by the frontend.

Although the backend APIs can responed with an object/error; the frontend mostly ignores these objects/errors (uncaught promise error), and
rather responds to events including error events.

All backend event constants are defined in `/internal/app/events.go`.
