package pkg

import (
	"errors"
	"os"
	"path"
)

// MkdirIfNotExist 如果目录不存在，则创建目录
func MkdirIfNotExist(dir string) error {
	if dir == "" {
		return errors.New("dir is empty")
	}

	dir = path.Dir(dir)
	if dir == "." || dir == "/" {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.Mkdir(dir, 0755)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
