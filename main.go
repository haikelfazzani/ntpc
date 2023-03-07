package main

import (
	"fmt"

	ntpc "github.com/haikelfazzani/ntpc/core"
)

func main() {
	ntpc := &ntpc.NTPC{
		Server: "pool.ntp.org",
		Port:   "123",
	}
	timeNow, _ := ntpc.Query()
	fmt.Println("NTP time:", timeNow)
}
