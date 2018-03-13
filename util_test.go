package goxc_test

import "os"

func withTempDir(f func(path string), path string) {
	tempDir := os.TempDir() + "/" + path
	_ = os.Mkdir(tempDir, 0775)
	//defer os.RemoveAll(tempDir)
	f(tempDir)
}

func Namespaces() map[string]string {
	return map[string]string{
		"foo": "foo",
		"w3c": "w3c",
	}
}
