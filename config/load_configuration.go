package config

import (
	"encoding/json"

	port "../port"
)

func LoadConfig(fileName string) error {
	var err error
	if err = port.Import(fileName, 100, func(bs []byte) {
		if json.Unmarshal(bs, &GlobalConfiguration) != nil {
			GlobalConfiguration.StartingStoryID = 0
			GlobalConfiguration.Port = 8080
		}
	}); err != nil {
		GlobalConfiguration.StartingStoryID = 0
		GlobalConfiguration.Port = 8080
	}

	return err
}
