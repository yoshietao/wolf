package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RouterGroupV1  string = "wolf/api/v1"
	SessionTokenId string = "session-id"
)

type HandlerEntry struct {
	Path      string
	Method    string
	Functions []gin.HandlerFunc
}

func Register(db *sql.DB) *gin.Engine {

	r := gin.New()
	// Todo: customer logger
	r.Use(gin.Logger(), gin.Recovery())

	// Authn: get and set session token
	authn := AuthnticationMiddleware(db)

	routes := []HandlerEntry{
		{
			Path:   "user",
			Method: http.MethodPost,
			Functions: []gin.HandlerFunc{
				CreateUserHandler(db),
			},
		},
		{
			Path:   "login",
			Method: http.MethodPost,
			Functions: []gin.HandlerFunc{
				UserLoginHandler(db),
			},
		},
		{
			Path:   "seat",
			Method: http.MethodPost,
			Functions: []gin.HandlerFunc{
				authn,
				SelectSeatHandler(db),
			},
		},
		{
			Path:   "status",
			Method: http.MethodGet,
			Functions: []gin.HandlerFunc{
				GetStatusHandler(db),
			},
		},
	}

	for _, route := range routes {
		apiGroup := r.Group(RouterGroupV1)
		apiGroup.Handle(route.Method, route.Path, route.Functions...)
	}

	return r
}
