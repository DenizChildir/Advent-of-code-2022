package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main1() {
	input, _ := os.Open("input.txt.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str := string(xxx)
	str = strings.ReplaceAll(str, "\r", "x")
	str = strings.ReplaceAll(str, "\n", "y")
	groups := strings.Split(str, "xyxy")
	//fmt.Printf("%q", groups)

	result := make([]int, len(groups))
	for i, group := range groups {
		//fmt.Printf("%q \n", group)
		re := regexp.MustCompile(`\d+`)
		numbers := re.FindAllString(group, -1)
		for _, k := range numbers {
			temp, _ := strconv.Atoi(k)
			result[i] = result[i] + temp
		}
	}
	sort.Ints(result)
	fmt.Printf("%v\n", result)
	//result2:=make([][]int,len(result))
	//for

	//result2 := [][]int{{1000, 2000, 3000}, {4000}, {5000, 6000}, {7000, 8000, 9000}, {10000}}
	//c := len(result2)
	//res := make([]int, c)
	//for i, _ := range result2 {
	//	for j, _ := range result2[i] {
	//		res[i] = res[i] + result2[i][j]
	//	}
	//
	//}
	//println(res)
	//fmt.Printf("%v\n", result)
}
