package main

import (
	//"album-management/services/albumservice/config"
	"album-management/services/albumservice/db"
	"album-management/services/albumservice/handler"
	"album-management/services/albumservice/repository"
	srv "album-management/services/albumservice/service"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	album "album-management/services/albumservice/proto/album"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.album"),
		micro.Version("1.0.0"),
	)

	// Initialise service
	service.Init()

	// Get instance of the broker using our defaults
	pubSub := srv.NewPublisher()

	// connect db
	//database, err := sql.Open("sqlite3", config.GetENV("DB_PATH"))
	database, err := sql.Open("mysql", "album:album@tcp(mysql_db:1306)/album")

	if err != nil {
		log.Log(err.Error())
	}

	db.CreateTables(database)

	// repositories
	albumRepo := repository.NewAlbumRepositoy(database)
	imageRepo := repository.NewImageRepositoy(database)

	// Register Handler
	album.RegisterAlbumServiceHandler(service.Server(), &handler.AlbumHandler{
		ImageRepository: imageRepo,
		AlbumRepository: albumRepo,
		Publisher:       pubSub,
	})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
