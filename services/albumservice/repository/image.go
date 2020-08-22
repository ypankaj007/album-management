package repository

import (
	album "album-management/services/albumservice/proto/album"
	"context"
	"database/sql"
	"strings"
)

// ImageRepository godoc
type ImageRepository interface {
	Create(context.Context, *album.Image) (int64, error)
	Delete(context.Context, int64) (bool, error)
	Get(context.Context, int64) (*album.Image, error)
	List(context.Context, int64, int64, int64) ([]*album.Image, int64, error)
	IsExists(context.Context, int64) (bool, error)
	DeleteInAlbum(context.Context, []interface{}) error
}

type imageRepository struct {
	DB *sql.DB
}

func NewImageRepositoy(db *sql.DB) ImageRepository {
	return &imageRepository{
		DB: db,
	}
}

func (img *imageRepository) Create(ctx context.Context, image *album.Image) (int64, error) {

	statement, err := img.DB.Prepare("INSERT INTO images (name, albumId, url, createdAt, updatedAt, isDeleted) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	defer statement.Close()
	res, err := statement.Exec(image.Name, image.AlbumId, image.Url, image.CreatedAt, image.UpdatedAt, false)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (img *imageRepository) Delete(ctx context.Context, id int64) (bool, error) {
	statement, err := img.DB.Prepare("UPDATE images SET isDeleted = 1 WHERE id = ?")
	if err != nil {
		return false, err
	}
	defer statement.Close()
	statement.Exec(id)
	return true, nil
}
func (img *imageRepository) IsExists(ctx context.Context, id int64) (bool, error) {
	_, err := img.Get(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (img *imageRepository) Get(ctx context.Context, id int64) (*album.Image, error) {
	var image album.Image
	err := img.DB.QueryRow("SELECT id, name, albumId, createdAt, updatedAt, url FROM images WHERE id = ? AND isDeleted = 0", id).Scan(&image.Id, &image.Name, &image.AlbumId, &image.CreatedAt, &image.UpdatedAt, &image.Url)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func (img *imageRepository) List(ctx context.Context, albumID int64, limit, offset int64) ([]*album.Image, int64, error) {
	list := make(chan []*album.Image)
	count := make(chan int64)
	go img.list(ctx, albumID, limit, offset, list)
	go img.count(ctx, albumID, count)
	l := <-list
	c := <-count
	return l, c, nil
}

func (img *imageRepository) list(ctx context.Context, albumID int64, limit, offset int64, result chan []*album.Image) {
	rows, err := img.DB.Query("SELECT id, name, albumId, createdAt, updatedAt, url FROM images WHERE albumId = ? AND isDeleted = 0 LIMIT ? OFFSET ?", albumID, limit, offset)
	if err != nil {
		result <- nil
	}
	var list []*album.Image
	for rows.Next() {
		var image album.Image
		rows.Scan(&image.Id, &image.Name, &image.AlbumId, &image.CreatedAt, &image.UpdatedAt, &image.Url)
		list = append(list, &image)
	}
	result <- list
}

func (img *imageRepository) count(ctx context.Context, albumID int64, result chan int64) {

	rows, err := img.DB.Query("SELECT count(*) FROM images WHERE albumId = ? AND isDeleted = 0", albumID)
	if err != nil {
		result <- 0
	}
	var count int64
	for rows.Next() {
		rows.Scan(&count)
	}
	result <- count
}

func (img *imageRepository) DeleteInAlbum(ctx context.Context, ids []interface{}) error {
	query := `UPDATE images SET isDeleted = 1 WHERE albumId in (?` + strings.Repeat(",?", len(ids)-1) + `)`
	statement, err := img.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	statement.Exec(ids...)
	return nil
}
