package models

import (
	uuid "github.com/satori/go.uuid"
)

type Photo struct {
	ID uuid.UUID
}

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
