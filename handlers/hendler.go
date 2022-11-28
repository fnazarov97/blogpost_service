package handlers

import (
	"article/config"
	"article/storage"
)

type Handler struct {
	IM   storage.StorageI
	Conf config.Config
}

func NewHandler(stg storage.StorageI, conf config.Config) Handler {
	return Handler{
		IM:   stg,
		Conf: conf,
	}
}
