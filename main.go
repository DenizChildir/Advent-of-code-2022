package main

import (
	"io/ioutil"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
}
