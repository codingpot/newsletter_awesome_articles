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
	parent_dest, _ := filepath.Abs("../../archive")

	dest := MoveFiles(filenames, parent_dest)

	files, _ := ioutil.ReadDir(dest)
	assert.Equal(t, 5, len(files), "not enough files")

	os.RemoveAll(dest)

	_, err := os.Stat(dest)
	assert.Equal(t, true, os.IsNotExist(err), "dest directory still exist")
}

func TestSequenceNumbering(t *testing.T) {
	// expect the archive directory exists
	dest, _ := filepath.Abs("../../archive")
	archiveNumber := GetSequenceNumberFromDirs(dest)
	assert.Equal(t, 1, archiveNumber)

	os.Mkdir(dest+"/1", 0700)
	archiveNumber = GetSequenceNumberFromDirs(dest)
	assert.Equal(t, 2, archiveNumber)

	os.Mkdir(dest+"/2", 0700)
	archiveNumber = GetSequenceNumberFromDirs(dest)
	assert.Equal(t, 3, archiveNumber)

	os.Mkdir(dest+"/shit", 0700)
	archiveNumber = GetSequenceNumberFromDirs(dest)
	assert.Equal(t, 3, archiveNumber)

	os.RemoveAll(dest)
	os.Mkdir(dest, 0700)
}
