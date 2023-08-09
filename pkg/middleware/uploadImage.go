package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(r gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// var ctx = context.Background()

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Photo doesn't exist": err.Error(),
			})
			return
		}

		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error Open File": err.Error(),
			})
			return
		}

		defer src.Close()

		tempFile, err := os.CreateTemp("./upload", "image-*.jpg")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error Create Temp File": err.Error(),
			})
			return
		}

		defer tempFile.Close()

		if _, err := io.Copy(tempFile, src); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error Copy File": err.Error(),
			})
			return
		}

		ctx := tempFile.Name()

		c.Set("UploadedFile", ctx)

		r(c)
	}
}
