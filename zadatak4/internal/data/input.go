package data

import (
	"encoding/json"
	"fmt"
	"os"

	"zadatak4.mihailoivic/internal/validator"
)

type Input struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Telephone string `json:"telephone"`
}

func ValidateInput(v *validator.Validator, input *Input) {
	v.Check(input.FirstName != "", "firstName", "must be provided")
	v.Check(input.LastName != "", "lastName", "must be provided")
	v.Check(input.Telephone != "", "telephone", "must be provided")
}

func Create(input *Input) error {
	allInputs, err := GetAll()
	if err != nil {
		return fmt.Errorf("error while getting the inputs")
	}

	if (len(allInputs)) > 0 {
		input.ID = (allInputs)[len(allInputs)-1].ID + 1
	} else {
		input.ID = 2
	}

	for _, i := range allInputs {
		if i.Telephone == input.Telephone {
			return fmt.Errorf("telephone already exists")
		}
	}

	allInputs = append(allInputs, input)

	file, err := os.Create("./internal/data/inputs.json")
	if err != nil {
		return fmt.Errorf("error while opening the file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(allInputs)
	if err != nil {
		return fmt.Errorf("error while encoding")
	}

	return err
}

func Get(id int64) (*Input, error) {
	inputs, err := GetAll()
	if err != nil {
		return nil, fmt.Errorf("error while getting the inputs")
	}

	for _, i := range inputs {
		if i.ID == id {
			return i, err
		}
	}

	return nil, fmt.Errorf("object not found")
}

func GetAll() ([]*Input, error) {
	file, err := os.Open("./internal/data/inputs.json")
	if err != nil {
		return nil, fmt.Errorf("error while opening the file")
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	var inputs []*Input

	err = decoder.Decode(&inputs)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the file")
	}

	return inputs, err

}

func Delete(id int64) error {
	inputs, err := GetAll()
	if err != nil {
		return fmt.Errorf("error while getting the inputs")
	}

	for idx, input := range inputs {
		if input.ID == id {
			inputs = append(inputs[0:idx], inputs[idx+1:]...)
		}
	}

	file, err := os.Create("./internal/data/inputs.json")
	if err != nil {
		return fmt.Errorf("error while opening the file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(inputs)
	if err != nil {
		return fmt.Errorf("error while encoding")
	}

	return err
}
