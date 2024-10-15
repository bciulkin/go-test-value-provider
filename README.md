## Summary
Simple value provider interface that could be used for testing purposes.

### Setup
`go get https://github.com/bciulkin/go-test-value-provider`

### Example

```
import (
    "github.com/bciulkin/go-test-value-provider"
    "fmt"
)


func example() {
    fmt.Println(value_provider.IntN(10)) // random integer from 0 to 10
    fmt.Println(value_provider.String()) // random 10 character length string
}
```
