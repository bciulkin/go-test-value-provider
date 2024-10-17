package value_provider

import (
  "fmt"
  "math/rand"
  "math"
)

const mixedCaseLetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const numericString = "0123456789"
const lowerCaseLetterBytes = "abcdefghijklmnopqrstuvwxyz"
const upperCaseLetterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Bool() bool {
  return IntN(2) == 1
}

func Int() int {
  return rand.Int()
}

func IntN(max int) int {
  if max > math.MaxInt {
    panic(fmt.Sprintf("Too big intiger. Max int is: %d", math.MaxInt))
  }

  return rand.Intn(max)
}

func IntNM(min, max int) int {
  if min > max {
    panic("Min is bigger than max.")
  }
  return IntN(max) + min
}

func Float32() float32 {
  return rand.Float32()
}

func Float32N(max float32) float32 {
  if max > math.MaxFloat32 {
    panic(fmt.Sprintf("Too big float32. Max float32 is: %f", math.MaxFloat32))
  }

  return Float32() * max
}

func Float64() float64 {
  return rand.Float64()
}

func Float64N(max float64) float64 {
  if max > math.MaxFloat64 {
    panic(fmt.Sprintf("Too big float64. Max float64 is: %f", math.MaxFloat64))
  }

  return Float64() * max
}


func String() string {
  return StringN(10)
}

func StringN(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = mixedCaseLetterBytes[rand.Intn(len(mixedCaseLetterBytes))]
  }
  return string(b)
}

func ByteArrayN(n int) []byte {
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(IntN(256))
  }
  return b
}

func ByteArray() []byte {
  return ByteArrayN(IntNM(1, 100))
}
