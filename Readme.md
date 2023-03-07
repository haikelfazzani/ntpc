# NTPv4 and SNTP Client (rfc5905)

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
	fmt.Println("NTP time:", timeNow)
}
```

[rfc5905](https://datatracker.ietf.org/doc/rfc5905/)
[Public Time Servers](https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453)

# License

MIT
