# NTP Client

# Usage

```go
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

	// update system date and time
	isUpdated := ntpc.UpdateSystem(dateTime)
	fmt.Println("> System time updated:", isUpdated)
}
```

[rfc5905](https://datatracker.ietf.org/doc/rfc5905/)
[Public Time Servers](https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453)

# License

MIT
