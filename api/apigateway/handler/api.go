package handler

import (
	api "album-management/api/apigateway/proto/api"
	"bufio"
	"io"
	"net/http"
	"strconv"

	"album-management/api/apigateway/util"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"gopkg.in/go-playground/validator.v9"
)

type Handler struct {
	AlbumService api.AlbumService
}

// CreateAlbum godoc
// @Tags album
// @Summary Create album
// @Description Create album to store the images
// @Accept  json
// @Produce  json
// @Param   albumId     path    int     true        "Parent Album ID"
// @Param   payload     body    albumReq     true        "Album data"
// @Success 200 {object} basicResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId} [post]
func (h *Handler) CreateAlbum(ctx *gin.Context) {
	log.Log("*************** CreateAlbum ************")
	var payload albumReq
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		log.Log("Error: " + err.Error())
		for _, fieldErr := range err.(validator.ValidationErrors) {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": util.ValidationErrorToText(fieldErr),
				"data":    nil,
			})
			return // exit on first error
		}
	}
	parentAlbumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	rsp, err := h.AlbumService.CreateAlbum(ctx, &api.Album{Name: payload.Name, ParentAlbum: parentAlbumID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": rsp.Msg,
		"data":    nil,
	})
}

// GetAlbum godoc
// @Tags album
// @Summary Get album
// @Description Get album by id
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Album ID"
// @Success 200 {object} albumResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/get [get]
func (h *Handler) GetAlbum(ctx *gin.Context) {
	albumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	rsp, err := h.AlbumService.GetAlbum(ctx, &api.GetRequest{Id: albumID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"data":    rsp.Album,
	})
}

// ListAlbum godoc
// @Tags album
// @Summary List album
// @Description List album in album
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Parent Album ID"
// @Param   offset     query    int64     true        "Offset"
// @Param   limit      query    int64     true        "Limit"
// @Success 200 {object} listAlbumResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/list [get]
func (h *Handler) ListAlbum(ctx *gin.Context) {
	parentAlbumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Param("limit"), 10, 64)
	offset, _ := strconv.ParseInt(ctx.Param("offset"), 10, 64)
	if limit == 0 {
		limit = 10
	}
	rsp, err := h.AlbumService.ListAlbum(ctx, &api.SearchRequest{AlbumId: parentAlbumID, Limit: limit, Offset: offset})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"data": map[string]interface{}{
			"list":  rsp.List,
			"count": rsp.Count,
		},
	})
}

// DeleteAlbum godoc
// @Tags album
// @Summary Get album
// @Description Get album by id
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Album ID"
// @Success 200 {object} basicResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId} [delete]
func (h *Handler) DeleteAlbum(ctx *gin.Context) {
	albumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	_, err := h.AlbumService.DeleteAlbum(ctx, &api.DeleteRequest{Id: albumID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Album successfully deleted",
		"data":    nil,
	})
}

// CreateImage godoc
// @Tags image
// @Summary Create Image
// @Description Create Image
// @Accept multipart/form-data
// @Produce  json
// @Param file formData file true  "images"
// @Param   albumId     path    int64     true        "Album ID"
// @Success 200 {object} basicResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/images [post]
func (h *Handler) CreateImage(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	defer file.Close()
	filename := header.Filename
	albumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	reader := bufio.NewReader(file)
	buffer := make([]byte, 2*1024)
	stream, err := h.AlbumService.CreateImage(ctx)
	req := &api.UploadImageRequest{
		Data: &api.UploadImageRequest_Info{
			Info: &api.ImageInfo{
				Name:    filename,
				AlbumId: albumID,
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	writing := true
	for writing {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			writing = false
		}
		if err != nil && err != io.EOF {
			log.Fatal("cannot send chunk to server: ", err)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, map[string]interface{}{
					"message": "Cannot send chunk to server",
					"data":    nil,
				})
				return
			}
		}

		req := &api.UploadImageRequest{
			Data: &api.UploadImageRequest_Chunk{
				Chunk: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
	}
	stream.Close()
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Image successfully uploaded",
		"data":    nil,
	})
}

// GetImage godoc
// @Tags image
// @Summary Get image
// @Description Get image by id
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Album ID"
// @Param   imageId     path    int64     true        "Image ID"
// @Success 200 {object} imageResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/images/{imageId}/get [get]
func (h *Handler) GetImage(ctx *gin.Context) {
	imageID, _ := strconv.ParseInt(ctx.Param("imageId"), 10, 64)
	rsp, err := h.AlbumService.GetImage(ctx, &api.GetRequest{Id: imageID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"data":    rsp.Image,
	})
}

// ListImage godoc
// @Tags image
// @Summary List image
// @Description List image in album
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Album ID"
// @Param   offset     query    int64     true        "Offset"
// @Param   limit      query    int64     true        "Limit"
// @Success 200 {object} listImageResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/images [get]
func (h *Handler) ListImage(ctx *gin.Context) {
	albumID, _ := strconv.ParseInt(ctx.Param("albumId"), 10, 64)
	limit, _ := strconv.ParseInt(ctx.Param("limit"), 10, 64)
	offset, _ := strconv.ParseInt(ctx.Param("offset"), 10, 64)
	if limit == 0 {
		limit = 10
	}
	rsp, err := h.AlbumService.ListImage(ctx, &api.SearchRequest{AlbumId: albumID, Limit: limit, Offset: offset})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "",
		"data": map[string]interface{}{
			"list":  rsp.List,
			"count": rsp.Count,
		},
	})
}

// DeleteImage godoc
// @Tags image
// @Summary Get image
// @Description Get image by id
// @Accept  json
// @Produce  json
// @Param   albumId     path    int64     true        "Album ID"
// @Param   imageId     path    int64     true        "Image ID"
// @Success 200 {object} basicResponse
// @Success 400 {object} basicResponse
// @Router /albums/{albumId}/images/{imageId} [delete]
func (h *Handler) DeleteImage(ctx *gin.Context) {
	imageID, _ := strconv.ParseInt(ctx.Param("imageId"), 10, 64)
	_, err := h.AlbumService.DeleteImage(ctx, &api.DeleteRequest{Id: imageID})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Image successfully deleted",
		"data":    nil,
	})
}
