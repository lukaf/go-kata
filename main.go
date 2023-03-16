package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"archive/zip"

	"github.com/nlpodyssey/gopickle/pickle"
	"github.com/nlpodyssey/gopickle/types"
)

func challenge0() {
	fmt.Println(int(math.Pow(2, 38)))
}

func challenge1() {
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

func challenge2() {
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

func challenge3() {
	// http://www.pythonchallenge.com/pc/def/equality.html
	const datafile string = "bodyguard.txt"
	data, err := os.ReadFile(datafile)
	if err != nil {
		panic(err)
	}
	output := []byte{}

	var re = regexp.MustCompile(`[^A-Z][A-Z]{3}(?P<foo>[a-z])[A-Z]{3}[^A-Z]`)

	for _, match := range re.FindAllSubmatch(data, -1) {
		output = append(output, match[1]...)
	}

	fmt.Println(string(output))
}

func challenge4() {
	url := "http://www.pythonchallenge.com/pc/def/linkedlist.php?nothing="
	nothing := "91706"
	re := regexp.MustCompile(`and the next nothing is (\d+)`)

	for {
		uri := url + nothing
		response, err := http.Get(uri)
		if err != nil {
			break
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		fmt.Println(string(body))

		if string(body) == "Yes. Divide by two and keep going." {
			nothingInt, err := strconv.Atoi(nothing)
			if err == nil {
				nothing = fmt.Sprintf("%v", nothingInt/2)
			}
		} else {
			if err == nil {
				match := re.FindAllSubmatch(body, 1)
				nothing = string(match[0][1])
			}
		}
	}
}

func challenge5() {
	// http://www.pythonchallenge.com/pc/def/peak.html
	// todo: get the banner.p from http://www.pythonchallenge.com/pc/def/banner.p
	foo, err := pickle.Load("banner.p")
	if err != nil {
		panic(err)
	}
	lines := foo.(*types.List)

	for line := 0; line < lines.Len(); line++ {

		tuple_list := lines.Get(line).(*types.List)
		for tuple := 0; tuple < tuple_list.Len(); tuple++ {
			t := tuple_list.Get(tuple).(*types.Tuple)
			fmt.Printf("%s", strings.Repeat(t.Get(0).(string), t.Get(1).(int)))
		}

		fmt.Printf("\n")
	}
}

func challenge6() {
	// http://www.pythonchallenge.com/pc/def/channel.html
	zipfile, err := zip.OpenReader("channel.zip")
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()
	// fmt.Printf("%#v", zipfile)
	currentNum := "90052"
	
	// fmt.Println(currentfile)

//	for _, y := range zipfile.File {
//		fmt.Println("archive content includes:", y.Name)
//	}
	for {
		currentfile := fmt.Sprintf("%s.txt", currentNum)
		f, err := zipfile.Open(currentfile)
		if err != nil {
			panic(err)
		}

		data, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		println(string(data))
		re := regexp.MustCompile(`Next nothing is (\d+)`)
		fmt.Println(string(data))
		match := re.FindAllSubmatch(data, 1)
		currentNum = string(match[0][1])
		// TODO: Collect the comments
	}
}

func main() {
	challenge6()
}
