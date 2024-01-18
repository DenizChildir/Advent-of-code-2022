package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	ar, size := parse()
	r := regexp.MustCompile(`(?s)(\d+)`)
	input, _ := os.Open("input.txt")
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
	move(ar, 3, 1, 9, size)
	//fmt.Println(ar)
	_ = ar
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
func move(arry [][]string, a, b, c, size int) [][]string {
	x := len(arry)
	y := len(arry[0])
	t := make([][]string, x)

	for i := 0; i < len(t); i++ {
		t[i] = make([]string, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			//fmt.Println(arry[j][i])
			t[j][i] = arry[i][j]
		}
	}
	for i := 0; i < len(t[0]); i++ {
		temp := t[b-1]
		if temp[i] == " " {
			temp = temp[1:]
			t[b-1] = temp
		} else {
			break
		}
	}

	temp := t[b-1]
	temp2 := temp[0:a]
	//t[c-1] = append(temp2, t[c-1]...)
	fmt.Println(len(t[c]))

	for i := len(t[c]) - 1; i > -1; i-- {
		if t[c-1][i] != " " {
			continue
		}
		t[c-1][i] = temp2[0]
		temp2 = temp2[1:]
	}

	slices.Reverse(temp2)
	t[c-1] = append(temp2, t[c-1]...)

	//fmt.Println(temp2)
	//fmt.Println(t[c-1])
	t[b-1] = t[b-1][a:]
	fmt.Println(t)

	return t
}
