package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	createFiles(1000000)

}

func createFiles(n int) {
	var newDir string = "newDir"

	os.Mkdir(newDir, 0700 )

	b,_ := filepath.Abs("/home/e4t4g/go/src/Go-2lvl-GB/newDir")

	for i:=0; i<=n; i++ {
		a, err := os.Create(b+"/"+fmt.Sprintf("file_%d.txt", i+1))
		if err != nil {
			a.Close()
		}
		file, _ := os.OpenFile("file_", os.O_WRONLY, 7777)
		log.Println(file)
		defer file.Close()
	}

}