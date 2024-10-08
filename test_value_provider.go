package value_provider

import (
  "fmt"
  "math/rand"
  "math"
)

func Bool() bool {
  if rand.Int() % 2 == 0 {
    return true
  } else {
    return false
  }
}

func Int() int {
  return rand.Int()
}

func UpperLimitedInt(upperLimit int) int {
  if upperLimit > math.MaxInt {
    panic(fmt.Sprintf("Too big intiger. Max int is: %d", math.MaxInt))
  }

  return rand.Int()
}
