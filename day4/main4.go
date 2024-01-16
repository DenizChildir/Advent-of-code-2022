package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

func main4() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	re := regexp.MustCompile(`(?s)([[:graph:]]+)`)
	str := re.FindAllString(str2, -1)
	//fmt.Printf("%q", str)
	sum := 0
	for _, i := range str {
		sum = sum + isover(i)
	}
	fmt.Println(sum)
}
func isover(str string) int {
	re := regexp.MustCompile(`(\d+)`)
	ar2 := re.FindAllString(str, -1)
	ar := make([]int, len(ar2))
	for i, j := range ar2 {
		ar[i], _ = strconv.Atoi(j)
	}
	if ar[0] < ar[2] && ar[1] < ar[2] {
		return 0
	}
	if ar[2] < ar[0] && ar[3] < ar[0] {
		return 0
	}
	return 1
}

//for the input ["2" "2" "19" "30"], why does the second if statement return true
