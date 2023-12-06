package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	var total = 0
	for scan.Scan() {
		line := scan.Text()
		d, err := regexp.Compile(`\d`)
		if err != nil {
			panic(err)
		}
		digits := d.FindAllString(line, -1)
		num := digits[0] + digits[len(digits)-1]
		toSum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		total = total + toSum
	}
	fmt.Printf("The total is %d \n", total)
}
