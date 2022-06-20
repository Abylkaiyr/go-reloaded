package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println("Error")
	}
}

func FormatInput(s string) []string {
	s1 := regexp.MustCompile(`\(.*?\)|\S*`).FindAllString(s, -1)
	return s1
}

func cap6(s []string) []string {
	for i := 0; i < len(s); i++ {
		digitCap := regexp.MustCompile(`\(cap, \d*?\)`).MatchString(s[i])
		numCap := regexp.MustCompile(`[\d]+`).FindAllString(s[i], -1)
		if digitCap {
			convToInt, _ := strconv.Atoi(numCap[0])
			for j := 0; j < convToInt; j++ {
				s[i-(j+1)] = strings.Title(s[i-(j+1)])
				s[i] = ""
			}
		}
	}
	return s
}

func cap(s []string) []string {
	for i := 0; i < len(s); i++ {
		justcap := regexp.MustCompile(`\(cap\)`).MatchString(s[i])
		if justcap {
			s[i-1] = (strings.Title(s[i-1]))
			s[i] = ""
		}
	}
	return s
}

func up6(s []string) []string {
	for i := 0; i < len(s); i++ {
		digitUp := regexp.MustCompile(`\(up,\s{1}\d*?\)`).MatchString(s[i])
		numCap := regexp.MustCompile(`[\d]+`).FindAllString(s[i], -1)
		if digitUp {
			convToInt, _ := strconv.Atoi(numCap[0])
			for j := 0; j < convToInt; j++ {
				s[i-(j+1)] = strings.ToUpper(s[i-(j+1)])
				s[i] = ""
			}
		}
	}
	return s
}

func up(s []string) []string {
	for i := 0; i < len(s); i++ {
		justup := regexp.MustCompile(`\(up\)`).MatchString(s[i])
		if justup {
			s[i-1] = (strings.ToUpper(s[i-1]))
			s[i] = ""
		}
	}
	return s
}

func low6(s []string) []string {
	for i := 0; i < len(s); i++ {
		digitLow := regexp.MustCompile(`\(low,\s{1}\d*?\)`).MatchString(s[i])
		numCap := regexp.MustCompile(`[\d]+`).FindAllString(s[i], -1)
		if digitLow {
			convToInt, _ := strconv.Atoi(numCap[0])
			for j := 0; j < convToInt; j++ {
				s[i-(j+1)] = strings.ToLower(s[i-(j+1)])
				s[i] = ""
			}
		}
	}
	return s
}

func low(s []string) []string {
	for i := 0; i < len(s); i++ {
		justlow := regexp.MustCompile(`\(low\)`).MatchString(s[i])
		if justlow {
			s[i-1] = (strings.ToLower(s[i-1]))
			s[i] = ""
		}
	}
	return s
}

func Hex(s []string) []string {
	for i := 0; i < len(s); i++ {
		HexParse := regexp.MustCompile(`\(hex\)`).MatchString(s[i])
		if HexParse {
			parsedHex, err := strconv.ParseInt(s[i-1], 16, 64)
			if err == nil {
				s[i-1] = strconv.Itoa(int(parsedHex))
			}
			s[i] = ""
		}
	}
	return s
}

func Bin(s []string) []string {
	for i := 0; i < len(s); i++ {
		BinParse := regexp.MustCompile(`\(bin\)`).MatchString(s[i])
		if BinParse {
			parsedBin, err := strconv.ParseInt(s[i-1], 2, 64)
			if err == nil {
				s[i-1] = strconv.Itoa(int(parsedBin))
			}
			s[i] = ""
		}
	}
	return s
}

func article(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		// An and a
		if s[i] == "a" && IsVowel(rune(s[i+1][0])) {
			s[i] = "an"
		}
		if s[i] == "A" && IsVowel(rune(s[i+1][0])) {
			s[i] = "An"
		}

	}
	return s
}

func IsVowel(letter rune) bool {
	for _, s := range "aeiouAEIOUhH" {
		if letter == s {
			return true
		}
	}
	return false
}

func formatText(s []string) []string {
	str1 := strings.Join(s, " ")
	var resStr string
	char := regexp.MustCompile(`[.,\?!:; ]+`)
	text := char.Split(str1, -1)

	punct := char.FindAllString(str1, -1)

	for i := 0; i < len(punct); i++ {
		resStr += strings.TrimSpace(text[i])
		resStr += strings.TrimSpace(punct[i])
		resStr += " "
	}

	resStr += text[len(text)-1]
	String1 := strings.Fields(resStr)
	return String1
}

func quote(s []string) []string {
	var logic bool = true
	for i := 0; i < len(s); i++ {
		if s[i] == "'" {
			if logic {
				s[i+1] = "'" + s[i+1]
				s[i] = ""
				logic = false
			} else {
				s[i-1] = s[i-1] + "'"
				s[i] = ""
				logic = true
			}
		}
	}
	return s
}

func main() {
	// Taking input
	arguments := os.Args[1:]
	if len(arguments) > 2 {
		fmt.Println("Enter only 2 arguments")
		return
	} else if len(arguments) < 2 {
		fmt.Println("Not enough arguments")
		return
	}
	content, err := os.ReadFile(arguments[0])
	ErrCheck(err)
	temp := FormatInput(string(content))

	// Processing
	var sampleText []string

	sampleText = cap6(temp)
	sampleText = cap(sampleText)
	sampleText = up6(sampleText)
	sampleText = up(sampleText)
	sampleText = low6(sampleText)
	sampleText = low(sampleText)
	sampleText = Hex(sampleText)
	sampleText = Bin(sampleText)
	sampleText = article(sampleText)
	sampleText = formatText(sampleText)
	sampleText = quote(sampleText)

	// Writing to output
	ResultText := strings.Join(sampleText, " ")
	finalResult := strings.Fields(ResultText)
	ResultText = strings.Join(finalResult, " ")
	res := ResultText + "\n"
	f, err := os.Create("result.txt")
	ErrCheck(err)
	f.WriteString(res)
}
