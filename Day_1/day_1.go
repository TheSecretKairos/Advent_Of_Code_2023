package main

import "fmt"
import "bufio"
import "os"
import "unicode"
import "strconv"
import "strings"

func ReadFile() []string{
	var str []string
	f,_:=os.Open("input.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		str = append(str,scanner.Text())
	}
	return str
}

func ReadString(s string) int{
	var nums_s []string
	for _,c:=range(s){
		if unicode.IsDigit(c){
			nums_s = append(nums_s,string(c))
		}
	}
	if len(nums_s)==0{
		return 0
	}
	if len(nums_s) == 1{
		str:=strings.Repeat(nums_s[0],2)
		num,_:=strconv.Atoi(str)
		return num
	}
	str:=nums_s[0]+nums_s[len(nums_s)-1]
	num,_:=strconv.Atoi(str)
	return num
}

func CalculateNumber(sl []string){
	var tot int
	for _,s:=range(sl){
		tot+=ReadString(s)
	}
	fmt.Println("Totale ->",tot)
}

func main(){
	lettura:=ReadFile()
	CalculateNumber(lettura)
}