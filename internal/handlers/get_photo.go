package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"io"
	"main/internal/models"
	"net/http"
	"os"
	"strconv"
)

func GetPhoto() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		photoId := vars["id"]
		quality, err := strconv.Atoi(r.URL.Query().Get("quality"))
		if quality != 100 && quality != 75 && quality != 50 && quality != 25 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("quality need to be 100/75/50/25")})
			return
		}

		id, err := uuid.FromString(photoId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("can't get id")})
			return
		}

		img, err := os.Open("./photos/" + id.String() + "/" + strconv.Itoa(quality) + ".jpeg")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("photo is not availabe now")})
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Itoa(quality)+".jpeg")
		w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusOK)
		_, err = io.Copy(w, img)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.ErrorResponse{Error: fmt.Sprintf("can't copy img")})
			return
		}
	}
}
