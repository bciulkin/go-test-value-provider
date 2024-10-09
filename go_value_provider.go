package value_provider

import (
  "fmt"
  "math/rand"
  "math"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

  return rand.Intn(upperLimit)
}

func String() string {
  b := make([]byte, 10)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))] // Pick random character
  }
  return string(b)
}

func StringN(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))] // Pick random character
  }
  return string(b)
}
