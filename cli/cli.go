package cli

import (
	"fmt"

	ntpc "github.com/haikelfazzani/ntpc/core"
)

func Execute() {
	ntpc := &ntpc.NewClient{
		Server:  "pool.ntp.org",
		Port:    "123",
		Timeout: 10,
	}

	timeNow, _ := ntpc.Query()
	fmt.Println("\n> NTP time:", timeNow.Local()) // 2023-03-07 21:01:53.084929939 +0000 UTC

	// dateTime := timeNow.Format("2006-01-02 15:04:05")

	// isUpdated := ntpc.UpdateSystem(dateTime)
	// fmt.Println("\n> System time updated:", isUpdated)
}
