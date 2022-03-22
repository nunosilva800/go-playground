package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidSyntax = errors.New("invalid syntax")

type Parser struct {
	cronLine string

	minutes     []int
	hours       []int
	daysOfMonth []int
	months      []int
	daysOfWeek  []int
	command     string
}

// New instantiates a new parser for the provided cron line and attempts to parse it.
// If the cron line cannot be parsed, it returns an error.
func New(cronLine string) (*Parser, error) {
	p := Parser{cronLine: cronLine}
	err := p.parse()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Parser) parse() error {
	if p.cronLine == "" {
		return fmt.Errorf("%w: empty line", ErrInvalidSyntax)
	}

	if strings.Contains(p.cronLine, "@yearly") {
		parts := strings.Split(p.cronLine, " ")
		if len(parts) != 2 {
			return fmt.Errorf("%w: invalid number of fields", ErrInvalidSyntax)
		}
		// store fields in Parser
		p.minutes = []int{FieldLow(Minute)}
		p.hours = []int{FieldLow(Hour)}
		p.daysOfMonth = []int{FieldLow(DayOfMonth)}
		p.months = []int{FieldLow(Month)}
		p.daysOfWeek = []int{1, 2, 3, 4, 5, 6, 7}
		p.command = parts[1]
	} else {
		err := p.parseFixedNumberFields()
		if err != nil {
			return err
		}
	}

	return nil
}

type Field int

const (
	Minute Field = iota
	Hour
	DayOfMonth
	Month
	DayOfWeek
)

func FieldLow(f Field) int {
	switch f {
	case Minute:
		return 0
	case Hour:
		return 0
	case DayOfMonth:
		return 1
	case Month:
		return 1
	case DayOfWeek:
		return 1
	default:
		panic("look at this")
	}
}

func FieldHigh(f Field) int {
	switch f {
	case Minute:
		return 59
	case Hour:
		return 23
	case DayOfMonth:
		return 31
	case Month:
		return 12
	case DayOfWeek:
		return 7
	default:
		panic("look at this")
	}
}

func FieldDefault(f Field) []int {
	switch f {
	case Minute:
		return FieldLow(f)
	case Hour:
		return FieldLow(f)
	case DayOfMonth:
		return FieldLow(f)
	case Month:
		return FieldLow(f)
	case DayOfWeek:
		return []int
	default:
		panic("look at this")
	}
}

func (p *Parser) parseFixedNumberFields() error {
	parts := strings.Split(p.cronLine, " ")
	if len(parts) != 6 {
		return fmt.Errorf("%w: invalid number of fields", ErrInvalidSyntax)
	}

	minutes, err := ParseFieldInBounds(parts[0], Minute)
	if err != nil {
		return fmt.Errorf("%w: invalid 'minute' field: %v", ErrInvalidSyntax, err.Error())
	}

	hours, err := ParseFieldInBounds(parts[1], Hour)
	if err != nil {
		return fmt.Errorf("%w: invalid 'hour' field: %v", ErrInvalidSyntax, err.Error())
	}

	daysOfMonth, err := ParseFieldInBounds(parts[2], DayOfMonth)
	if err != nil {
		return fmt.Errorf("%w: invalid 'day of month' field: %v", ErrInvalidSyntax, err.Error())
	}

	months, err := ParseFieldInBounds(parts[3], Month)
	if err != nil {
		return fmt.Errorf("%w: invalid 'month' field: %v", ErrInvalidSyntax, err.Error())
	}

	daysOfWeek, err := ParseFieldInBounds(parts[4], DayOfWeek)
	if err != nil {
		return fmt.Errorf("%w: invalid 'day of week' field: %v", ErrInvalidSyntax, err.Error())
	}

	// store fields in Parser
	p.minutes = minutes
	p.hours = hours
	p.daysOfMonth = daysOfMonth
	p.months = months
	p.daysOfWeek = daysOfWeek
	p.command = parts[5]

	return nil
}

// Possible input shapes:
// asterisk:        *    -> 0,1,2,3,4,5...59
// asterisk over N: */15 -> 0 15 30 45
// set of values:   1,15 -> 1 15
// range of values: 1-5  -> 1 2 3 4 5
// single value:    42   -> 42
func ParseFieldInBounds(str string, field Field) ([]int, error) {
	var result []int

	low := FieldLow(field)
	high := FieldHigh(field)

	// asterisk
	if str == "*" {
		for i := low; i <= high; i++ {
			result = append(result, i)
		}
		return result, nil
	}

	// asterisk over N
	if strings.Contains(str, "/") {
		values := strings.Split(str, "/")
		if len(values) != 2 {
			return nil, fmt.Errorf("'%v' is invalid", str)
		}
		increment, err := fieldToIntInBounds(values[1], low, high)
		if err != nil {
			return nil, err
		}
		for i := low; i < high; i += increment {
			result = append(result, i)
		}
		return result, nil
	}

	// set of values
	if strings.Contains(str, ",") {
		values := strings.Split(str, ",")
		for _, val := range values {
			v, err := fieldToIntInBounds(val, low, high)
			if err != nil {
				return nil, err
			}
			result = append(result, v)
		}
		return result, nil
	}

	// range of values
	if strings.Contains(str, "-") {
		values := strings.Split(str, "-")
		if len(values) != 2 {
			return nil, fmt.Errorf("'%v' is an invalid range", str)
		}

		// check that this is not a single negative value
		// if that is the case, continue
		if values[0] != "" {
			from, err := fieldToIntInBounds(values[0], low, high)
			if err != nil {
				return nil, err
			}

			to, err := fieldToIntInBounds(values[1], low, high)
			if err != nil {
				return nil, err
			}

			if from < to {
				for i := from; i <= to; i++ {
					result = append(result, i)
				}
			} else {
				for i := from; i <= high; i++ {
					result = append(result, i)
				}
				for i := low; i <= to; i++ {
					result = append(result, i)
				}
			}

			return result, nil
		}
	}

	// single value
	val, err := fieldToIntInBounds(str, low, high)
	if err != nil {
		return nil, err
	}
	result = append(result, val)

	return result, nil
}

func fieldToIntInBounds(str string, low int, high int) (int, error) {
	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	if val < low || val > high {
		return 0, fmt.Errorf("'%v' is not within %d..%d", val, low, high)
	}
	return val, nil
}

// PrettyPrint builds a set of formatted strings where each cron field is expanded to showcase their
// respective times.
//
// Example, for the cron line "*/15 0 1,15 * 1-5 /usr/bin/find", generates:
//
// minute        0 15 30 45
// hour          0
// day of month  1 15
// month         1 2 3 4 5 6 7 8 9 10 11 12
// day of week   1 2 3 4 5
// command       /usr/bin/find
func (p *Parser) PrettyPrint() []string {
	result := make([]string, 6)

	label0 := "minute        "
	label1 := "hour          "
	label2 := "day of month  "
	label3 := "month         "
	label4 := "day of week   "
	label5 := "command       "

	result[0] = fmt.Sprintf("%v %v", label0, p.PrintMinutes())
	result[1] = fmt.Sprintf("%v %v", label1, p.PrintHours())
	result[2] = fmt.Sprintf("%v %v", label2, p.PrintDaysOfMonth())
	result[3] = fmt.Sprintf("%v %v", label3, p.PrintMonths())
	result[4] = fmt.Sprintf("%v %v", label4, p.PrintDaysOfWeek())
	result[5] = fmt.Sprintf("%v %v", label5, p.command)

	return result
}

func (p *Parser) PrintMinutes() string {
	return strings.Join(intSliceToStringSlice(p.minutes), " ")
}

func (p *Parser) PrintHours() string {
	return strings.Join(intSliceToStringSlice(p.hours), " ")
}

func (p *Parser) PrintDaysOfMonth() string {
	return strings.Join(intSliceToStringSlice(p.daysOfMonth), " ")
}

func (p *Parser) PrintMonths() string {
	return strings.Join(intSliceToStringSlice(p.months), " ")
}

func (p *Parser) PrintDaysOfWeek() string {
	return strings.Join(intSliceToStringSlice(p.daysOfWeek), " ")
}

func intSliceToStringSlice(in []int) []string {
	parts := []string{}
	for _, v := range in {
		parts = append(parts, fmt.Sprintf("%d", v))
	}
	return parts
}
