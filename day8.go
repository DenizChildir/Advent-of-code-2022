package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("coding_qual_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wordMap := make(map[int]string)
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\d+) (\w+)`)
		match := re.FindStringSubmatch(line)
		if match != nil {
			position, _ := strconv.Atoi(match[1])
			wordMap[position] = match[2]
		}
	}
	var message []string
	for i, n := 1, 0; ; i++ {
		n = i * (i + 1) / 2
		if word, ok := wordMap[n]; ok {
			message = append(message, word)
		} else {
			break
		}
	}
	printMessage := strings.Join(message, " ")
	fmt.Println("Decoded message:", printMessage)
}
