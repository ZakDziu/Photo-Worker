package services

import (
	uuid "github.com/satori/go.uuid"
	"image/jpeg"
	"log"
	"os"
)

func OptimizePhoto(id uuid.UUID) {
	imgIn, err := os.Open("./photos/" + id.String() + "/100.jpeg")
	if err != nil {
		log.Fatal("open" + err.Error())
	}

	img, err := jpeg.Decode(imgIn)
	if err != nil {
		log.Fatal("decode" + err.Error())
	}

	defer imgIn.Close()

	out75, err := os.Create("./photos/" + id.String() + "/75.jpeg")
	if err != nil {
		log.Fatal("out75" + err.Error())
	}

	out50, err := os.Create("./photos/" + id.String() + "/50.jpeg")
	if err != nil {
		log.Fatal("out3" + err.Error())
	}

	out25, err := os.Create("./photos/" + id.String() + "/25.jpeg")
	if err != nil {
		log.Fatal("out4" + err.Error())
	}

	if err = jpeg.Encode(out75, img, &jpeg.Options{Quality: 75}); err != nil {
		log.Fatal("encode out75" + err.Error())
	}
	if err = jpeg.Encode(out50, img, &jpeg.Options{Quality: 50}); err != nil {
		log.Fatal("encode out50" + err.Error())
	}
	if err = jpeg.Encode(out25, img, &jpeg.Options{Quality: 25}); err != nil {
		log.Fatal("encode out25" + err.Error())
	}
}
