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

func IntN(upperLimit int) int {
  if upperLimit > math.MaxInt {
    panic(fmt.Sprintf("Too big intiger. Max int is: %d", math.MaxInt))
  }

  return rand.Intn(upperLimit)
}

func Float32() float32 {
  return rand.Float32()
}

func Float32N(upperLimit float32) float32 {
  if upperLimit > math.MaxFloat32 {
    panic(fmt.Sprintf("Too big float32. Max float32 is: %f", math.MaxFloat32))
  }

  return rand.Float32() * upperLimit
}

func Float64() float64 {
  return rand.Float64()
}

func Float64N(upperLimit float64) float64 {
  if upperLimit > math.MaxFloat64 {
    panic(fmt.Sprintf("Too big float64. Max float64 is: %f", math.MaxFloat64))
  }

  return rand.Float64() * upperLimit
}


func String() string {
  b := make([]byte, 10)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}

func StringN(length int) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = letterBytes[rand.Intn(len(letterBytes))]
  }
  return string(b)
}


