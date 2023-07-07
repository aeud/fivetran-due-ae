package main

import (
	"due/fivetranio"
	"strings"
	"testing"
)

func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestConvertVoidState(t *testing.T) {
	initialStateReader := strings.NewReader("{}")
	step1 := "step1"
	step2 := "step2"
	step3 := "step3"
	step4 := "step4"
	allSteps := []string{step1, step2, step3, step4}
	expectedCurrentStep := step1
	expectedRemainingSteps := []string{step2, step3, step4}
	expectedNextPageNumber := 1

	s, err := fivetranio.NewStateFromReader(initialStateReader, allSteps)
	if err != nil {
		t.Fatalf("failed when decoding the initial state: %s", err.Error())
		t.Fail()
	}
	if s.CurrentStep != expectedCurrentStep {
		t.Fatalf("should get %s as an expected current step. got %s", expectedCurrentStep, s.CurrentStep)
		t.Fail()
	}
	if !EqualStringSlices(s.RemainingSteps, expectedRemainingSteps) {
		t.Fatalf("should get %s as an expected remaining steps. got %s", expectedRemainingSteps, s.RemainingSteps)
		t.Fail()
	}
	if s.NextPageNumber != expectedNextPageNumber {
		t.Fatalf("should get %d as an expected next page number. got %d", expectedNextPageNumber, s.NextPageNumber)
		t.Fail()
	}
	// Check if there is no cursor
	if l := len(s.GetAllCursors()); l > 0 {
		t.Fatalf("should get no cursors. got %d", l)
		t.Fail()
	}
}
