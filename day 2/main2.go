package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//type play struct {
//	play string
//	win   string
//	loose string
//	tie string
//	score int
//}
//var rock play{"X","Z","B","A",1}
//var paper play{"Y","A","Z","B",2}
//var siz play{"Z","Y","A","C",3}
//036   123
mp := map[string]int{
"AY": 4,
"AX": 8,
"AZ": 9,
"YA": 2,
"YB": 6,
"YZ":
}
func main2() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	str2 = strings.ReplaceAll(str2, "\r", "x")
	//str2 = strings.ReplaceAll(str2, "\n", "y")
	str2 = strings.ReplaceAll(str2, " ", "")
	str := strings.Split(str2, "\n")
	fmt.Printf("%v", str)
}
func rpc(str string) int {


}
