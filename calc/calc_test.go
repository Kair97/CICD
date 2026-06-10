package calc

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 99},
		{"negative numbers", -2, -3, -5},
		{"with zero", 7, 0, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	got, err := Div(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != 5 {
		t.Errorf("Div(10, 2) = %v, want 5", got)
	}
}

func TestDivByZero(t *testing.T) {
	_, err := Div(1, 0)
	if !errors.Is(err, ErrDivByZero) {
		t.Errorf("expected ErrDivByZero, got %v", err)
	}
}
