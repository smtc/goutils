package goutils

import "testing"
import "os"

func TestCreateDir(t *testing.T) {
	err := CreateDirIfNotExist("./test")
	if err != nil {
		t.Fatal(err)
	}
	err = CreateDirIfNotExist("./test/a/b/c")
	if err != nil {
		t.Fatal(err)
	}
	os.RemoveAll("./test")
}
