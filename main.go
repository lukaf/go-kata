package main

import "fmt"
import "math"

func challenge1() {
	fmt.Println(int(math.Pow(2, 38)))
}

func main() {
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