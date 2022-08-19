package readingjsonexample

import (
	"encoding/json"
	"io/ioutil"
)

type DataType map[string]interface{}

func GetSomeJSON() DataType {
	var data = DataType{}
	var jsonPath = "readingjsonexample/some.json"

	content, err := ioutil.ReadFile(jsonPath)

	if err == nil {
		json.Unmarshal(content, &data)
	}

	return data

}
