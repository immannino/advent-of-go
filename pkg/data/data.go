package data

import (
	"io/ioutil"
	"log"
)

func ReadAsString(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func ReadAsBytes(path string) []byte {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)

		return data
	}

	return data
}
