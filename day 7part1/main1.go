package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type folder struct {
	name  string
	sub   []*folder // Changed to a slice of pointers to folder structs
	files []*file
	size  int
	par   *folder
}

func main1() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	xxx, _ := ioutil.ReadAll(input)
	str2 := string(xxx)
	str := strings.Split(str2, "\n")
	//fmt.Printf("%q", str)
	root := parse(str)
	size(root)
	fmt.Println(count(root))

}
func count(cur *folder) int {
	counter := 0
	for _, child := range cur.sub {
		counter += count(child)
	}
	if cur.size <= 100000 {
		counter += cur.size
	}
	return counter
}

func size(cur *folder) {

	for _, child := range cur.sub {
		size(child)
	}
	for _, child := range cur.sub {
		cur.size = cur.size + child.size
	}
	for _, file := range cur.files {
		cur.size = cur.size + file.size
	}
}

func parse(str []string) *folder {
	root := &folder{name: "root"} // Changed to a pointer
	parsehelp(str, root, 1)
	return root
}

func parsehelp(str []string, fold *folder, j int) { // Changed to a pointer
	if j >= len(str) {
		return
	} else if strings.HasPrefix(str[j], "$ ls") {
		//fmt.Println("ls ", j)
		parsehelp(str, fold, j+1)
	} else if strings.HasPrefix(str[j], "$ cd ..") {
		if fold.par != nil {
			parsehelp(str, fold.par, j+1)
		} else {
			parsehelp(str, fold, j+1) // Stay in root if no parent
		}
	} else if strings.HasPrefix(str[j], "$ cd ") && !strings.HasPrefix(str[j], "$ cd ..") {
		re := regexp.MustCompile(`(\w+)`)
		name := re.FindAllStringSubmatch(str[j], -1)
		//fmt.Println("cd ", name[1][1], " ", j)
		for _, subfolder := range fold.sub { // Range over pointers
			if subfolder.name == name[1][1] {
				parsehelp(str, subfolder, j+1) // Pass the pointer
			}
		}
	} else if strings.HasPrefix(str[j], "dir") {
		re := regexp.MustCompile(`(\w+)`)
		splitStr := re.FindAllStringSubmatch(str[j], -1)
		//fmt.Println("dir ", splitStr[1][1], " ", j)
		child := &folder{name: splitStr[1][1], par: fold, size: 0} // Create a new folder as a pointer
		fold.sub = append(fold.sub, child)                         // Append the pointer
		parsehelp(str, fold, j+1)
	} else {
		re := regexp.MustCompile(`(\d+) (\w+.*)`)
		ar := re.FindStringSubmatch(str[j])

		if len(ar) >= 3 { // Ensure that there are at least three elements (including the full match)
			//fmt.Println(ar[1], " ", ar[2], " ", j)
			size, _ := strconv.Atoi(ar[1])
			file := &file{name: ar[2], size: size}
			fold.files = append(fold.files, file)
		}
		parsehelp(str, fold, j+1)
	}
}
