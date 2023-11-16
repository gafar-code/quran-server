package api

import (
	"fmt"
	"log"
	"time"

	db "github.com/gafar-code/quran-server/db/sqlc"
	"github.com/gafar-code/quran-server/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	page   *db.Page
	router *gin.Engine
}

type ResponseData struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseErr struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func errorHandler(c *gin.Context, err error, code int) {
	if err != nil {
		c.JSON(code, ResponseErr{
			Code:    int32(code),
			Message: err.Error(),
		})
	}
}

func NewServer(page *db.Page) *Server {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Tidak dapat memuat config:", err)
	}

	router := gin.Default()
	server := &Server{page: page}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "api_key", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	wrapper := &WrapperServer{
		server: server,
	}

	middleware := []MiddlewareFunc{
		// AuthMiddleware,
	}

	baseUrl := fmt.Sprintf("/api/v%v", config.Version)

	RegisterHandlersWithOptions(router, wrapper, GinServerOptions{
		BaseURL:      baseUrl,
		Middlewares:  middleware,
		ErrorHandler: errorHandler,
	})

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
