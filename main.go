package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input2, _ := os.Open("input2.txt")
	defer input2.Close()
	xxx, _ := ioutil.ReadAll(input2)
	str2 := string(xxx)
	curStack := strings.Split(str2, "\n")
	re := regexp.MustCompile(`(\d+)`)
	max2 := re.FindAllString(curStack[len(curStack)-1], -1)
	max3 := max2[len(max2)-1]
	//fmt.Println(max3)
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
				//fmt.Printf("%q", ar[i][j])
			}
		}
		b = 0
		println("")
	}
	//for i := 0; i < len(ar); i++ {
	//	for j := 0; j < len(ar[i]); j++ {
	//
	//		//if r.MatchString(ar[i][j]) || ar[i][j] == " " {
	//		//	ar2[i][j] = ar[i][j]
	//		//}
	//		fmt.Printf("%q", ar[i][j])
	//	}
	//}
	fmt.Printf("%q", ar2)
	_ = ar

}
