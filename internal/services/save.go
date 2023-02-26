package services

import (
	uuid "github.com/satori/go.uuid"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func SaveOriginalPhoto(photoId uuid.UUID, img image.Image) {
	if err := os.Mkdir("./photos/"+photoId.String(), os.ModePerm); err != nil {
		log.Fatal(err)
	}
	out100, err := os.Create("./photos/" + photoId.String() + "/" + "100.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	if err = jpeg.Encode(out100, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatal(err)
	}
}
