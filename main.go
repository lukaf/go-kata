package main

import (
	"archive/zip"
	"fmt"
	"image/png"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

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

	commentHash := map[string]string{}

	for _, file := range zipfile.File {
		commentHash[file.Name] = file.Comment
	}

	for {
		currentfile := fmt.Sprintf("%s.txt", currentNum)
		fmt.Printf("%s", commentHash[currentfile])
		f, err := zipfile.Open(currentfile)
		if err != nil {
			panic(err)
		}

		data, err := io.ReadAll(f)
		if err != nil {
			panic(err)
		}
		re := regexp.MustCompile(`Next nothing is (\d+)`)
		match := re.FindAllSubmatch(data, 1)
		if len(match) < 1 {
			break
		}

		currentNum = string(match[0][1])
		f.Close()
	}
}

func challenge7() {
	// http://www.pythonchallenge.com/pc/def/oxygen.html

	file, err := os.Open("oxygen.png")
	if err != nil {
		log.Fatal("Kaput")
	}
	defer file.Close()

	image, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
	}
	bounds := image.Bounds()
	y := 47
	s := ""
	for x := bounds.Min.X; x < bounds.Max.X; x+=7 {
		colour := image.At(x, y)
		r16, g16, b16, _ := colour.RGBA()
		r8 := uint8(r16 >> 8)
		g8 := uint8(g16 >> 8)
		b8 := uint8(b16 >> 8)

		// fmt.Printf("r8=%v, g8=%v, b8=%v, a8=%v\n", r8, g8, b8, a8)

		if r8 == g8 && g8 == b8 {
		s += string(r8)
		}
	}
	fmt.Println(s)

	start := strings.Index(s, "[")
	end := strings.Index(s, "]")
	if start == -1 || end == -1 {
		fmt.Println("Could not find slice in string")
		return
	}
	sliceStr := s[start+1 : end]
	
	// Convert slice to ASCII characters
	asciiChars := ""
	for _, numStr := range strings.Split(sliceStr, ", ") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatal("Kaput")
		}
		asciiChars += string(num)
	}
	fmt.Println(asciiChars)
	
}

func main() {
	challenge7()
}
