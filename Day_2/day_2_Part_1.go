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

func Leggi() []string{
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

func LeggiNumeroGioco(s string) int{
	var str string
	for _,c:=range(s){
		if unicode.IsDigit(c){
			str+=string(c)
		}
	}
	num,_:=strconv.Atoi(str)
	return num
}

func LeggiColore(s string) bool{
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
	switch(color){
	case "red":
		if digit>RED{
			return false
		}
		return true
	case "blue":
		if digit>BLUE{
			return false
		}
		return true
	case "green":
		if digit>GREEN{
			return false
		}
		return true
	default: 
		return false
	}
}

func LeggiEstrazione(s string) bool{
	str:=strings.Split(s,",")
	for _,estrazione:=range(str){
		if !LeggiColore(estrazione){
			return false
		}
	}
	return true
}

func LeggiEstrazioni(s string) bool{
	str:=strings.Split(s,";")
	for _,estrazione:=range(str){
		if !LeggiEstrazione(estrazione){
			return false
		}
	}
	return true
}

func AnalizzaStringa(s string) (bool,int){
	var game_num int
	str:=strings.Split(s,":")
	game_num = LeggiNumeroGioco(str[0])
	return LeggiEstrazioni(str[1]),game_num
}

func AnalizzaStringhe(sl []string){
	var tot int 
	for _,s:=range(sl){
		check,num:=AnalizzaStringa(s)
		if check{
			tot+=num
		}
	}
	fmt.Println("totale ->",tot)
}

func main(){
	testo:=Leggi()
	AnalizzaStringhe(testo)
}