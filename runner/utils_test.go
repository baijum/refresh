package runner

import "testing"

func TestIsWatchedFile(t *testing.T) {
	validExtensions := []string{
		"test.go",
		"test.tpl",
		"test.tmpl",
		"test.html"}

	for _, fileName := range validExtensions {
		if isWatchedFile(fileName) != true {
			t.Error("File not watched:", fileName)
		}
	}

	invalidExtensions := []string{
		"test.css",
		"test-executable"}

	for _, fileName := range invalidExtensions {
		if isWatchedFile(fileName) != false {
			t.Error("File watched:", fileName)
		}
	}

	filesInTmp := []string{
		"./tmp/test.go",
		"./tmp/test.html"}

	for _, fileName := range filesInTmp {
		if isWatchedFile(fileName) != false {
			t.Error("File watched:", fileName)
		}
	}
}
