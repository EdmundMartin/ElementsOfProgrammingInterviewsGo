package ch8_stacks_queues

import (
	"testing"
)

func TestEvaluateRPN(t *testing.T) {
	exampleRPN := "1,1,+,-2,*"
	result := EvaluateRPN(exampleRPN)

	if result != -4 {
		t.Error("unexpected value")
	}
}
