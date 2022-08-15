package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshietao/wolf/server/apimodels"
	"github.com/yoshietao/wolf/server/service"
)

func SelectSeatHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req apimodels.UserInput
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := service.CreateUser(ctx, db, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("user %s is registered", req.UserName)})
	}
}
