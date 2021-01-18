package utils

import (
	"io/ioutil"
	"log"
)

func GetData(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
