package main

import (
	"fmt"
	"os"
)


func main(){
	var err error
	FileInfo, err := os.Stat("PEBEWE")
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	if FileInfo.IsDir(){
		fmt.Println("PEBEWE adalah sebuah direktori.")
	}else{
		fmt.Println("PEBEWE adalah sebuah file.")
	}
}