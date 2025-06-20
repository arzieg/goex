package storage

import "os"

type Storage interface {
	Save() error
}

type FileStorage struct {
	Filename string
	Content  string
}

func (f *FileStorage) Save() error {
	return os.WriteFile(f.Filename, []byte(f.Content), 0644)
}
