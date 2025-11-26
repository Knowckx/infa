package infa

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 1, expected: 1},
		{input: -1, expected: 1},
		{input: 0, expected: 0},
		{input: 100, expected: 100},
		{input: -100, expected: 100},
	}

	for _, test := range tests {
		actual := Abs(test.input)
		if actual != test.expected {
			t.Errorf("Abs(%d): expected %d, got %d", test.input, test.expected, actual)
		}
	}
}

func TestAverager(t *testing.T) {
	// Test case 1: Add integers
	a := Averager{}
	a.AddNumber(10)
	a.AddNumber(20)
	a.AddNumber(30)

	if a.count != 3 {
		t.Errorf("Expected count to be 3, got %d", a.count)
	}
	if a.sum != 60.0 {
		t.Errorf("Expected sum to be 60.0, got %f", a.sum)
	}
	if a.Avg != 20.0 {
		t.Errorf("Expected average to be 20.0, got %f", a.Avg)
	}

	// Test case 2: Add float64
	a = Averager{}
	a.AddNumber(10.5)
	a.AddNumber(20.5)
	a.AddNumber(30.0)

	if a.count != 3 {
		t.Errorf("Expected count to be 3, got %d", a.count)
	}
	if a.sum != 61.0 {
		t.Errorf("Expected sum to be 61.0, got %f", a.sum)
	}
	if a.Avg != 20.333333333333332 { // Expected value for (10.5 + 20.5 + 30.0) / 3
		t.Errorf("Expected average to be 20.333333333333332, got %f", a.Avg)
	}

	// Test case 3: Mix integers and float64 (should convert int to float)
	a = Averager{}
	a.AddNumber(10)
	a.AddNumber(20.0)

	if a.count != 2 {
		t.Errorf("Expected count to be 2, got %d", a.count)
	}
	if a.sum != 30.0 {
		t.Errorf("Expected sum to be 30.0, got %f", a.sum)
	}
	if a.Avg != 15.0 {
		t.Errorf("Expected average to be 15.0, got %f", a.Avg)
	}
}
