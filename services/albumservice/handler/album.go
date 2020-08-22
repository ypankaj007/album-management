package handler

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"album-management/services/albumservice/service"
	"album-management/services/albumservice/util"

	"github.com/micro/go-micro/util/log"

	album "album-management/services/albumservice/proto/album"
	repository "album-management/services/albumservice/repository"
)

type AlbumHandler struct {
	AlbumRepository repository.AlbumRepository
	ImageRepository repository.ImageRepository
	Publisher       *service.PubSub
}

func (handler *AlbumHandler) CreateAlbum(ctx context.Context, req *album.Album, rsp *album.Response) error {
	log.Log("Received Album.CreateAlbum request")
	data := &album.Album{
		Name:        req.Name,
		ParentAlbum: req.ParentAlbum,
	}
	if data.ParentAlbum == 0 {
		data.Path = "/"
	} else {
		parent, err := handler.AlbumRepository.Get(ctx, data.ParentAlbum)
		if err != nil {
			log.Error(err)
			return err
		}
		if parent.ParentAlbum == 0 {
			data.Path = fmt.Sprintf("%s%d", parent.Path, parent.Id)
		} else {
			data.Path = fmt.Sprintf("%s/%d", parent.Path, parent.Id)
		}
	}
	data.CreatedAt = util.CurrentTimeInMilli()
	data.UpdatedAt = util.CurrentTimeInMilli()
	id, err := handler.AlbumRepository.Create(ctx, data)
	if err == nil {
		path := fmt.Sprintf("%s/%d", data.Path, id)
		if data.ParentAlbum == 0 {
			path = fmt.Sprintf("%s%d", data.Path, id)
		}
		err = service.NewDisk().CreateFolder(path)
		if err != nil {
			log.Error(err)
			go handler.AlbumRepository.Delete(ctx, id)
			return err
		}
		handler.Publisher.Log(fmt.Sprintf("%s '%s' %s", "Album", data.Name, "created"))
		rsp.Msg = "Album successfully created"
	}
	return err
}

// CreateImage ..
func (handler *AlbumHandler) CreateImage(ctx context.Context, stream album.AlbumService_CreateImageStream) error {
	log.Log("Received Album.CreateImage request")
	var fileName string
	var albumID int64
	imageData := bytes.Buffer{}
	for {
		req, err := stream.Recv()
		if err != nil {
			log.Error(err.Error())
			if err == io.EOF {
				break
			}
			return err
		}
		chunk := req.GetChunk()
		if chunk != nil {
			log.Log("Received chunk")
			_, err = imageData.Write(chunk)
		}
		imgInf := req.GetInfo()
		if imgInf != nil {
			log.Log("Received image info")
			fileName = imgInf.Name
			albumID = imgInf.AlbumId
		}

	}
	path := ""
	var alm *album.Album
	var err error
	if albumID != 0 {
		alm, err = handler.AlbumRepository.Get(ctx, albumID)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	if albumID == 0 {
		path = ""
	} else if alm.ParentAlbum == 0 {
		path = fmt.Sprintf("%s%d", alm.Path, albumID)
	} else {
		path = fmt.Sprintf("%s/%d", alm.Path, albumID)
	}
	log.Log("uploading....")
	imageID, err := service.NewDisk().Save(path, imageData)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Log("DONE")
	_, err = handler.ImageRepository.Create(ctx, &album.Image{
		Name:      fileName,
		Url:       fmt.Sprintf("%s/%s", path, imageID),
		AlbumId:   albumID,
		CreatedAt: util.CurrentTimeInMilli(),
		UpdatedAt: util.CurrentTimeInMilli(),
	})
	handler.Publisher.Log(fmt.Sprintf("%s '%s' %s", "Image", fileName, "uploaded"))
	return err
}

func (handler *AlbumHandler) DeleteAlbum(ctx context.Context, req *album.DeleteRequest, rsp *album.Response) error {
	log.Log("Received Album.DeleteAlbum request")
	parent, err := handler.AlbumRepository.Get(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	ids, err := handler.AlbumRepository.Delete(ctx, req.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	go handler.ImageRepository.DeleteInAlbum(ctx, ids)
	if parent.ParentAlbum == 0 {
		service.NewDisk().Remove(fmt.Sprintf("%s%d", parent.Path, req.Id))
	} else {
		service.NewDisk().Remove(fmt.Sprintf("%s/%d", parent.Path, req.Id))
	}
	handler.Publisher.Log("Album deleted")
	return nil
}

func (handler *AlbumHandler) DeleteImage(ctx context.Context, req *album.DeleteRequest, rsp *album.Response) error {
	log.Log("Received Album.Call request")
	image, err := handler.ImageRepository.Get(ctx, req.Id)
	if err != nil {
		return err
	}
	result, err := handler.ImageRepository.Delete(ctx, req.Id)
	if err != nil {
		return err
	}
	if result {
		go service.NewDisk().Remove(image.Url)
	}
	handler.Publisher.Log("Image deleted")
	return nil
}
func (handler *AlbumHandler) GetAlbum(ctx context.Context, req *album.GetRequest, rsp *album.AlbumResponse) error {
	log.Log("Received Album.GetAlbum request")
	album, err := handler.AlbumRepository.Get(ctx, req.Id)
	rsp.Album = album
	return err
}
func (handler *AlbumHandler) GetImage(ctx context.Context, req *album.GetRequest, rsp *album.ImageResponse) error {
	log.Log("Received Album.GetImage request")
	image, err := handler.ImageRepository.Get(ctx, req.Id)
	rsp.Image = image
	return err
}

func (handler *AlbumHandler) ListAlbum(ctx context.Context, req *album.SearchRequest, rsp *album.Albums) error {
	log.Log("Received Album.ListAlbum request")
	list, count, err := handler.AlbumRepository.List(ctx, req.AlbumId, req.Limit, req.Offset)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.List = list
	rsp.Count = count
	return err
}
func (handler *AlbumHandler) ListImage(ctx context.Context, req *album.SearchRequest, rsp *album.Images) error {
	log.Log("Received Album.ListImage request")
	list, count, err := handler.ImageRepository.List(ctx, req.AlbumId, req.Limit, req.Offset)
	if err != nil {
		log.Error(err)
		return err
	}
	rsp.List = list
	rsp.Count = count
	return err
}
