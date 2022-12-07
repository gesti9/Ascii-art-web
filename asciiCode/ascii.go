package asciiCode

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func Test(text string, banner string) {
	slice, err := getData(text)
	fileName := banner
	if err != nil {
		fmt.Println(err)
		return
	}
	if !checkFile(banner) {
		fmt.Println("the file does not match")
		return
	}

	file, err := ioutil.ReadFile("./asciiCode/" + fileName + ".txt")
	if err != nil {
		fmt.Println("read fail", err)
	}
	file = []byte(strings.ReplaceAll(string(file), "\r\n", "\n"))
	the := strings.Trim(string(file), "\n")

	sliceRead := strings.Split(string(the), "\n\n") // 3 файл не делит на массив, читает все одной сторокой
	var str [][]string
	arr := []string{"n"}

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			str = append(str, diVisionStr(sliceRead[slice[i][j]-32]))
		}
		str = append(str, arr)
	}

	newLine(str)
}

func checkFile(banner string) bool {
	fileName := banner
	switch fileName {
	case "standard":
		if FileMD5("./asciiCode/"+fileName+".txt") != "10a476398a02a81473aee0d0a5f813f9" {
			return false
		}
	case "shadow":
		if FileMD5("./asciiCode/"+fileName+".txt") != "a49d5fcb0d5c59b2e77674aa3ab8bbb1" {
			return false
		}
	case "thinkertoy":
		if FileMD5("./asciiCode/"+fileName+".txt") != "eef471ad03be9d13027560644dda8359" {
			return false
		}
	}
	return true
}

func FileMD5(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func getData(s string) ([]string, error) {
	count := 0

	Arg := s

	for _, v := range Arg {
		if v < 32 || v > 126 {
			return nil, nil
		}
	}

	if Arg == "" {
		return nil, nil
	}
	for _, v := range Arg {
		if string(v) == "\\" {
			count++
		}
	}
	if count*2 == len(Arg) {
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return nil, nil
	}

	if len(Arg) == 0 {
		return nil, nil
	}

	slice := strings.ReplaceAll(Arg, "\\n", "\n")
	sliceOne := strings.Split(slice, "\n")
	for i := 0; i < len(sliceOne); i++ {
		if Arg[i] < 32 || Arg[i] > 126 {
			if Arg[i] != '\n' {
				return nil, errors.New("error with argument2")
			}
		}
	}

	return sliceOne, nil
}

func diVisionStr(s string) []string {
	slice := strings.Split(s, "\n")
	return slice
}

func form(slice [][]string, file *os.File) {
	var str []string
	var line string

	for i := 0; i < 8; i++ {
		for j := 0; j < len(slice); j++ {
			if slice[j][0] == "" {
				slice[j][0] = "      "
			}
			line += slice[j][i]
		}
		str = append(str, line)
		line = ""
	}

	for _, v := range str {
		if v == "" {
			// fmt.Println("")
			return
		}
		file.WriteString(v)
		// fmt.Println(v)

		file.WriteString("\r\n")
	}
}

func newLine(slice [][]string) {
	file, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var str [][]string
	for i := 0; i < len(slice); i++ {
		if slice[i][0] == "n" {
			form(str, file)
			str = nil
		} else {
			str = append(str, slice[i])
		}
		count := 0
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == "" {
				count++
			}
		}

	}
}
