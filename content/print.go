package home

import (
	"io/ioutil"
	"log"
)

func Content() []byte {
	content, err := ioutil.ReadFile("result.txt")
	if err != nil {
		log.Fatal(err)
	}
	return content
}
