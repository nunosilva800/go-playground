package manipulator_test

import (
	manipulator "ex-shop-go/text-manipulator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManipulator(t *testing.T) {
	input := "Hello World"
	m := manipulator.NewManipulator(input)

	out, pos, err := m.Command("lll")
	assert.NoError(t, err)
	assert.Equal(t, 3, pos)
	assert.Equal(t, "Hello World", out)

	out, pos, err = m.Command("hhh")
	assert.NoError(t, err)
	assert.Equal(t, 0, pos)
	assert.Equal(t, "Hello World", out)

	out, pos, err = m.Command("hhlhllhlhhll")
	assert.NoError(t, err)
	assert.Equal(t, 2, pos)
	assert.Equal(t, "Hello World", out)

	t.Run("with command 'r'", func(t *testing.T) {
		_, _, err := m.Command("r")
		assert.EqualError(t, err, "missing char after 'r'")

		out, pos, err = m.Command("rz")
		assert.NoError(t, err)
		assert.Equal(t, 0, pos)
		assert.Equal(t, "zello World", out)

		out, pos, err = m.Command("lrz")
		assert.NoError(t, err)
		assert.Equal(t, 1, pos)
		assert.Equal(t, "Hzllo World", out)

		out, pos, err = m.Command("rhllllllrw")
		assert.NoError(t, err)
		assert.Equal(t, 6, pos)
		assert.Equal(t, "hello world", out)
	})
}

func TestManipulator_Repeats(t *testing.T) {
	input := "Hello World"
	m := manipulator.NewManipulator(input)

	out, pos, err := m.Command("3l")
	assert.NoError(t, err)
	assert.Equal(t, 3, pos)
	assert.Equal(t, "Hello World", out)
}
