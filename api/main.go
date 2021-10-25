package api

import (
	_ "github.com/Yusuf1901-cloud/message_sender/api/docs"
	"github.com/Yusuf1901-cloud/message_sender/api/handler"
	"github.com/Yusuf1901-cloud/message_sender/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	HttpPort string
	Engine   *gin.Engine
}

func New(cfg config.Config) Server {
	engine := gin.Default()

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	engine.POST("/send", handler.SendMessage)

	return Server{
		HttpPort: cfg.HttpPort,
		Engine:   engine,
	}
}

func (s Server) Run() error {
	return s.Engine.Run(s.HttpPort)
}
