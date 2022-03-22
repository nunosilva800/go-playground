package manipulator

import (
	"errors"
	"fmt"
)

func NewManipulator(txt string) *manipulator {
	return &manipulator{text: txt}
}

type manipulator struct {
	text string
}

func (m *manipulator) Command(cmd string) (string, int, error) {
	output := m.text
	pos := 0
	maxPos := len(m.text) - 1

	for idx := 0; idx < len(cmd); idx++ {
		char := cmd[idx]

		switch char {
		case 'h':
			pos = m.moveLeft(pos)

		case 'l':
			pos = m.moveRight(pos, maxPos)

		case 'r':
			if idx+1 >= len(cmd) {
				return output, pos, errors.New("missing char after 'r'")
			}
			charToReplaceWith := string(cmd[idx+1])
			output = m.replace(charToReplaceWith, output, pos)
			idx++

		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// scan until the next non digit char
			numToRepeat := 123
			// execute X amount of times the next non digit char
			for i := 0; i < numToRepeat; i++ {
				m.moveLeft(pos)
			}

		default:
			return output, pos, fmt.Errorf("command '%v' is not supported", string(char))
		}
	}

	return output, pos, nil
}

func (m *manipulator) moveLeft(currentPos int) int {
	currentPos--
	if currentPos <= 0 {
		currentPos = 0
	}
	return currentPos
}

func (m *manipulator) moveRight(currentPos, maxPos int) int {
	currentPos++
	if currentPos >= maxPos {
		currentPos = maxPos
	}
	return currentPos
}

func (m *manipulator) replace(charToReplaceWith, output string, pos int) string {
	pre := output[0:pos]
	after := output[pos+1:]
	return fmt.Sprintf("%v%v%v", pre, charToReplaceWith, after)
}
