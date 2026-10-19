package main

import (
	"os"

	"github.com/Paymentbox-com/proto-service-mesh-api/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
