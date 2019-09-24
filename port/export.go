package port

import (
	"bufio"
	"encoding/json"
	"os"
)

func Export(fileName string, objects []interface{}) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, object := range objects {
		bs, err := json.Marshal(object)
		if err != nil {
			return err
		}

		_, err = writer.Write(bs)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
