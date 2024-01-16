package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"
import "unicode"

const RED = 12
const GREEN = 13
const BLUE = 14
var COLORS = []string{"red","green","blue"}

func ReadFile() []string{
	var testo []string
	f,_:=os.Open("input.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		testo = append(testo,scanner.Text())
	}
	for i:=0;i<len(testo);i++{
		testo[i]=strings.ReplaceAll(testo[i]," ","")
	}
	return testo
}

func ReadColor(s string) (string,int){
	var digit int
	var temp string
	var color string
	for _,c:=range(s){
		if unicode.IsDigit(c){
			temp+=string(c)
		}else{
			color+=string(c)	
		}
	}
	digit,_=strconv.Atoi(temp)
	return color,digit
}

func ReadEstraction(s string) map[string]int{
	//fmt.Println("stringa ->",s)
	var mapMaxColor map[string]int
	mapMaxColor = make(map[string]int)
	str:=strings.Split(s,",")
	for _,estrazione:=range(str){
		color,num:=ReadColor(estrazione)
		if _,ok:=mapMaxColor[color];ok{
			if mapMaxColor[color]>num{
				mapMaxColor[color]=num
			}
		}else{
			mapMaxColor[color]=num
		}
	}
	return mapMaxColor
}

func LeggiEstrazioniColore(s string) int{
	var tot int = 1
	str:=strings.Split(s,";")
	var mapGame map[string]int
	mapGame= make(map[string]int)
	for _,estrazione:=range(str){
		mapRec:=ReadEstraction(estrazione)
		for _,col:=range(COLORS){
			if mapGame[col]<mapRec[col]{
				mapGame[col]=mapRec[col]
			}
		}
	}
	for _,col:=range(COLORS){
		if _,ok:=mapGame[col];ok{
			if mapGame[col]!=0{
				tot*=mapGame[col]
			}
		}
	} 
	return tot
}

func AnalizeString(s string) (int){
	str:=strings.Split(s,":")
	return LeggiEstrazioniColore(str[1])
}

func AnalizeStrings(sl []string){
	var tot int 
	for _,s:=range(sl){
		num:=AnalizeString(s)
		tot+=num
	}
	fmt.Println("totale colore ->",tot)
}


func main(){
	testo:=ReadFile()
	AnalizeStrings(testo)
}