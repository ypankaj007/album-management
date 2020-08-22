package handler

type albumReq struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=50" example:"My Collection"`
}

type basicResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type albumResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID          int64  `json:"id"`
		ParentAlbum int64  `json:"ParentAlbumId"`
		Name        string `json:"name"`
		Path        string `json:"path"`
		CreatedAT   int64  `json:"createdAt"`
		UpdatedAT   int64  `json:"updatedAt"`
	} `json:"data"`
}
type imageResponse struct {
	Message string `json:"message"`
	Data    struct {
		ID        int64  `json:"id"`
		AlbumID   int64  `json:"albumId"`
		Name      string `json:"name"`
		URL       string `json:"url"`
		CreatedAT int64  `json:"createdAt"`
		UpdatedAT int64  `json:"updatedAt"`
	} `json:"data"`
}
type listAlbumResponse struct {
	Message string `json:"message"`
	Data    struct {
		List []struct {
			ID          int64  `json:"id"`
			ParentAlbum int64  `json:"ParentAlbumId"`
			Name        string `json:"name"`
			Path        string `json:"path"`
			CreatedAT   int64  `json:"createdAt"`
			UpdatedAT   int64  `json:"updatedAt"`
		} `json:"list"`
		Count int64 `json:"count"`
	} `json:"data"`
}
type listImageResponse struct {
	Message string `json:"message"`
	Data    struct {
		List []struct {
			ID        int64  `json:"id"`
			AlbumID   int64  `json:"albumId"`
			Name      string `json:"name"`
			URL       string `json:"url"`
			CreatedAT int64  `json:"createdAt"`
			UpdatedAT int64  `json:"updatedAt"`
		} `json:"list"`
		Count int64 `json:"count"`
	} `json:"data"`
}
