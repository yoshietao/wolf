package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func GetStatusHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
