// main.go
package main

import (
    "fmt"
    "os"

    "github.com/paywithclay/clay-relay/cli"
)

func main() {
    fmt.Println("📦 Welcome to Clay Relay — Your Private Payment Gateway Bridge")
    if len(os.Args) < 2 {
        fmt.Println("❌ No command provided. Use: setup | start")
        return
    }

    command := os.Args[1]

    switch command {
    case "setup":
        cli.Setup()
    case "start":
        cli.Start()
    default:
        fmt.Println("❌ Unknown command. Use: setup | start")
    }
}