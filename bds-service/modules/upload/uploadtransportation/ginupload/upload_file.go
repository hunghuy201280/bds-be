package ginupload

import (
	"bds-service/common"
	"bds-service/component"
	"bds-service/modules/upload/uploadbiz"
	"bds-service/modules/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func Upload(ctx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		//err = c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename))
		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(common.ErrInvalidRequest(err))
			}
		}(file)

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		imgStore := uploadstorage.NewSQLStore(ctx.GetMainDbConnection())

		biz := uploadbiz.NewUploadBiz(ctx.GetUploadProvider(), imgStore)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
