# browser_go

`browser_go` is a lightweight, cross-platform Go package that lets you open URLs in the user's default web browser. It supports Windows, macOS, Linux, and Windows Subsystem for Linux (WSL).

## Why use browser_go?

When developing web applications or APIs (for example, with the Gin web framework), you often want your dashboard or homepage to open automatically in the browser as soon as your server starts.  
`browser_go` makes this effortless and improves your development workflow.

---

## Features

- **Cross-platform**: Works seamlessly on Windows, macOS, Linux, and WSL.
- **WSL-aware**: Detects WSL and opens URLs in the native Windows browser.
- **Simple API**: Just call `OpenURL` with your URL.
- **No external dependencies**: Uses only Go’s standard library and built-in system utilities.

---

## Installation

Install the package with:

```bash
go get github.com/Abin-Antu/browser_go
```

---

## Usage

You can use `browser_go` to open any web page—such as your dashboard, API UI, or documentation—automatically when your Go web server starts.

### Example: Open a dashboard or homepage when your server starts

```go
package main

import (
    "github.com/Abin-Antu/browser_go/browser"
)

func main() {
    url := "http://localhost:8080/dashboard"
    browser.OpenURL(url) // This will open the dashboard in the default browser

    // Start your web server here (example with Gin or net/http)
    // router := gin.Default()
    // router.Run(":8080")
}
```

Just call `browser.OpenURL(url)` with the URL you want to open.  
This is especially handy for automatically launching your web dashboard, homepage, or API documentation for local development.

---

## License

MIT License. See [LICENSE](LICENSE) for details.