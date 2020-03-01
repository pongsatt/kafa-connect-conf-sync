package file

import (
	"io/ioutil"
	"path"
	"strings"
)

// Reader read configs from files
type Reader struct {
	rootPath string
}

// NewFileReader creates new file reader instance
func NewFileReader(rootPath string) *Reader {
	return &Reader{
		rootPath: rootPath,
	}
}

// ListConfigs lists all configs file under rootPath
func (fileReader *Reader) Read() ([][]byte, error) {
	files, err := ioutil.ReadDir(fileReader.rootPath)

	if err != nil {
		return nil, err
	}

	results := make([][]byte, 0)

	for _, file := range files {
		filename := file.Name()

		if !strings.HasSuffix(filename, ".json") {
			continue
		}

		filePath := path.Join(fileReader.rootPath, filename)
		data, err := ioutil.ReadFile(filePath)

		if err != nil {
			return nil, err
		}

		results = append(results, data)
	}
	return results, nil
}
