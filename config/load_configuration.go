package config

import (
	"encoding/json"

	port "../port"
)

func LoadConfig(fileName string) error {
	return port.Import(fileName, 100, func(bs []byte) {
		json.Unmarshal(bs, &GlobalConfiguration)
	})
}
