package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main6() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str := string(xxx)
	for i, _ := range str {
		mp := map[string]int{}
		for j := 0; j < 14; j++ {
			x := string(str[i+j])
			mp[x]++
		}
		if len(mp) == 14 {
			fmt.Println(i + 14)
			break
		}
	}
}
