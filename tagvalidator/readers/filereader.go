package readers

import (
	"io/ioutil"
)

//FileReader - get html data from file
type FileReader struct{}

func (reader *FileReader) Read(filename string, params ...interface{}) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
