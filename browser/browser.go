// Package browser provides utilities to open URLs in the default web browser
// across multiple platforms, including Windows, macOS, Linux, and WSL.
package browser



import (
    "os/exec"
    "runtime"
    "strings"
)

func OpenURL(url string) error {
    var cmd string
    var args []string

    switch runtime.GOOS {
    case "windows":
        cmd = "cmd"
        args = []string{"/c", "start"}
    case "darwin":
        cmd = "open"
        args = []string{url}
    default:
        if isWSL() {
            cmd = "cmd.exe"
            args = []string{"/c", "start", url}
        } else {
            cmd = "xdg-open"
            args = []string{url}
        }
    }

    if len(args) > 1 {
        args = append(args[:1], append([]string{""}, args[1:]...)...)
    }

    return exec.Command(cmd, args...).Start()
}

func isWSL() bool {
    releaseData, err := exec.Command("uname", "-r").Output()
    if err != nil {
        return false
    }
    return strings.Contains(strings.ToLower(string(releaseData)), "microsoft")
}

