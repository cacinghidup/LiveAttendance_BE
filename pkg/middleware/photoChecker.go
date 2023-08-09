package middleware

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PhotoChecker(n gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {

		file, err := c.FormFile("checker")
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

		tempFile, err := os.CreateTemp("./uploadChecker", "image-*.jpg")
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

		c.Set("UploadedCheckerFile", ctx)

		n(c)
	}
}

// func (r *checkerByFile) VerificationPicture(c *gin.Context, Id int) {
// 	var Employee models.Employee

// 	err := r.db.First(&Employee, Id).Error
// 	if err != nil {
// 		panic(err)
// 	}

// 	PictEmployee := Employee.Picture

// 	rec, err := face.NewRecognizer(modelsDir)
// 	if err != nil {
// 		log.Fatalf("Error saat membuat recognizer: %v", err)
// 	}
// 	defer rec.Close()

// 	// Load gambar pertama
// 	img1Path := PictEmployee
// 	img1, err := rec.RecognizeSingleFile(img1Path)
// 	if err != nil {
// 		log.Fatalf("Error saat mengenali wajah gambar 1: %v", err)
// 	}

// 	// Load gambar kedua
// 	test, _ := c.Get("UploadedCheckerFile")
// 	testus := test.(string)

// 	img2, err := rec.RecognizeSingleFile(testus)
// 	if err != nil {
// 		log.Fatalf("Error saat mengenali wajah gambar 2: %v", err)
// 	}

// 	// Perbandingan kesamaan wajah
// 	match := rec.
// 		fmt.Printf("Similarity: %.2f%%\n", (1.0-match)*100)
// }
