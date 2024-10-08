package value_provider

import (
  "testing"
  "math"
  "github.com/stretchr/testify/assert"
  "fmt"
  "github.com/bciulkin/go-test-value-provider"
)

func TestInt(t *testing.T) {
  result := value_provider.Int()

  assert.GreaterOrEqual(t, result, math.MinInt, "Result should be greater or equal than minInt")
  assert.LessOrEqual(t, result, math.MaxInt, "Result should be less or equal than maxInt")
}

func UpperLimitedTestInt(t *testing.T) {
  result := value_provider.UpperLimitedInt(123)

  fmt.Println(result)
  assert.GreaterOrEqual(t, result, math.MinInt, "Result should be greater or equal than minInt")
  assert.LessOrEqual(t, result, 123, "Result should be less or equal than limited value")
}
