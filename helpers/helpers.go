package helpers

import (
	"fmt"
	"io/ioutil"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func ReadFile(targetFile string) string {
	data, err := ioutil.ReadFile(targetFile)
	errorHandler(err)
	return string(data)
}
