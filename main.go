package main

import (
	"fmt"

	ntpc "github.com/haikelfazzani/ntpc/core"
)

func main() {
	timeNow, _ := ntpc.NtpQuery("time.google.com", "123")
	fmt.Println("NTP time:", timeNow)
}
