package file

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type File struct {
	filename string
}

func NewFileStorage(filename string) *File {
	return &File{
		filename: filename,
	}
}

func (f File) Set(i interface{}) error {
	data, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(f.filename, data, os.ModePerm)
}

func (f File) Get(i interface{}) error {
	data, err := ioutil.ReadFile(f.filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}
