package workers

import (
	"encoding/json"
	"github.com/levor/seeBattle/internal/types"
	"io/ioutil"
)

const FILENAME = "data.json"

func GetData() ([]types.Subject, error) {
	file, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		return nil, err
	}
	subjects := make([]types.Subject, 0)
	if err = json.Unmarshal(file, &subjects); err != nil {
		return nil, err
	}
	return subjects, nil
}

func UpdateData(subjects []types.Subject) error {
	data, err := json.Marshal(subjects)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(FILENAME, data, 0); err != nil {
		return err
	}
	return nil
}
