# NTP
Simple implementation of NTP client

# Install

```
go get github.com/haikelfazzani/ntpc
```

# Usage

```go
package main

import (
	"fmt"

	ntpc "github.com/haikelfazzani/ntpc/core"
)

func main() {
	ntpc := &ntpc.NewClient{
		Server:  "pool.ntp.org",
		Port:    "123",
		Timeout: 10,
	}

	timeNow, _ := ntpc.Query()
	fmt.Println("\n> NTP time:", timeNow.Local())

	dateTime := timeNow.Format("2006-01-02 15:04:05")

	isUpdated := ntpc.UpdateSystem(dateTime)
	fmt.Println("\n> System time updated:", isUpdated)
}
```

[rfc5905](https://datatracker.ietf.org/doc/rfc5905/)
[Public Time Servers](https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453)

# License

MIT
