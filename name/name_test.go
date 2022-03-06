package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestNameCheckerInputBelowMinimumLengthExpectInvalid(t *testing.T) {
	assert.Equal(t, nameChecker(""), "invalid")
}

func TestNameCheckerInputMinimumLengthExpectValid(t *testing.T) {
	assert.Equal(t, nameChecker("T"), "valid")
}

func TestNameCheckerInputAboveMinimumLengthExpectValid(t *testing.T) {
	assert.Equal(t, nameChecker("Tu"), "valid")
}

func TestNameCheckerInputNominalLengthExpectValid(t *testing.T) {
	assert.Equal(t, nameChecker("Tom"), "valid")
}

func TestNameCheckerInputBelowMaxLengthExpectValid(t *testing.T) {
	assert.Equal(t, nameChecker("Underwood"), "valid")
}

func TestNameCheckerInputMaxLengthExpectValid(t *testing.T) {
	assert.Equal(t, nameChecker("Rosalind M"), "valid")
}

func TestNameCheckerInputAboveMaxLengthExpectInvalid(t *testing.T) {
	assert.Equal(t, nameChecker("Rosalind MC"), "invalid")
}
