package main

import "fmt"
import "bufio"
import "os"
import "unicode"
import "strconv"
//import "strings"

func ReadFile() string{
	var s string
	f,_:=os.Open("input.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		s+=scanner.Text()
	}
	fmt.Println("Testo letto")
	fmt.Println(s)
	return s
}

/*

		if unicode.IsDigit(c){
				s_num+=string(c)
			}else{
				if symbol && len(s_num)!=0{
					num,_:=strconv.Atoi(s_num)
					tot+=num
					s_num=""
					prev_symbol = true
				}
			}
			symbol = true
*/

func IsOkay(s string) (bool,int){
	var num_s string
	for i,c:=range(s){
		if unicode.IsDigit(c){
			bt:=rune(s[i-1])
			if !unicode.IsDigit(bt) && len(num_s) !=0{
				return false,0
			} 
			num_s+=string(c)
		}
	}
	num,_:=strconv.Atoi(num_s)
	fmt.Println("Numero ->",num,"\ttrue")
	return true,num
}

func AnalizeText(s string){
//	var s_num string
	var tot int 
	var prev_symbol bool = false
//	var symbol bool = false
	var prev_index int
	for i,c:=range(s){
		if c != '.'{
			if !(unicode.IsDigit(c)){
				if prev_symbol{
					fmt.Println("sottostringa ->",s[prev_index:i+1])
					ok,n:=IsOkay(s[prev_index:i+1])
					if ok{
						tot+=n
					}
					prev_index = i
					prev_symbol = false
				}
				prev_symbol = true
				prev_index = i
			}
		}
	}
	fmt.Println("Totale ->",tot)
}

func main(){
	testo:=ReadFile()
	AnalizeText(testo)
}