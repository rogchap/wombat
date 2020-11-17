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

During development it's common to want to test proto files and gRPC APIs. When Wombat lanuches in "Bridge/Debug Mode" it
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

## Frontend

The frontend is built with [Svelte](https://svelte.dev/); and should be fairly straight forward if you have done any
type of web-based frontend development before.

During development, you can use any browser (and it's dev-tools), but it's important to remember that Wombat does not
bundle a browser (like Chromium), but rather uses the system's built-in web view component. For Mac this will be a
Safari web view; on Linux a WebKit web view and on Windows mshtml (basically IE11).

mshtml does not have support for CSS variables (which is used heavily by Wombat frontend); because of this limitation we
pre-compile the CSS via a plugin for rollup. This only works for CSS declared in `style` blocks; to use in the JS world
we inject a `window.isWin` bool flag (just `isWin` for short) so that we can pass the real value rather than the CSS
variable.

This is a short term fix; Wails v2 will built against
[WebView2](https://docs.microsoft.com/en-us/microsoft-edge/webview2/) which is based off Chromium and will resolve most
of the Windows "issues".

If you need to add any styles please update both `frontend/src/views/App.svelte` and `frontend/rollup.config.js`.

After getting your UI correct; it's always good to validate in a real build:

```zsh
$ wails build -d
$ ./build/wombat
```

The `-d` flag will build Wombat in debug mode so you will still get all the debug output in the terminal, but also runs
the development gRPC server too.

In debug mode you can still "Right/Option Click" -> "Inspect Element" and bring up the development tools (mac and linux
only). This is disabled in production builds.

## Changelog

We maintain a changelog for all versions of Wombat. If you create a PR please add the change to `CHANGELOG.md`. Don't
forget to add your GitHub handle to make sure you get the credit!
