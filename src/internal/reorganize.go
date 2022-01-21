package internal

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func createDirIfNotExist(dest string) bool {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		os.Mkdir(dest, 0700)
		return true
	}

	return false
}

func MoveFiles(filenames []string, to string) string {
	createDirIfNotExist(to)
	seq_num := GetSequenceNumberFromDirs(to)
	target_to := to + "/" + strconv.Itoa(seq_num)
	createDirIfNotExist(target_to)

	for _, filename := range filenames {
		original_filename := filename
		_, filename := filepath.Split(filename)

		err := os.Rename(original_filename, target_to+"/"+filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	return target_to
}

func GetSequenceNumberFromDirs(in string) int {
	if createDirIfNotExist(in) {
		return -1
	}

	files, err := ioutil.ReadDir(in)
	if err != nil {
		log.Fatal(err)
	}

	count := 1
	for _, f := range files {
		if f.IsDir() {
			if _, err := strconv.Atoi(f.Name()); err == nil {
				count++
			}
		}
	}

	return count
}
