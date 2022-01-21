package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func creatEmptyFile(path string) string {
	emptyFile, _ := os.Create(path)
	emptyFile.Close()
	path, _ = filepath.Abs(path)
	return path
}

func TestMoveFiles(t *testing.T) {
	path1 := creatEmptyFile("../../current/empty1.yaml")
	path2 := creatEmptyFile("../../current/empty2.yaml")
	path3 := creatEmptyFile("../../current/empty3.yaml")
	path4 := creatEmptyFile("../../current/empty4.yaml")
	path5 := creatEmptyFile("../../current/empty5.yaml")

	filenames := []string{path1, path2, path3, path4, path5}
	dest, _ := filepath.Abs("../../archive")

	MoveFiles(filenames, dest)

	files, _ := ioutil.ReadDir(dest)
	assert.Equal(t, 5, len(files), "not enough files")

	os.RemoveAll(dest)
	os.Mkdir(dest, 0700)

	files, _ = ioutil.ReadDir(dest)
	assert.Equal(t, 0, len(files), "not enough files")
}

func TestArchiveNumbering(t *testing.T) {
	// dest, _ := filepath.Abs("../../archive")

	// archiveNumber := GetArchiveNumber(dest)

	// path1 := creatEmptyFile("../../current/empty1.yaml")
	// path2 := creatEmptyFile("../../current/empty2.yaml")
	// path3 := creatEmptyFile("../../current/empty3.yaml")
	// path4 := creatEmptyFile("../../current/empty4.yaml")
	// path5 := creatEmptyFile("../../current/empty5.yaml")

	// filenames := []string{path1, path2, path3, path4, path5}
	// dest, _ := filepath.Abs("../../archive")

	// MoveFiles(filenames, dest)
}
