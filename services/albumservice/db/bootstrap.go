package db

import (
	"album-management/services/albumservice/util"
	"database/sql"
	"log"
)

// CreateTables godoc
func CreateTables(database *sql.DB) {
	statement, err := database.Prepare(`CREATE TABLE IF NOT EXISTS albums(
		id INT PRIMARY KEY AUTO_INCREMENT,
		path TEXT,
		name TEXT,
		parentAlbumId INT,
		isDeleted BOOLEAN,
		createdAt BIGINT,
		updatedAt BIGINT) ENGINE=InnoDB DEFAULT CHARSET=latin1`)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	statement, err = database.Prepare(`CREATE TABLE IF NOT EXISTS images (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name TEXT,
		url TEXT,
		albumId INT,
		isDeleted BOOLEAN,
		createdAt BIGINT,
		updatedAt BIGINT) ENGINE=InnoDB DEFAULT CHARSET=latin1`)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	//initilaizeRootFolder(database)
}

func initilaizeRootFolder(database *sql.DB) {
	if isRootExists(database) {
		return
	}
	statement, err := database.Prepare("INSERT INTO albums (name, parentAlbumId, path, isDeleted, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("root", 0, "/", 0, util.CurrentTimeInMilli(), util.CurrentTimeInMilli())
}

func isRootExists(database *sql.DB) bool {
	var id int32
	err := database.QueryRow("SELECT id FROM albums WHERE parentAlbumId = ? ", 0).Scan(&id)
	if err != nil {
		return false
	}
	return id > 0
}
