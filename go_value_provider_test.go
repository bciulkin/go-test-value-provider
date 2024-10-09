package value_provider

import (
  "testing"
  "math"
  "github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
  result := Int()

  assert.GreaterOrEqual(t, result, math.MinInt, "Result should be greater or equal than minInt")
  assert.LessOrEqual(t, result, math.MaxInt, "Result should be less or equal than maxInt")
}

func TestUpperLimitedInt(t *testing.T) {
  result := UpperLimitedInt(123)

  assert.LessOrEqual(t, result, 123, "Result should be less or equal than limited value")
  assert.GreaterOrEqual(t, result, math.MinInt, "Result should be greater or equal than minInt")
}
