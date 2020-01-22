package main

import (
    "fmt"
    "os/exec"
    "strings"
)

func main() {
    var resolved, unresolved int
    cmd := exec.Command("hg", "resolve", "--list")
    out, err := cmd.Output()
    if err != nil {
        return
    }
    lines := strings.Split(string(out), "\n")
    for _, line := range(lines) {
        e := strings.Split(line, " ")
        if len(e) != 2 {
            continue
        }
        switch e[0] {
        case "R":
            resolved++
        case "U":
            unresolved++
        }
    }
    fmt.Printf("%d|%d", resolved, unresolved)
}
