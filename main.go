package main

import (
	"fmt"
	"math"
	"os"
)

func challenge1() {
	fmt.Println(int(math.Pow(2, 38)))
}

func challenge2() {
	const min = int('a')
	const max = int('z')

	//input := "g fmnc wms bgblr rpylqjyrc gr zw fylb. rfyrq ufyr amknsrcpq ypc dmp. bmgle gr gl zw fylb gq glcddgagclr ylb rfyr'q ufw rfgq rcvr gq qm jmle. sqgle qrpgle.kyicrpylq() gq pcamkkclbcb. lmu ynnjw ml rfc spj."
	input := "http://www.pythonchallenge.com/pc/def/map.html"
	output := make([]rune, 0)
	for _, letter := range input {
		// fmt.Printf("initial letter: %s  => ", string(letter))
		if int(letter) >= min && int(letter) <= max {
			n := int(letter) + 2
			if n > int('z') {
				n -= (int('z') - int('a') + 1)
			}
			letter = rune(n)
		}
		// fmt.Printf("%s\n", string(letter))
		output = append(output, letter)
	}

	fmt.Println(string(output))
}

func challenge3() {
	// http://www.pythonchallenge.com/pc/def/ocr.html
	const datafile string = "data.txt"
	data, err := os.ReadFile(datafile)
	if err != nil {
		panic(err)
	}

	type counter struct {
		index int
		count int
	}
	characterCount := map[rune]counter{}
	for i, r := range string(data) {
		value, ok := characterCount[r]
		if ok == false {
			characterCount[r] = counter{index: i, count: 0}
		}
		value.count++
		characterCount[r] = value
	}
	unorderedCharacterCount := []counter{}
	for _, v := range characterCount {
		if v.count == 1 {
			unorderedCharacterCount = append(unorderedCharacterCount, v)
		}
	}
	// TODO: Sort the `unorderedCharacterCount` with interface
}

func challenge4() {
	// http://www.pythonchallenge.com/pc/def/equality.html
}

func main() {
	challenge3()
}
