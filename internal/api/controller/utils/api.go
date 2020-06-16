package utils

import (
	"github.com/alimnastaev/testapi/pkg/db"

	"log"
)

// Utils jkfldsjklfd
type Utils struct {
	DBSvc db.API
}

// New returns a new implementation of utilities
func New() *Utils {
	return &Utils{}
}

// Login will login to something
func (u *Utils) Login(email, password string) (authToken string, err error) {
	log.Println("here")
	return
}
