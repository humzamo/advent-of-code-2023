package helpers

import (
	"bufio"
	"log"
	"os"
)

// LoadStringList loads the input as a slice of strings
func LoadStringList(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	return list
}
