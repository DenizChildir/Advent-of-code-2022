package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("coding_qual_input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	strArry := strings.Split(str2, "\n")
	a := len(strArry)
	b := len(strArry[0]) - 1
	//ar := [a][a]int{}
	ar := make([][]int, a)
	for i, _ := range ar {
		ar[i] = make([]int, b)
	}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			ar[i][j] = int(strArry[i][j] - '0')
		}
	}
	//fmt.Printf("%v", ar)
	ar2 := make([][]bool, a)
	for i, _ := range ar {
		ar2[i] = make([]bool, b)
	}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			ar2[i][j] = help(ar, i, j)
		}
	}
	//for i := 0; i < len(ar); i++ {
	//	for j := 0; j < len(ar[0]); j++ {
	//		if ar2[i][j] {
	//			fmt.Print("1 ")
	//		} else {
	//			fmt.Print("0 ")
	//		}
	//	}
	//	println("")
	//}
	count := 0
	for i, _ := range ar {
		for j, _ := range ar[i] {
			if ar2[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)

}
func help(ar [][]int, i, j int) bool {
	visibleLeft := true
	visibleRight := true
	visibleTop := true
	visibleDown := true

	// Check visibility to the left
	for x := 0; x < i; x++ {
		if ar[x][j] >= ar[i][j] {
			visibleTop = false
			break
		}
	}

	// Check visibility to the right
	for x := i + 1; x < len(ar); x++ {
		if ar[x][j] >= ar[i][j] {
			visibleDown = false
			break
		}
	}

	// Check visibility upwards
	for y := 0; y < j; y++ {
		if ar[i][y] >= ar[i][j] {
			visibleLeft = false
			break
		}
	}

	// Check visibility downwards
	for y := j + 1; y < len(ar[0]); y++ {
		if ar[i][y] >= ar[i][j] {
			visibleRight = false
			break
		}
	}

	return visibleLeft || visibleRight || visibleTop || visibleDown
}
