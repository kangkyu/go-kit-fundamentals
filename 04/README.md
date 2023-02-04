# 04

```sh
go mod init github.com/kangkyu/go-kit-fundamentals/04
```

```go
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getGreeting(t *testing.T) {
	expected := "Welcome to Go kit 0.12 Fundamentals!"

	actual := getGreeting()

	assert.True(t, expected == actual, "~2|Test expected message: %s, not message %s~", expected, actual)
}
```

```sh
go mod tidy
go test
# undefined: getGreeting
```

```go
package main

import (
	"fmt"
)

func getGreeting() string {
	return fmt.Sprintf("Welcome to Go kit 0.12 Fundamentals!")
}
```

```sh
go test
# PASS
```

```go
func main() {
	fmt.Println(getGreeting())
}
```

```sh
go run main.go
# Welcome to Go kit 0.12 Fundamentals!
```
