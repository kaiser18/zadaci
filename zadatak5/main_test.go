package main

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "NoDuplicates",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Duplicates",
			input: []int{1, 1, 2, 2, 3},
			want:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := deduplicate(tt.input)

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("got %v, want %v", data, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "AlreadySorted",
			input: []int{1, 2, 3},
			want:  []int{1, 2, 3},
		},
		{
			name:  "Empty",
			input: []int{},
			want:  []int{},
		},
		{
			name:  "Unsorted",
			input: []int{2, 3, 1, 4, 2},
			want:  []int{1, 2, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := sort(tt.input)

			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("got %v, want %v", data, tt.want)
			}
		})
	}
}
