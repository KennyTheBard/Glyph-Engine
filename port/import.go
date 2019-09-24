package port

import (
	"bufio"
	"os"
)

func Import(fileName string, buflen int, unmarshalFunc func([]byte)) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, buflen)

	start := 0
	end := 0
	count := 0
	var currObject string

	for n := 1; n > 0; n, err = reader.Read(buf) {
		if err != nil {
			return err
		}

		// obtaint the current json object
		for i, b := range buf[:n] {
			if b == '{' {
				if start <= end && count == 0 {
					start = i
				}

				count += 1
			}

			if b == '}' {
				count -= 1
				if end <= start {
					end = i
				}

				if count == 0 {
					currObject += string(buf[start : end+1])
					unmarshalFunc([]byte(currObject))
					currObject = ""
					start = end
				}
			}
		}

		if end < start {
			currObject += string(buf[start:])
		} else {
			currObject = ""
		}

		start = 0
		end = 0
	}

	return nil
}
