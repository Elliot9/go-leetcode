package main

import (
	"testing"
)

func TestParseBoolExpr(t *testing.T) {
	testCases := []struct {
		name       string
		expression string
		expected   bool
	}{
		{
			name:       "Simple true",
			expression: "t",
			expected:   true,
		},
		{
			name:       "Simple false",
			expression: "f",
			expected:   false,
		},
		{
			name:       "Not operation",
			expression: "!(f)",
			expected:   true,
		},
		{
			name:       "And operation",
			expression: "&(t,f)",
			expected:   false,
		},
		{
			name:       "Or operation",
			expression: "|(t,f)",
			expected:   true,
		},
		{
			name:       "Complex expression 1",
			expression: "&(t,t,f)",
			expected:   false,
		},
		{
			name:       "Complex expression 2",
			expression: "|(&(t,f,t),!(t))",
			expected:   false,
		},
		{
			name:       "Nested expression",
			expression: "&(|(f,f,f,t),&(t,t,t))",
			expected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := parseBoolExpr(tc.expression)
			if result != tc.expected {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
