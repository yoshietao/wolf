package router

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/yoshietao/wolf/server/authz"
	"github.com/yoshietao/wolf/server/db/models"
)

func AuthnticationMiddleware(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var sessionToken string
		cookie, err := ctx.Request.Cookie(SessionTokenId)
		if err == nil {
			sessionToken = cookie.Value
		}

		user, err := models.Users(models.UserWhere.UUID.EQ(sessionToken), qm.Load(models.UserRels.UserIdPlayers)).One(ctx, db)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fmt.Println("User", user.UserName, user.R.UserIdPlayers)

		ctx.Set(authz.UserKey, user)
	}
}
