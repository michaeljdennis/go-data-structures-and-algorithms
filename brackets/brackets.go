package brackets

import (
	"log"
)

var m map[string]string = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func IsBalanced(str string) bool {
	var st stack

	for _, r := range str {
		s := string(r)

		// If opening bracket, push onto stack and continue
		if isOpeningBracket(s) {
			st.push(s)
			continue
		}

		// Pop last value from the stack
		poppedBracket, ok := st.pop()
		if !ok {
			log.Fatalln("unable to pop bracket")
		}

		// Get opening bracket for current bracket in loop
		matchedOpeningBracket, ok := m[s]
		if !ok {
			log.Fatalf("character detected that is not in list of valid brackets: %s", m[s])
		}
		// Check if popped value is opening bracket for
		// current closing bracket
		if poppedBracket != matchedOpeningBracket {
			return false
		}
	}

	return true
}

func isOpeningBracket(str string) bool {
	if str == "(" || str == "[" || str == "{" {
		return true
	}
	return false
}

type stack struct {
	data []string
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *stack) pop() (string, bool) {
	if len(s.data) < 1 {
		return "", false
	}

	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return v, true
}
