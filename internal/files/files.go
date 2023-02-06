package files

import "os"

func WriteToFile(value string, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}

	_, err = f.WriteString(value)
	return
}

func WriteBytesToFile(value []byte, path string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}

	_, err = f.Write(value)
	return
}
