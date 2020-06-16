package controller

import (
	"github.com/alimnastaev/testapi/internal/api/controller/utils"
)

// Controller combines all of the different types of controllers
type Controller struct {
	*utils.Utils
}

// New returns a new implementation of the controllers
func New() *Controller {
	return &Controller{
		Utils: utils.New(),
	}
}
