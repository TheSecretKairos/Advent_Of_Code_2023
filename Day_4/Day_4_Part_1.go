package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func ReadFile() []string{
	var testo []string
	f,_:=os.Open("input.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		testo = append(testo,scanner.Text())
	}
	return testo
}

func CheckNumbers(sl_w,sl_n []int) int{
	fmt.Println("Controllo su ")
	fmt.Println(sl_w)
	fmt.Println("-----")
	fmt.Println(sl_n)
	var cont int
	for i:=0;i<len(sl_w);i++{
		for j:=0;j<len(sl_n);j++{
			if sl_w[i]==sl_n[j]{
				fmt.Println(sl_w[i],sl_n[j])
				if cont>1{
					cont*=2
				}else{
					cont++
				}
			}
		}
	}
	fmt.Println("numeri trovati ->",cont)
	return cont
}

func ConvertSlice(sl string) []int{
	var ints []int
	s:=strings.Split(sl," ")
	//fmt.Println("splittata ->",s)
	for _,nms:=range(s){
		num,_:=strconv.Atoi(nms)
		if num!=0{
		//	fmt.Println("num conv ->",num)
			ints = append(ints,num)
		}
	}
//	fmt.Println("Risultato conv",ints)
	return ints
}


func ReadScratchcards(s string)int {
	//var tot int
	fmt.Println("Stringa in esame ->",s)
	sequenze:=strings.Split(s,"|")
	seq_1:=ConvertSlice(sequenze[0])
	seq_2:=ConvertSlice(sequenze[1])
	//fmt.Println(sequenze[0])
	nums:=CheckNumbers(seq_1,seq_2)
	return nums
}

func AnalizeString(s string) int{
	str:=strings.Split(s,":")
	return ReadScratchcards(str[1])
}

func AnalizeStrings(sl []string){
	var tot int
	for _,s:=range(sl){
		tot+=AnalizeString(s)
	}
	fmt.Println("Totale ->",tot)
}

func main(){
	testo:=ReadFile()
	AnalizeStrings(testo)
}