package main

import (
	"time"

	_ "album-management/api/apigateway/docs"
	"album-management/api/apigateway/handler"
	api "album-management/api/apigateway/proto/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Album Management
// @version 1.0
// @description Album and image management apis

// @contact.name API Support
// @contact.email ypankaj007@gmail.com

// @BasePath /api/v1
func main() {
	// New Service
	service := web.NewService(
		web.Name("go.micro.srv.api"),
		web.Address(":8080"),
	)

	// Initialise service
	service.Init()

	// Register Service
	albumSevice := api.NewAlbumService("go.micro.srv.album", client.DefaultClient)
	controller := &handler.Handler{AlbumService: albumSevice}
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.MaxMultipartMemory = 30 << 20 // 8 MiB

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Static("/static", "/var/db")
	v1 := r.Group("/api/v1")
	{
		v1.POST("/albums/:albumId", controller.CreateAlbum)
		v1.GET("/albums/:albumId/list", controller.ListAlbum)
		v1.GET("/albums/:albumId/get", controller.GetAlbum)
		v1.DELETE("/albums/:albumId", controller.DeleteAlbum)
		v1.POST("/albums/:albumId/images", controller.CreateImage)
		v1.DELETE("/albums/:albumId/images/:imageId", controller.DeleteImage)
		v1.GET("/albums/:albumId/images/:imageId/get", controller.GetImage)
		v1.GET("/albums/:albumId/images", controller.ListImage)
	}

	service.Handle("/", r)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
