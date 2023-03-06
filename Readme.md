# NTP version 4 and SNTP

# Usage

```go
package main

import (
	"fmt"

	ntpc "github.com/haikelfazzani/core/ntpc/core"
)

func main() {
	timeNow, _ := ntpc.NtpQuery("time.google.com", "123")
	fmt.Println("NTP time:", timeNow)
}
```

[](https://gist.github.com/mutin-sa/eea1c396b1e610a2da1e5550d94b0453)

# License

MIT
