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
	"compress/bzip2"

	"github.com/nlpodyssey/gopickle/pickle"
	"github.com/nlpodyssey/gopickle/types"
	"github.com/fogleman/gg"
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
		asciiChars += string(rune(num))
	}
	fmt.Println(asciiChars)
	
}

func challenge8() {
	un := "BZh91AY&SYA\xaf\x82\r\x00\x00\x01\x01\x80\x02\xc0\x02\x00 \x00!\x9ah3M\x07<]\xc9\x14\xe1BA\x06\xbe\x084"
	pw := "BZh91AY&SY\x94$|\x0e\x00\x00\x00\x81\x00\x03$ \x00!\x9ah3M\x13<]\xc9\x14\xe1BBP\x91\xf08"
	un_reader := strings.NewReader(un)
	pw_reader := strings.NewReader(pw)
	unDecrypted := bzip2.NewReader(un_reader)
	pwDecrypted := bzip2.NewReader(pw_reader)
	username, err := io.ReadAll(unDecrypted)
	if err != nil {
		panic(err)
	}
	password, err := io.ReadAll(pwDecrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(username), string(password))
	// Huge:File
}

func challenge9() {
first := []int{146,399,163,403,170,393,169,391,166,386,170,381,170,371,170,355,169,346,167,335,170,329,170,320,170,
	310,171,301,173,290,178,289,182,287,188,286,190,286,192,291,194,296,195,305,194,307,191,312,190,316,
	190,321,192,331,193,338,196,341,197,346,199,352,198,360,197,366,197,373,196,380,197,383,196,387,192,
	389,191,392,190,396,189,400,194,401,201,402,208,403,213,402,216,401,219,397,219,393,216,390,215,385,
	215,379,213,373,213,365,212,360,210,353,210,347,212,338,213,329,214,319,215,311,215,306,216,296,218,
	290,221,283,225,282,233,284,238,287,243,290,250,291,255,294,261,293,265,291,271,291,273,289,278,287,
	279,285,281,280,284,278,284,276,287,277,289,283,291,286,294,291,296,295,299,300,301,304,304,320,305,
	327,306,332,307,341,306,349,303,354,301,364,301,371,297,375,292,384,291,386,302,393,324,391,333,387,
	328,375,329,367,329,353,330,341,331,328,336,319,338,310,341,304,341,285,341,278,343,269,344,262,346,
	259,346,251,349,259,349,264,349,273,349,280,349,288,349,295,349,298,354,293,356,286,354,279,352,268,
	352,257,351,249,350,234,351,211,352,197,354,185,353,171,351,154,348,147,342,137,339,132,330,122,327,
	120,314,116,304,117,293,118,284,118,281,122,275,128,265,129,257,131,244,133,239,134,228,136,221,137,
	214,138,209,135,201,132,192,130,184,131,175,129,170,131,159,134,157,134,160,130,170,125,176,114,176,
	102,173,103,172,108,171,111,163,115,156,116,149,117,142,116,136,115,129,115,124,115,120,115,115,117,
	113,120,109,122,102,122,100,121,95,121,89,115,87,110,82,109,84,118,89,123,93,129,100,130,108,132,110,
	133,110,136,107,138,105,140,95,138,86,141,79,149,77,155,81,162,90,165,97,167,99,171,109,171,107,161,
	111,156,113,170,115,185,118,208,117,223,121,239,128,251,133,259,136,266,139,276,143,290,148,310,151,
	332,155,348,156,353,153,366,149,379,147,394,146,399}

second := []int{156,141,165,135,169,131,176,130,187,134,191,140,191,146,186,150,179,155,175,157,168,157,163,157,159,
157,158,164,159,175,159,181,157,191,154,197,153,205,153,210,152,212,147,215,146,218,143,220,132,220,
125,217,119,209,116,196,115,185,114,172,114,167,112,161,109,165,107,170,99,171,97,167,89,164,81,162,
77,155,81,148,87,140,96,138,105,141,110,136,111,126,113,129,118,117,128,114,137,115,146,114,155,115,
158,121,157,128,156,134,157,136,156,136}
	// http://www.pythonchallenge.com/pc/return/good.html

	dc := gg.NewContext(600, 600)
	dc.SetRGB(1,1,1)
	dc.Clear()
	dc.SetRGB(1,0,0)

	for i:=0; i < len(first);i+=2 {
		x := first[i]
		y := first[i+1]
		dc.DrawPoint(float64(x),float64(y), 3.0)
	}
	dc.SetRGB(0,1,0)
	for i:=0; i < len(second);i+=2 {
		x := second[i]
		y := second[i+1]
		dc.DrawPoint(float64(x),float64(y), 3.0)
	}
	dc.Stroke()
	dc.SavePNG("foo.png")
	// Draws a male cow
}

func challenge10() {
	// http://www.pythonchallenge.com/pc/return/bull.html
}

func main() {
	challenge9()
}
