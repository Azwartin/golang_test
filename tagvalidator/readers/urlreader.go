package readers

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

//URLReader - get html data from url
type URLReader struct{}

func (reader *URLReader) Read(url string, params ...interface{}) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}

	contentType := strings.ToUpper(r.Header.Get("Content-Type"))
	expectedType := strings.ToUpper("text/html")

	if !strings.Contains(contentType, expectedType) {
		return "", errors.New("Incorrect page Content-Type")
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
