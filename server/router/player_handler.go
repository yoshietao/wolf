package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshietao/wolf/server/apimodels"
	"github.com/yoshietao/wolf/server/service"
)

func SelectSeatHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var player apimodels.PlayerInput
		if err := ctx.ShouldBindJSON(&player); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.SelectSeat(ctx, db, player.SeatId)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		// ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("user %s is registered", req.UserName)})
	}
}
