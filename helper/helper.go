package helper

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadInput takes a file name and splits it into lines based on provided separator
func ReadInput(fileName, separator string) []string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err.Error())
		os.Exit(1)
	}

	lines := strings.Split(string(bytes), separator)

	return lines
}

// Convert a slice of strings to ints, no error handling bc
func StringtoIntSlice(stringList []string) []int {
	var intSlice = make([]int, len(stringList))

	for i := 0; i < len(stringList); i++ {
		intSlice[i], _ = strconv.Atoi(stringList[i])
	}

	return intSlice
}
