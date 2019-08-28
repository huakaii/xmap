package utils

import (
	"fmt"
	"io/ioutil"
)

func FileRead(path string) []byte {

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("File read error !")
	}
	return data
}
