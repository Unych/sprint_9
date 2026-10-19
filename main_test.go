package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 7, 8, 1000} {
		data := generateRandomElements(size)
		wantSize := size
		if wantSize < 0 {
			wantSize = 0
		}
		if len(data) != wantSize {
			t.Fatalf("size %d: got length %d, want %d", size, len(data), wantSize)
		}
		for _, value := range data {
			if value <= 0 {
				t.Fatalf("size %d: expected a positive value, got %d", size, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	testMaximum(t, maximum)
}

func TestMaxChunks(t *testing.T) {
	testMaximum(t, maxChunks)
}

func testMaximum(t *testing.T, findMaximum func([]int) int) {
	t.Helper()
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{42}, 42},
		{"single negative", []int{-42}, -42},
		{"zeros", []int{0, 0, 0}, 0},
		{"equal", []int{5, 5, 5}, 5},
		{"first maximum", []int{9, 2, 1}, 9},
		{"middle maximum", []int{1, 9, 2}, 9},
		{"last maximum", []int{1, 2, 9}, 9},
		{"negative", []int{-5, -2, -9}, -2},
		{"eight elements", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"negative chunks", []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMaximum(tt.data)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksPartitions(t *testing.T) {
	for size := 1; size <= 2*CHUNKS+1; size++ {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			data[position] = 100
			if got := maxChunks(data); got != 100 {
				t.Fatalf("size %d, position %d: got %d, want 100", size, position, got)
			}
		}
	}
}
