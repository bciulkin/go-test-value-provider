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

func TestIntN(t *testing.T) {
  t.Run("successful", func(t *testing.T) {
    result := IntN(123)

    assert.LessOrEqual(t, result, 123, "Result should be less or equal than limited value")
    assert.GreaterOrEqual(t, result, math.MinInt, "Result should be greater or equal than minInt")
  })
}

func TestIntNM(t *testing.T) {
  t.Run("successful", func(t *testing.T) {
    result := IntNM(1,10)

    assert.LessOrEqual(t, result, 10, "Result should be less or equal than limited value")
    assert.GreaterOrEqual(t, result, 1, "Result should be greater or equal than minInt")
  })
}


func TestFloat32(t *testing.T) {
  result := Float32()

  assert.GreaterOrEqual(t, result, float32(0.0), "Result should be greater or equal than min float")
  assert.LessOrEqual(t, result, float32(math.MaxFloat32), "Result should be less or equal than maxFloat32")
}

func TestFloat32N(t *testing.T) {
  t.Run("successful", func(t *testing.T) {
    result := Float32N(float32(123.0))

    assert.LessOrEqual(t, result, float32(123.0), "Result should be less or equal than limited value")
    assert.GreaterOrEqual(t, result, float32(0.0), "Result should be greater or equal than min float")
  })
}

func TestFloat64(t *testing.T) {
  result := Float64()

  assert.GreaterOrEqual(t, result, float64(0.0), "Result should be greater or equal than min float")
  assert.LessOrEqual(t, result, math.MaxFloat64, "Result should be less or equal than maxFloat64")
}

func TestFloat64N(t *testing.T) {
  t.Run("successful", func(t *testing.T) {
    result := Float64N(float64(123.0))

    assert.LessOrEqual(t, result, float64(123.0), "Result should be less or equal than limited value")
    assert.GreaterOrEqual(t, result, float64(0.0), "Result should be greater or equal than minInt")
  })
}

func TestString(t *testing.T) {
  t.Run("successful String", func(t *testing.T) {
    result := String()

    assert.Equal(t, len(result), 10, "Result should be less or equal than limited value")
  })

  t.Run("successful StringN", func(t *testing.T) {
    result := StringN(12)
    
    assert.Equal(t, len(result), 12, "Result should be less or equal than limited value")
  })
}

func TestByteArray(t *testing.T) {
  t.Run("successful byteArray", func(t *testing.T) {
    result := ByteArray()

    assert.LessOrEqual(t, len(result), 100, "Result should be less than 100")
    assert.GreaterOrEqual(t, len(result), 1, "Result should be greated than 1")
  })
}

func TestByteArrayN(t *testing.T) {
  t.Run("successful byteArrayN", func(t *testing.T) {
    result := ByteArrayN(123)

    assert.LessOrEqual(t, len(result), 123, "Result should be less than 100")
    assert.GreaterOrEqual(t, len(result), 1, "Result should be greated than 1")
  })
}

