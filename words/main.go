package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

func removeNoLetters(s string) string {
	text := strings.ReplaceAll(s, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\t", "")
	return text
}

func countVowelLetters(s string) (vowelCount int) {
	word := []byte(s)
	for utf8.RuneCount(word) > 0 {
		r, size := utf8.DecodeRune(word)
		word = word[size:]
		if r == 'A' || r == 'a' || r == 260 || r == 261 || r == 'E' || r == 'e' || r == 280 || r == 281 ||
			r == 'I' || r == 'i' || r == 'O' || r == 'o' || r == 211 || r == 243 || r == 'U' || r == 'u' {
			vowelCount++
		}
	}
	return vowelCount
}

func writeToFile(f os.File, s string) {
	_, err := f.WriteString(s + " ")
	if err != nil {
		panic(err)
	}
}

func whichFile(filePath string) string {
	file, _ := os.Open(filePath)
	buffer, _ := ioutil.ReadAll(file)
	return string(buffer)

}

func words(filePath string) (even []string, odd []string) {
	fileEven, _ := os.Create("even.txt")
	fileOdd, _ := os.Create("odd.txt")

	fileAsString := whichFile(filePath)

	stringSlice := strings.Split(removeNoLetters(fileAsString), " ")
	for i := 0; i < len(stringSlice); i++ {
		if countVowelLetters(stringSlice[i])%2 == 0 {
			even = append(even, stringSlice[i])
			writeToFile(*fileEven, stringSlice[i])
		} else {
			odd = append(odd, stringSlice[i])
			writeToFile(*fileOdd, stringSlice[i])
		}
	}

	return even, odd
}

func main() {
	var flagFile string
	flag.StringVar(&flagFile, "file", "lorem.txt", "file to read")
	flag.Parse()
	
	fmt.Println(words(flagFile))
}
