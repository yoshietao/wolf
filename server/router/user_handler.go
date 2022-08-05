package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshietao/wolf/server/apimodels"
	"github.com/yoshietao/wolf/server/service"
)

func CreateUserHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req apimodels.UserInput
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uuid, err := service.CreateUser(ctx, db, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		ctx.SetCookie("session_token", uuid, 60*60*24, "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("user %s is registered", req.UserName)})
	}
}

func UserLoginHandler(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req apimodels.UserInput
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uuid, err := service.LoginUser(ctx, db, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		ctx.SetCookie("session_token", uuid, 60*60*24, "/", "localhost", false, false)
		ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprintf("user %s is logged in, session token returned", req.UserName)})
	}
}
