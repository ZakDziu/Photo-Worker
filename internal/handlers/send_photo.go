package handlers

import (
	"encoding/json"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"image/jpeg"
	"log"
	"main/internal/models"
	"main/internal/services"
	"main/utils"
	"net/http"
)

func SendPhoto(bus *utils.Bus) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("photo")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("can't get file")})
			return
		}
		defer file.Close()

		photoId := uuid.NewV4()
		img, err := jpeg.Decode(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("can't decode file")})
			return
		}

		services.SaveOriginalPhoto(photoId, img)
		photo := models.Photo{ID: photoId}

		photoPayload, err := json.Marshal(photo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("can't marshal id")})
			return
		}

		bus.Sender.Send(photoPayload, photoId)

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(models.MessageResponse{Message: fmt.Sprintf("file saved with id %v", photoId)})
		if err != nil {
			log.Panic(err)
			return
		}
		return
	}
}
