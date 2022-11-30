package main

import (
	"fmt"
	helper "github.com/babulal107/go-examples/pkg/utils"
	"os"
)

func main() {

	currentDir, _ := os.Getwd()
	fileName := "Indaram_Sales Invoice_1.pdf"
	tmpFileDir := currentDir + "/assets/idstats/" + fileName
	fmt.Println("removed file path : ", tmpFileDir)
	err := helper.RemoveFile(tmpFileDir)
	if err != nil {
		fmt.Println("error while delete file : ", err.Error())
	}

}
