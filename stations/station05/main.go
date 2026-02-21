package main

import (
	"fmt"
	"os"
)

func main() {
	sampleFile := "./sample.txt"
	CheckFileExist(sampleFile)

	notExistFile := "./abc.txt"
	CheckFileExist(notExistFile)
}

func CheckFileExist(path string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Printf("ファイルが見つかりました")
	} else if os.IsNotExist(err){
		fmt.Printf("ファイルが見つかりませんでした")
	}
}
