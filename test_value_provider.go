package main

import (
  "fmt"
)

type TestValueProvider interface {
  Int() int
}

type TestValueProvider {

}

func NewTestValueProvider() TestValueProvider {
  return &TestValueProvider{
  }
}
