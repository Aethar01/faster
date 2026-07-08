package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Aethar01/faster/internal/faster"
)

func main() {
	if err := faster.Run(os.Args[1:], os.Stdout); err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) {
			fmt.Fprintln(os.Stderr, "No internet connection.")
			os.Exit(1)
		}
		log.Fatal(err)
	}
}
