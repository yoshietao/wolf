package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yoshietao/wolf/server/db/models"
)

func AuthnticationMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sessionToken string
		cookie, err := ctx.Request.Cookie(SessionTokenId)
		if err == nil {
			sessionToken = cookie.Value
		}

		user, err := models.Users(models.UserWhere.UUID.EQ(sessionToken)).One(ctx, db)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(UserKey, user)
	}
}
