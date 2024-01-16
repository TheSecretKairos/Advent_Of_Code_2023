package main

import "fmt"
import "bufio"
import "os"
import "unicode"
import "strconv"
import "strings"

// func Leggi() []string{
// 	var str []string
// 	scanner:=bufio.NewScanner(os.Stdin)
// 	for scanner.Scan(){
// 		str = append(str,scanner.Text())
// 	}
// 	return str
// }

func Leggi() []string{
	var str []string
	f,_:=os.Open("input.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		str = append(str,scanner.Text())
	}
	return str
}

// func Leggi() []string{
// 	var str []string
// 	var f *os.File
// 	var err error
// 	var line string
// 	f,err=os.Open("input.txt")
// 	defer f.Close()
// 	r:=bufio.NewReader(f)
// 	for err==nil{
// 		line,err=r.ReadString('\n')
// 		line = strings.TrimSuffix(line,"\n")
// 		str = append(str,line)
// 	}
// 	return str
// }


func LeggiStringa(s string) int{
	var nums_s []string
	for _,c:=range(s){
		if unicode.IsDigit(c){
			nums_s = append(nums_s,string(c))
		}
	}
	if len(nums_s)==0{
		fmt.Println("S ->",s," ---\tNumero in riga: ",0)
		return 0
	}
	if len(nums_s) == 1{
		str:=strings.Repeat(nums_s[0],2)
		num,_:=strconv.Atoi(str)
		fmt.Println("S ->",s,"---\tNumero in riga: ",num)
		return num
	}
	str:=nums_s[0]+nums_s[len(nums_s)-1]
	num,_:=strconv.Atoi(str)
	fmt.Println("S ->",s,"---\tNumero in riga: ",num)
	return num
}

func CalcolaNumero(sl []string){
	var tot int
	for _,s:=range(sl){
		tot+=LeggiStringa(s)
	}
	fmt.Println("Totale ->",tot)
}

func main(){
	lettura:=Leggi()
	//	fmt.Println("Lettoura ->")
	//fmt.Println(lettura)
	CalcolaNumero(lettura)
}