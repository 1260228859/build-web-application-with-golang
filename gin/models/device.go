package models

import (
	db "gin/database"
	"log"
)

type Device struct {
	id int    `json:"id" from:"id"`
	sn string `json:"sn" from:"sn"`
}

func (device *Device) AddDevice() (id int, err error) {

}
