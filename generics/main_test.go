package main

import (
	"testing"
)

// RETURN CORRECT SUM OF INT64 VALUE
func TestSumIntsOrFloatsInt64Values(t *testing.T) {
	// Initialize a map with int64 values
	m := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Call the SumIntsOrFloats function
	result := SumIntsOrFloats(m)

	// Assert that the result is correct
	expected := int64(46)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

// RETURN CORRECT SUM OF NEGATIVE MAP VALUES
func TestSumIntsOrFloatsNegativeValues(t *testing.T) {
	// Initialize a map with negative int64 values
	m := map[string]int64{
		"first":  -34,
		"second": -12,
	}

	// Call the SumIntsOrFloats function
	result := SumIntsOrFloats(m)

	// Assert that the result is correct
	expected := int64(-46)
	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
