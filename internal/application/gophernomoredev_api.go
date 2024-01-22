package application

import logger "github.com/sirupsen/logrus"

//go:generate go run github.com/golang/mock/mockgen -package gophernomoredevapimocks -destination=./gophernomoredevapimocks/gophernomoredevapi.go . GopherNoMoreDevAPI
type GopherNoMoreDevAPI interface {
}

type gophernomoredevapi struct {
	log *logger.Logger
}

func NewGopherNoMoreDevAPI(log *logger.Logger) *gophernomoredevapi {
	return &gophernomoredevapi{
		log: log,
	}
}
