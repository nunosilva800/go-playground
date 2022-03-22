package parser_test

import (
	"testing"

	"github.com/nunosilva800/cron-parser-go/parser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParser_InitErrors(t *testing.T) {
	scenarios := []struct {
		name     string
		cronLine string
		errorStr string
	}{
		{
			name:     "empty line",
			cronLine: "",
			errorStr: "invalid syntax: empty line",
		},
		{
			name:     "invalid number of fields (too many)",
			cronLine: "* * * * * * cmd",
			errorStr: "invalid syntax: invalid number of fields",
		},
		{
			name:     "invalid number of fields (too few)",
			cronLine: "* * * * cmd",
			errorStr: "invalid syntax: invalid number of fields",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			prsr, err := parser.New(scenario.cronLine)
			assert.Nil(t, prsr)
			assert.ErrorIs(t, err, parser.ErrInvalidSyntax)
			assert.EqualError(t, err, scenario.errorStr)
		})
	}
}

func TestParser_ParseFieldInBounds(t *testing.T) {
	t.Run("not in bounds", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("-1", parser.DayOfWeek)
		assert.Empty(t, res)
		assert.EqualError(t, err, "'-1' is not within 0..9")

		res, err = parser.ParseFieldInBounds("10", parser.DayOfWeek)
		assert.Empty(t, res)
		assert.EqualError(t, err, "'10' is not within 0..9")
	})

	t.Run("asterisk", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("*", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7}, res)
	})

	t.Run("asterisk over N", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("*/2", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{2, 4, 6}, res)

		t.Run("errors", func(t *testing.T) {
			res, err := parser.ParseFieldInBounds("*/20", parser.DayOfWeek)
			assert.Empty(t, res)
			assert.EqualError(t, err, "'20' is not within 0..9")
		})
	})

	t.Run("set of values", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("1,2,3", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3}, res)
	})

	t.Run("range of values", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("1-5", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, res)

		t.Run("errors", func(t *testing.T) {
			res, err = parser.ParseFieldInBounds("1-3-5", parser.DayOfWeek)
			assert.Empty(t, res)
			assert.EqualError(t, err, "'1-3-5' is an invalid range")
		})
	})

	t.Run("range of values - wrapping around", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("4-2", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{4, 5, 6, 7, 1, 2}, res)

		res, err = parser.ParseFieldInBounds("15-5", parser.Hour)
		require.NoError(t, err)
		assert.Equal(t, []int{15, 16, 17, 18, 19, 20, 21, 22, 23, 0, 1, 2, 3, 4, 5}, res)
	})

	t.Run("single value", func(t *testing.T) {
		res, err := parser.ParseFieldInBounds("1", parser.DayOfWeek)
		require.NoError(t, err)
		assert.Equal(t, []int{1}, res)
	})
}

func TestParser_PrettyPrint(t *testing.T) {
	t.Run("with the fixed amount of fields", func(t *testing.T) {
		prsr, err := parser.New("*/15 0 1,15 * 1-5 /usr/bin/find")
		require.NoError(t, err)

		res := prsr.PrettyPrint()
		require.Len(t, res, 6)

		assert.Equal(t, "minute         0 15 30 45", res[0])
		assert.Equal(t, "hour           0", res[1])
		assert.Equal(t, "day of month   1 15", res[2])
		assert.Equal(t, "month          1 2 3 4 5 6 7 8 9 10 11 12", res[3])
		assert.Equal(t, "day of week    1 2 3 4 5", res[4])
		assert.Equal(t, "command        /usr/bin/find", res[5])
	})

	t.Run("with the yearly field", func(t *testing.T) {
		prsr, err := parser.New("@yearly /usr/bin/find")
		require.NoError(t, err)

		res := prsr.PrettyPrint()
		require.Len(t, res, 6)

		assert.Equal(t, "minute         0", res[0])
		assert.Equal(t, "hour           0", res[1])
		assert.Equal(t, "day of month   1", res[2])
		assert.Equal(t, "month          1", res[3])
		assert.Equal(t, "day of week    1 2 3 4 5 6 7", res[4])
		assert.Equal(t, "command        /usr/bin/find", res[5])
	})
}

func TestParser_PrintMinutes(t *testing.T) {
	scenarios := []struct {
		name   string
		line   string
		output string
	}{
		{
			name:   "minutes: asterisk",
			line:   "* * * * * cmd",
			output: "0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52 53 54 55 56 57 58 59",
		},
		{
			name:   "minutes: asterisk over N",
			line:   "*/10 * * * * cmd",
			output: "0 10 20 30 40 50",
		},
		{
			name:   "minutes: set of values",
			line:   "1,2,3 * * * * cmd",
			output: "1 2 3",
		},
		{
			name:   "minutes: range of values",
			line:   "1-5 * * * * cmd",
			output: "1 2 3 4 5",
		},
		{
			name:   "minutes: single value",
			line:   "1 * * * * cmd",
			output: "1",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			prsr, err := parser.New(scenario.line)
			require.NoError(t, err)

			res := prsr.PrintMinutes()
			assert.Equal(t, scenario.output, res)
		})
	}
}
