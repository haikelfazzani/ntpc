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
	fmt.Println("> NTP time:", timeNow) // 2023-03-07 21:01:53.084929939 +0000 UTC

	dateTime := timeNow.Format("2006-01-02 15:04:05")

	isUpdated := ntpc.UpdateSystem(dateTime)
	fmt.Println("> System time updated:", isUpdated)
}
