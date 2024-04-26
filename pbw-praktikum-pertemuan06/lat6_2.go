package main

import (
	"fmt"
	"os"
)



func main(){
	// membuat direktori baru
	err:= os.Mkdir("Afni Puspita Zahra", 0755)
	if err!= nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'Afni Puspita Zahra' berhasil dibuat.")
}