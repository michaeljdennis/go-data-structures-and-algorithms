package brackets

import "testing"

func TestIsBalanced(t *testing.T) {
	balanced := "([])[]({})"
	unbalanced := "([)]"

	if !IsBalanced(balanced) {
		t.Errorf("return value should be true for the following brackets: %s", balanced)
	}

	if IsBalanced(unbalanced) {
		t.Errorf("return value should be false for the following brackets: %s", unbalanced)
	}
}
