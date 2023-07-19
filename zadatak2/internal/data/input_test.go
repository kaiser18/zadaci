package data

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  []int64
	}{
		{
			name:  "NoDuplicates",
			input: []int64{1, 2, 3},
			want:  []int64{1, 2, 3},
		},
		{
			name:  "Empty",
			input: []int64{},
			want:  []int64{},
		},
		{
			name:  "Duplicates",
			input: []int64{1, 1, 2, 2, 3},
			want:  []int64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Deduplicate(tt.input)

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("got %v, want %v", data, tt.want)
			}
		})
	}
}

func TestGetPairs(t *testing.T) {
	tests := []struct {
		name  string
		input []int64
		want  map[int64]int
	}{
		{
			name:  "NoPairs",
			input: []int64{1, 2, 3},
			want:  map[int64]int{},
		},
		{
			name:  "Empty",
			input: []int64{},
			want:  map[int64]int{},
		},
		{
			name:  "Pairs",
			input: []int64{1, 1, 2, 2, 3},
			want:  map[int64]int{1: 2, 2: 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := GetPairs(tt.input)

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("got %v, want %v", data, tt.want)
			}
		})
	}
}
