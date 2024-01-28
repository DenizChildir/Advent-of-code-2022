package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	val   int
	left  *node
	right *node
	top   *node
	dow   *node
	vis   bool
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	strArry := strings.Split(str2, "\n")
	ar := [5][5]int{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ar[i][j] = int(strArry[i][j] - '0')
		}
	}
	ar2 := [5][5]bool{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			ar2[i][j] = help(ar, i, j)
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(help(ar, 2, 3))
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			//if ar2[i][j] {
			//	fmt.Print("1 ")
			//} else {
			//	fmt.Print("0 ")
			//}
		}
		//println("")
	}

}
func help(ar [5][5]int, i, j int) bool {
	visl := true
	visr := true
	vist := true
	visd := true
	for x := 0; x < 5; x++ {
		fmt.Print(ar[i][x], " ")
	}
	fmt.Println()
	for x := 0; x < i; x++ {
		if ar[x][j] >= ar[i][j] {
			visl = false
		}
	}
	for x := i; x < 5; x++ {
		//fmt.Print(x, " ", ar[x][j])
		if ar[x][j] >= ar[i][j] {
			//fmt.Print(" I ran")
			visr = false
		}
		//println()
	}
	for x := 0; x < j; x++ {
		if ar[i][x] >= ar[i][j] {
			vist = false
		}
	}
	for x := j; x < 5; x++ {
		if ar[i][x] >= ar[i][j] {
			visd = false
		}
	}
	return visr || visl || vist || visd

}
