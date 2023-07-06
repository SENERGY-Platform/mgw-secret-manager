package files

import (
	"os"
	"path/filepath"
)

func WriteToFile(value string, path string) (err error) {
	if err := os.MkdirAll(filepath.Dir(path), 0770); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return
	}

	defer f.Close()
	_, err = f.WriteString(value)
	return
}

func WriteBytesToFile(value []byte, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = f.Write(value)
	return
}
