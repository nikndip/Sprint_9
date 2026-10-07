package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "positive size", size: 100},
		{name: "zero size", size: 0},
		{name: "negative size", size: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			expectedLength := test.size
			if expectedLength < 0 {
				expectedLength = 0
			}

			if len(data) != expectedLength {
				t.Fatalf("generateRandomElements(%d) returned %d elements; want %d", test.size, len(data), expectedLength)
			}

			for index, value := range data {
				if value <= 0 {
					t.Errorf("element %d is not positive: %d", index, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "several elements", data: []int{3, 15, 7, 9}, want: 15},
		{name: "maximum repeated", data: []int{10, 3, 10}, want: 10},
		{name: "negative elements", data: []int{-7, -2, -10}, want: -2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d; want %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "fewer elements than chunks", data: []int{5, 11, 3}, want: 11},
		{name: "few negative elements", data: []int{-5, -2, -9}, want: -2},
		{name: "exactly eight elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "length with remainder", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100, 9}, want: 100},
		{name: "negative elements", data: []int{-8, -3, -12, -1, -9, -4, -6, -2}, want: -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks(%v) = %d; want %d", test.data, got, test.want)
			}
		})
	}
}
