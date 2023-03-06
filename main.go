package main

import (
	"fmt"
)

func main() {
	timeNow, _ := NtpQuery("time.google.com", "123")
	fmt.Println("NTP time:", timeNow)
}
