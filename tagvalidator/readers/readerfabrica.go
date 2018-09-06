package readers

import "net/url"

//CreateReader creates Reader by path, if path is url - returns UrlReader, else - File
func CreateReader(path string) ReaderInterface {
	_, err := url.ParseRequestURI(path)

	if err != nil {
		return &FileReader{}
	}

	return &URLReader{}
}
