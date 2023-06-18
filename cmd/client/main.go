package main

import (
	"fmt"
	"time"

	"github.com/ksusonic/gophkeeper/internal/client/cli"
)

var (
	Version = "dev"
	Date    = time.Now().String()
)

func main() {
	fmt.Print(cli.Greeting)
	fmt.Printf("Version: %s from %s\n", Version, Date)
}
