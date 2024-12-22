package controllers

import "backend-practice/models"

type Controller struct {
	DB *models.Database
}

// ControllerInit creates a new UserController.
func ControllerInit(db *models.Database) *Controller {
	return &Controller{DB: db}
}
