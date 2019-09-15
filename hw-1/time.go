package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %s", err)
		os.Exit(1)
	}
	fmt.Printf("Hello, now is %s", time)
}
