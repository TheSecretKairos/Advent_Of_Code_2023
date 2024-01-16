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
	//defer f
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		testo = append(testo,scanner.Text())
	}
	// for _,s:=range(testo){
	// 	fmt.Println(s)	
	// }
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
	//fmt.Println("digit ->",digit,"color->",color)
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
	//fmt.Println("stringa ->",s)
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

func LeggiColorePerColore(s string) (string,int){
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
	fmt.Println("digit ->",digit,"color->",color)
	return color,digit
}

func LeggiEstrazioneColore(s string) map[string]int{
	//fmt.Println("stringa ->",s)
	var mapMaxColor map[string]int
	mapMaxColor = make(map[string]int)
	str:=strings.Split(s,",")
	for _,estrazione:=range(str){
		color,num:=LeggiColorePerColore(estrazione)
		if _,ok:=mapMaxColor[color];ok{
			if mapMaxColor[color]>num{
				mapMaxColor[color]=num
			}
		}else{
			mapMaxColor[color]=num
		}
	}
	//fmt.Println("mapMaxColor ->",mapMaxColor)
	return mapMaxColor
}

func LeggiEstrazioniColore(s string) int{
	var tot int = 1
	str:=strings.Split(s,";")
	var mapGame map[string]int
	mapGame= make(map[string]int)
	for _,estrazione:=range(str){
		mapRec:=LeggiEstrazioneColore(estrazione)
		for _,col:=range(COLORS){
			if mapGame[col]<mapRec[col]{
				mapGame[col]=mapRec[col]
			}
		}
	}
//	fmt.Println("mapGame ->",mapGame)
	for _,col:=range(COLORS){
		if _,ok:=mapGame[col];ok{
			if mapGame[col]!=0{
				tot*=mapGame[col]
			}
		}
	} 
	return tot
}

func AnalizzaStringaColore(s string) (int){
	//var game_num int
	str:=strings.Split(s,":")
	//game_num = LeggiNumeroGioco(str[0])
	return LeggiEstrazioniColore(str[1])
}

func AnalizzaStringheColore(sl []string){
	var tot int 
	for _,s:=range(sl){
		num:=AnalizzaStringaColore(s)
		tot+=num
	}
	fmt.Println("totale colore ->",tot)
}


func main(){
	testo:=Leggi()
	AnalizzaStringhe(testo)
	AnalizzaStringheColore(testo)
}