package utils

import (
	"encoding/json"
	"log"
)

func JsonToStruct(model interface{}, modelJson string) error {
	err := json.Unmarshal([]byte(modelJson), model)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
