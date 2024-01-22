package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main5() {

	ar := [][]string{
		{"S", "M", "R", "N", "W", "J", "V", "T"},
		{"B", "W", "D", "J", "Q", "P", "C", "V"},
		{"B", "J", "F", "H", "D", "R", "P"},
		{"F", "R", "P", "B", "M", "N", "D"},
		{"H", "V", "R", "P", "T", "B"},
		{"C", "B", "P", "T"},
		{"B", "J", "R", "P", "L"},
		{"N", "C", "S", "L", "T", "Z", "B", "W"},
		{"L", "S", "G"}}
	r := regexp.MustCompile(`(?s)(\d+)`)
	input, _ := os.Open("input.txt.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	mv := r.FindAllString(string(xxx), -1)
	mv2 := make([][]int, len(mv)/3)
	for i := 0; i < len(mv2); i++ {
		mv2[i] = make([]int, 3)
	}
	counter := 0
	for i := 0; i < len(mv2); i++ {
		for j := 0; j < 3; j++ {
			mv2[i][j], _ = strconv.Atoi(mv[counter])
			counter++
		}
	}

	_, _ = ar, mv
	//fmt.Println(mv2)
	for i := 0; i < len(mv2); i++ {
		ar = move(ar, mv2[i][0], mv2[i][1]-1, mv2[i][2]-1)
	}
	fmt.Println(ar)
}
func parse() ([][]string, int) {
	input2, _ := os.Open("input2.txt")
	defer input2.Close()
	xxx, _ := ioutil.ReadAll(input2)
	str2 := string(xxx)
	curStack := strings.Split(str2, "\n")
	re := regexp.MustCompile(`(\d+)`)
	max2 := re.FindAllString(curStack[len(curStack)-1], -1)
	max3 := max2[len(max2)-1]
	str := strings.Split(str2, "\n")
	max4, _ := strconv.Atoi(max3)
	ar := make([][]string, max4-1)
	for i := 0; i < len(str)-1; i++ {
		temp := make([]string, len(str[i]))
		ar[i] = temp
		for j := 1; j < len(str[i]); j = j + 4 {
			ar[i][j] = string(str[i][j])
		}
	}
	r := regexp.MustCompile(`\w+`)
	ar2 := make([][]string, 9)

	for i := range ar2 {
		ar2[i] = make([]string, 9) // Assuming 3 columns based on your desired output
	}
	b := 0
	for i := 0; i < len(ar); i++ {

		for j := 0; j < len(ar[i]); j++ {
			if r.MatchString(ar[i][j]) || ar[i][j] == " " {
				ar2[i][b] = ar[i][j]
				b++
			}
		}
		b = 0
	}

	_ = ar
	ret := int(str2[len(str2)-1])
	return ar2, ret
}

func move(ar [][]string, N, a, b int) [][]string {
	temp := ar[a][len(ar[a])-N:]
	//slices.Reverse(temp)
	ar[b] = append(ar[b], temp...)
	ar[a] = ar[a][:len(ar[a])-N]
	//fmt.Println(ar)
	return ar
}

//func main() {
//	ar, size := parse()
//	r := regexp.MustCompile(`(?s)(\d+)`)
//	input.txt, _ := os.Open("input.txt.txt")
//	defer input.txt.Close()
//	xxx, _ := ioutil.ReadAll(input.txt)
//	mv := r.FindAllString(string(xxx), -1)
//	mv2 := make([][]int, len(mv)/3)
//	for i := 0; i < len(mv2); i++ {
//		mv2[i] = make([]int, 3)
//	}
//	counter := 0
//	for i := 0; i < len(mv2); i++ {
//		for j := 0; j < 3; j++ {
//			mv2[i][j], _ = strconv.Atoi(mv[counter])
//			counter++
//		}
//	}
//	for i := 0; i < len(mv2); i++ {
//		ar = move(ar, mv2[i][0], mv2[i][1], mv2[i][2], size)
//	}
//
//	//fmt.Println(ar)
//	_, _ = ar, size
//}
//func parse() ([][]string, int) {
//	input2, _ := os.Open("input2.txt")
//	defer input2.Close()
//	xxx, _ := ioutil.ReadAll(input2)
//	str2 := string(xxx)
//	curStack := strings.Split(str2, "\n")
//	re := regexp.MustCompile(`(\d+)`)
//	max2 := re.FindAllString(curStack[len(curStack)-1], -1)
//	max3 := max2[len(max2)-1]
//	str := strings.Split(str2, "\n")
//	max4, _ := strconv.Atoi(max3)
//	ar := make([][]string, max4-1)
//	for i := 0; i < len(str)-1; i++ {
//		temp := make([]string, len(str[i]))
//		ar[i] = temp
//		for j := 1; j < len(str[i]); j = j + 4 {
//			ar[i][j] = string(str[i][j])
//		}
//	}
//	r := regexp.MustCompile(`\w+`)
//	ar2 := make([][]string, 9)
//
//	for i := range ar2 {
//		ar2[i] = make([]string, 9) // Assuming 3 columns based on your desired output
//	}
//	b := 0
//	for i := 0; i < len(ar); i++ {
//
//		for j := 0; j < len(ar[i]); j++ {
//			if r.MatchString(ar[i][j]) || ar[i][j] == " " {
//				ar2[i][b] = ar[i][j]
//				b++
//			}
//		}
//		b = 0
//	}
//
//	_ = ar
//	ret := int(str2[len(str2)-1])
//	return ar2, ret
//}
//func move(arry [][]string, a, b, c, size int) [][]string {
//	x := len(arry)
//	y := len(arry[0])
//	t := make([][]string, x)
//	for i := 0; i < len(t); i++ {
//		t[i] = make([]string, y)
//	}
//	for i := 0; i < x; i++ {
//		for j := 0; j < len(arry[i]); j++ {
//			fmt.Println(i, " ", j)
//			t[j][i] = arry[i][j]
//		}
//	}
//	fmt.Println(t)
//
//	for i := 0; i < len(t[0]); i++ {
//		temp := t[b-1]
//		if temp[i] == " " {
//			temp = temp[1:]
//			t[b-1] = temp
//		} else {
//			break
//		}
//	}
//	temp := t[b-1]
//	temp2 := temp[0:a]
//	for i := len(t[0]) - 1; i > -1; i-- {
//		if t[c-1][i] == " " {
//			t[c-1] = t[c-1][1:]
//		}
//
//	}
//
//	slices.Reverse(temp2)
//	newSlice := make([]string, 0)
//	newSlice = append(newSlice, temp2...)
//	newSlice = append(newSlice, t[c-1]...)
//	t[c-1] = newSlice
//
//	t[b-1] = t[b-1][a:]
//	//fmt.Println(t)
//
//	// Find the length of the longest inner slice
//	maxLen := 0
//	for _, row := range t {
//		if len(row) > maxLen {
//			maxLen = len(row)
//		}
//	}
//	ret := make([][]string, 9)
//	for i := 0; i < len(t); i++ {
//		ret[i] = make([]string, 0)
//	}
//	for i := 0; i < len(t); i++ {
//		//ret[i] = make([]string, 1)
//		//fmt.Println("a")
//		for j := 0; j < len(t[i]); j++ {
//			//fmt.Println("b")
//			ret[i] = append(ret[i], t[i][j])
//		}
//		//temp := []string{}
//		//ret = append(ret, temp)
//		fmt.Println("")
//	}
//	fmt.Println(ret)
//	return ret
//}
