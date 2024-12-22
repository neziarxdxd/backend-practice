package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	handler *gin.Engine
	db      *gorm.DB
	// TODO: "tw" should be renamed as "sms" for consistency
}
