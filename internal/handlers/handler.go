package handlers

import (
	v1 "go.dsig.cn/shortener/internal/handlers/v1"
)

type Handler struct {
	ShortenHandler *v1.ShortenHandler
}

var (
	// Handle expose the handler to outside
	Handle *Handler
)

func init() {
	Handle = &Handler{
		ShortenHandler: v1.NewShortenHandler(),
	}
}
