package day3

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main3() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	re := regexp.MustCompile(`(?s)([[:alnum:]]+)`)
	stz := re.FindAllString(string(xxx), -1)
	sum := 0
	for i := 0; i < len(stz)-1; i = i + 3 {
		sum = sum + find(stz[i], stz[i+1], stz[i+2])
		fmt.Println(stz[i], stz[i+1], stz[i+2])
	}
	fmt.Println(sum)
}

/*
	func main() {
		input, _ := os.Open("input.txt")
		defer input.Close()
		xxx, _ := ioutil.ReadAll(input)
		stz := string(xxx)
		re:=regexp.MustCompile(`(?s)([[:alnum:]])`)
		stz=re.FindAllString(stz,-1)
		//stz = strings.ReplaceAll(stz, "\r", "\n")
		//fmt.Printf("%q", stz)
		stz2 := strings.Split(stz, "\n")
		_ = stz2
		sum := 0
		fmt.Printf("%q", stz2)
		//for i := 0; i < len(stz2)-1; i = i + 3 {
		//	sum = sum + find(stz2[i], stz2[i+1], stz2[i+2])
		//	fmt.Println(stz2[i], stz2[i+1], stz2[i+2])
		//}
		fmt.Println(sum)

}
*/
func find(str1, str2, str3 string) int {
	//fmt.Println(len(str1), " ", len(str2), " ", len(str3))
	for i, _ := range str1 {
		for j, _ := range str2 {
			for k, _ := range str3 {
				if str1[i] == str2[j] && str2[j] == str3[k] {
					fmt.Println(string(str1[i]))
					j := str1[i]
					if j < 91 {
						return int(j - 64 + 26)
					} else {
						return int(j - 96)
					}
				}
			}
		}
	}
	return 0

}

/*
func find(str string) int {
	l := len(str)
	str1 := str[l/2:]
	str2 := str[:l/2]
	for _, j := range str1 {
		if strings.Contains(str2, string(j)) {
			fmt.Println(string(j))
			if j < 91 {
				return int(j - 64 + 26)
			} else {
				return int(j - 96)
			}
		}
	}
	return 0

}
*/
//65 90 97 122
