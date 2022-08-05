package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatusHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Hello world.")
		ctx.JSON(http.StatusOK, gin.H{"player": ctx.Value("player_name"), "seat id": ctx.Value("seat_id")})
	}
}
