package main

import "fmt"
import "os"
import "bufio"

func ReadFile() []string{
	var testo []string
	f,_:=os.OpenFile("input_test.txt")
	scanner:=bufio.NewScanner(f)
	for scanner.Scan(){
		testo = append(testo,scanner.Text())
	}
	return testo
}

func main(){

}