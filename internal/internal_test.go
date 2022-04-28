package internal

import "testing"

func TestInRange(t *testing.T) {
	tests := []struct {
		src  uint8
		min  uint8
		max  uint8
		want bool
	}{
		{15, 10, 20, true},
		{10, 10, 20, true},
		{20, 10, 20, true},
		{1, 10, 20, false},
	}
	for _, tt := range tests {
		if got := inRange(tt.src, tt.min, tt.max); got != tt.want {
			t.Errorf("want %v, got %v\n", tt.want, got)
		}
	}
}

func TestIsRegularLower(t *testing.T) {
	tests := []struct {
		src  []byte
		want bool
	}{
		{[]byte("a"), true},
		{[]byte("z"), true},
		{[]byte("A"), false},
		{[]byte("Z"), false},
		{[]byte("0"), false},
		{[]byte("9"), false},
		{[]byte("あ"), false},
		{[]byte(""), false},
	}

	for _, tt := range tests {
		if got := isRegularLower(tt.src); got != tt.want {
			t.Errorf("input: %s, got %v, want %v\n", string(tt.src), got, tt.want)
		}
	}
}

func TestIsRegularUpper(t *testing.T) {
	tests := []struct {
		src  []byte
		want bool
	}{
		{[]byte("A"), true},
		{[]byte("Z"), true},
		{[]byte("a"), false},
		{[]byte("z"), false},
		{[]byte("0"), false},
		{[]byte("9"), false},
		{[]byte("あ"), false},
		{[]byte(""), false},
	}

	for _, tt := range tests {
		if got := isRegularUpper(tt.src); got != tt.want {
			t.Errorf("input: %s, got %v, want %v\n", string(tt.src), got, tt.want)
		}
	}
}

func TestIsRegularDigit(t *testing.T) {
	tests := []struct {
		src  []byte
		want bool
	}{
		{[]byte("0"), true},
		{[]byte("9"), true},
		{[]byte("a"), false},
		{[]byte("z"), false},
		{[]byte("A"), false},
		{[]byte("Z"), false},
		{[]byte("あ"), false},
		{[]byte(""), false},
	}

	for _, tt := range tests {
		if got := isRegularDigit(tt.src); got != tt.want {
			t.Errorf("input: %s, got %v, want %v\n", string(tt.src), got, tt.want)
		}
	}
}
