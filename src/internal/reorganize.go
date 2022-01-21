package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func MoveFiles(filenames []string, to string) {
	for _, filename := range filenames {
		original_filename := filename
		_, filename := filepath.Split(filename)

		err := os.Rename(original_filename, to+"/"+filename)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func GetArchiveNumber(in string) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
