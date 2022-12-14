package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInt(t *testing.T) {
	assert.Equal(t, 2, MinInt(2, 4))
	assert.Equal(t, 3, MinInt(3, 3))
}

func TestMaxInt(t *testing.T) {
	assert.Equal(t, 4, MaxInt(2, 4))
	assert.Equal(t, 3, MaxInt(3, 3))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 2, Abs(2))
	assert.Equal(t, 0, Abs(0))
	assert.Equal(t, 5, Abs(-5))
}
func TestPosMod(t *testing.T) {
	assert.Equal(t, 2, PosMod(2, 4))
	assert.Equal(t, 0, PosMod(4, 4))
	assert.Equal(t, 1, PosMod(5, 4))
	assert.Equal(t, 3, PosMod(-1, 4))
	assert.Equal(t, 2, PosMod(-6, 4))
}

func TestClampToOne(t *testing.T) {
	assert.Equal(t, 1, ClampToOne(15))
	assert.Equal(t, 0, ClampToOne(0))
	assert.Equal(t, -1, ClampToOne(-7))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, 6, Lcm(2, 3))
	assert.Equal(t, 12, Lcm(3, 12))
	assert.Equal(t, 12, Lcm(4, 6))
	assert.Equal(t, 56, Lcm(7, 8))
}

func TestSumUpTo(t *testing.T) {
	assert.Equal(t, 0, SumUpTo(0))
	assert.Equal(t, 1, SumUpTo(1))
	assert.Equal(t, 10, SumUpTo(4))
	assert.Equal(t, 28, SumUpTo(7))
}

func TestSerializeCoord(t *testing.T) {
	assert.Equal(t, "1 2", SerializeCoordRaw(1, 2))
}

func TestDeserializeCoord(t *testing.T) {
	x, y := DeserializeCoordRaw("1 2")
	assert.Equal(t, 1, x)
	assert.Equal(t, 2, y)
}

func TestSerializeCoord3(t *testing.T) {
	assert.Equal(t, "1 2 4", SerializeCoord3(Coord3{1, 2, 4}))
}

func TestDeserializeCoord3(t *testing.T) {
	coord3 := DeserializeCoord3("1 2 4")
	assert.Equal(t, 1, coord3.X)
	assert.Equal(t, 2, coord3.Y)
	assert.Equal(t, 4, coord3.Z)
}
