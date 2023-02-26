package utils

import (
	"log"
	"os"
)

func CreateWorkDirectory() {
	if err := os.Mkdir("photos", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
