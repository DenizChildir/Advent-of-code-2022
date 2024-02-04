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
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	re := regexp.MustCompile(`(\w+) (-?\d+)`)
	strArry := strings.Split(str2, "\n")
	//q := []int{}
	//timer := 0
	reg := 1
	cyc := 1
	fmt.Println(len(strArry))
	for i := 0; i < len(strArry); i++ {

		op := re.FindAllStringSubmatch(strArry[i], -1)
		if len(op) == 0 {
			//fmt.Println("cyc=", cyc, " reg=", reg, " i=", i, "noop")
			cyc++
			fmt.Println("cyc=", cyc, " reg=", reg, " i=", i, "noop")
			continue
		} else {
			temp, _ := strconv.Atoi(op[0][2])
			//fmt.Println("cyc=", cyc, " reg=", reg, " i=", i, " temp=", temp)
			cyc += 2
			fmt.Println("cyc=", cyc-1, " reg=", reg, " i=", i, " temp=", temp)
			reg += temp
			fmt.Println("cyc=", cyc, " reg=", reg, " i=", i, " temp=", temp)
			continue
		}

	}
	//for i := 0; i < len(strArry); i++ {
	//	fmt.Println(i, " ", reg)
	//	op := re.FindAllStringSubmatch(strArry[i], -1)
	//	if timer%2 == 0 && len(q) > 0 {
	//		reg = reg + q[0]
	//		q = append(q[:0], q[0+1:]...)
	//
	//	}
	//	if len(op) == 0 {
	//		continue
	//	} else {
	//		timer++
	//		temp, _ := strconv.Atoi(op[0][2])
	//		q = append(q, temp)
	//	}
	//
	//	if i == 20 || i == 60 || i == 100 || i == 140 || i == 180 || i == 220 {
	//		//sum = sum + (reg * i)
	//	}
	//
	//}
	fmt.Println(reg)

}
