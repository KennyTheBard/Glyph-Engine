package port

import (
	"bufio"
	"os"
)

func Export(fileName string, objects interface{})) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for object := range objects {
		bs, err := json.Marshal(object)
		if err != nil {
			return err
		}

		_, err := writer.Write(bs)
		if err != nil {
			return err
		}
	}

	writer.Flush()
	return nil
}
