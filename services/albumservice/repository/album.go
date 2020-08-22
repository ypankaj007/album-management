package repository

import (
	album "album-management/services/albumservice/proto/album"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/micro/go-micro/util/log"
)

type AlbumRepository interface {
	Create(context.Context, *album.Album) (int64, error)
	Delete(context.Context, int64) ([]interface{}, error)
	Get(context.Context, int64) (*album.Album, error)
	List(context.Context, int64, int64, int64) ([]*album.Album, int64, error)
	IsExists(context.Context, int64) (bool, error)
}

type albumRepository struct {
	DB *sql.DB
}

// NewAlbumRepositoy initialize repo
func NewAlbumRepositoy(db *sql.DB) AlbumRepository {
	return &albumRepository{
		DB: db,
	}
}

func (alb *albumRepository) Create(ctx context.Context, album *album.Album) (int64, error) {
	// parent, err := alb.Get(ctx, album.ParentAlbum)
	// if err != nil {
	// 	return err
	// }
	// album.Path = fmt.Sprintf("%s/%d", parent.Path, parent.Id)
	statement, err := alb.DB.Prepare("INSERT INTO albums (name, parentAlbumId, path, isDeleted, createdAt, updatedAt) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}
	fmt.Println("--------------")
	fmt.Println(fmt.Sprintf("%v", album))
	fmt.Println(album.CreatedAt)
	fmt.Println("--------------")
	defer statement.Close()
	res, err := statement.Exec(album.Name, album.ParentAlbum, album.Path, false, album.CreatedAt, album.UpdatedAt)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}
func (alb *albumRepository) Delete(ctx context.Context, id int64) ([]interface{}, error) {
	return alb.deletedAllAlbumOnPathAndReturnIds(ctx, id)
}
func (alb *albumRepository) IsExists(ctx context.Context, id int64) (bool, error) {
	_, err := alb.Get(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (alb *albumRepository) Get(ctx context.Context, id int64) (*album.Album, error) {
	var result album.Album
	err := alb.DB.QueryRow("SELECT id, name, parentAlbumId, path, createdAt, updatedAt FROM albums WHERE id = ? AND isDeleted = 0", id).Scan(&result.Id, &result.Name, &result.ParentAlbum, &result.Path, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &result, nil
}

func (alb *albumRepository) List(ctx context.Context, parentAlbumID, limit, offset int64) ([]*album.Album, int64, error) {
	list := make(chan []*album.Album)
	count := make(chan int64)
	go alb.list(ctx, parentAlbumID, limit, offset, list)
	go alb.count(ctx, parentAlbumID, count)
	l := <-list
	c := <-count
	fmt.Println(l, c)
	return l, c, nil
}

func (alb *albumRepository) list(ctx context.Context, parentAlbumID, limit, offset int64, result chan []*album.Album) {
	log.Log("List func()", offset, limit)
	rows, err := alb.DB.Query("SELECT id, name, parentAlbumId, path, createdAt, updatedAt FROM albums WHERE parentAlbumId = ? AND isDeleted = 0 LIMIT ? OFFSET ?", parentAlbumID, limit, offset)
	if err != nil {
		log.Log("List func() error")
		result <- make([]*album.Album, 0)
	}
	var list []*album.Album
	for rows.Next() {
		var alm album.Album
		rows.Scan(&alm.Id, &alm.Name, &alm.ParentAlbum, &alm.Path, &alm.CreatedAt, &alm.UpdatedAt)
		list = append(list, &alm)
	}
	log.Log("List ch done")
	result <- list
}

func (alb *albumRepository) count(ctx context.Context, parentAlbumID int64, result chan int64) {
	log.Log("Count func()")
	rows, err := alb.DB.Query("SELECT count(*) FROM albums WHERE parentAlbumId = ? AND isDeleted = 0", parentAlbumID)
	if err != nil {
		log.Log("Count func() error")
		result <- 0
	}
	var count int64
	for rows.Next() {
		rows.Scan(&count)
	}
	log.Log("Count ch done")
	result <- count
}

func (alb *albumRepository) deletedAllAlbumOnPathAndReturnIds(ctx context.Context, id int64) ([]interface{}, error) {
	rows, err := alb.DB.Query("SELECT id from albums WHERE path LIKE ? OR path LIKE ? AND isDeleted = ?", fmt.Sprintf("%s%d", "%/", id), fmt.Sprintf("%s%d%s", "%/", id, "/%"), 0)
	if err != nil {
		return nil, err
	}
	var ids []interface{}
	for rows.Next() {
		var childID int64
		rows.Scan(&childID)
		ids = append(ids, childID)
	}
	ids = append(ids, id)
	stmt := `UPDATE albums SET isDeleted = 1 WHERE id in (?` + strings.Repeat(",?", len(ids)-1) + `)`
	statement, err := alb.DB.Prepare(stmt)
	if err != nil {
		return nil, err
	}
	defer statement.Close()
	statement.Exec(ids...)

	return ids, nil
}
