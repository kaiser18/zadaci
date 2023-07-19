package data

import (
	"fmt"

	"github.com/google/uuid"
	"zadatak2.mihailoivic/internal/validator"
)

type Input struct {
	Operation string  `json:"operation"`
	Data      []int64 `json:"data"`
}

type Output struct {
	ID        string      `json:"id"`
	Operation string      `json:"operation"`
	Data      interface{} `json:"data"`
}

func ValidateInput(v *validator.Validator, input *Input) {
	v.Check(input.Operation != "", "operation", "must be provided")
	v.Check(input.Data != nil, "data", "must be provided")
}

func DoTheOperation(input *Input) (*Output, error) {
	output := Output{
		ID:        uuid.NewString(),
		Operation: input.Operation,
		Data:      make([]int64, 0),
	}

	if input.Operation == "deduplicate" {
		result := Deduplicate(input.Data)
		output.Data = result

		return &output, nil
	} else if input.Operation == "getPairs" {
		result := GetPairs(input.Data)
		output.Data = result

		return &output, nil
	}

	return nil, fmt.Errorf("there was an error")
}

func Deduplicate(data []int64) []int64 {
	allKeys := make(map[int64]bool)
	list := []int64{}

	for _, item := range data {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func GetPairs(data []int64) map[int64]int {
	dict := make(map[int64]int)
	for _, num := range data {
		dict[num] = dict[num] + 1
	}

	for key, val := range dict {
		if val == 1 {
			delete(dict, key)
		}
	}

	return dict
}
