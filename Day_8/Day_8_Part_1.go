package main

import "fmt"
import "bufio"
import "os"
import "strings"

func ReadInput() (string,[]string){
	var res []string
	file,_:=os.Open("input.txt")
	defer file.Close()
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		res = append(res,scanner.Text())
	}
	steps:=res[0]
	coordinates:=res[2:]
	return steps,coordinates
}



func FormatCoordinates(cordinate []string) (map[string][]string){
	var mappatura map[string][]string
	mappatura = make(map[string][]string)
	for _,cord:=range(cordinate){
		var slRL []string = make([]string,2)
		cord=strings.ReplaceAll(cord," ","")
		splitted:=strings.Split(cord,"=")
		temp:=strings.Split(splitted[1],",")
		slRL[0]=temp[0][1:]
		slRL[1]=temp[1][:len(temp[1])-1]
		mappatura[splitted[0]]=slRL
	}
	return mappatura
}

func FindDestination(steps string,coordinates map[string][]string){
	var next string = "AAA"
	var cont int
	var found bool 
	for !found{
		for _,c:=range(steps){
			if c=='R'{
				next=coordinates[next][1]
				cont++
				if next=="ZZZ"{
					found = true
					break
				}
			}else{
				next=coordinates[next][0]
				cont++
				if next=="ZZZ"{
					found = true
					break
				}
			}
		}
	}
	fmt.Println("Result ->",cont)
}

func main(){
	steps,coordinates_daFormattare:=ReadInput()
	coordinates:=FormatCoordinates(coordinates_daFormattare)
	FindDestination(steps,coordinates)
}