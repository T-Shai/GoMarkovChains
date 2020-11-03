package mchains

import (
	"io/ioutil"
	"log"
)

// ReadFile : read a file and return the content as string
func ReadFile(name string) string {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
