package errors

import (
	"log"
	"os"
)

func FileClose(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(err)
	}
}
