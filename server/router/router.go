package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RouterGroupV1 string = "wolf/api/v1"
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

	routes := []HandlerEntry{
		{
			Path:   "user",
			Method: http.MethodPost,
			Functions: []gin.HandlerFunc{
				CreateUserHandler(db),
			},
		},
		{
			Path:      "vote",
			Method:    http.MethodPost,
			Functions: []gin.HandlerFunc{},
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
