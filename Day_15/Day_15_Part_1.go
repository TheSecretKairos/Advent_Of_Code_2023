package main

import "fmt"
import "os"
import "bufio"
import "strings"

func ReadText() []string {
	var read string
	var res []string
	f,_:=os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		read += scanner.Text()
	}
	res = strings.Split(read, ",")
	return res
}

func GetCode(s string) int {
	var current_value int
	for _, c := range s {
		fmt.Println(string(c), "->", int(c))
		current_value = (current_value + int(c)) * 17 % 256
	}
	fmt.Println("tot of ", s, "=", current_value)
	return current_value
}

func GetCodes(codici []string) {
	var tot int
	for _, codice := range codici {
		tot += GetCode(codice)
	}
	fmt.Println("tot ->", tot)
}

func main() {
	read := ReadText()
	GetCodes(read)
}
