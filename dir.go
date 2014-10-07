package goutils

// create dir if it not exist

import (
	"os"
)

func CreateDirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	// 目录存在
	if err == nil {
		return nil
	}

	// 其他错误
	if !os.IsNotExist(err) {
		return err
	}
	return os.MkdirAll(dir, os.ModeDir|0755)
}
