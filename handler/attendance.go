package handler

import (
	dto "TestBE/dto/result"
	"TestBE/models"
	"TestBE/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Kagami/go-face"
	"github.com/gin-gonic/gin"
)

const dataDir = "handler"

type handlerAttendance struct {
	AttendanceRepository repository.AttendanceRepository
}

func HandlerAttendance(AttendanceRepository repository.AttendanceRepository) *handlerAttendance {
	return &handlerAttendance{AttendanceRepository}
}

func (h *handlerAttendance) FindAll(c *gin.Context) {
	attendance, err := h.AttendanceRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{
		Code: http.StatusOK,
		Data: attendance,
	})
}

func (h *handlerAttendance) CreateAttendanceCheckIn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	StatusCheckIn := "CheckIn"
	PhotoDatabaseEmployee, _ := h.AttendanceRepository.GetByID(id)
	PhotoUploadEmployee, _ := c.Get("UploadedCheckerFile")

	// fmt.Println(PhotoDatabaseEmployee)

	currentTime := time.Now()
	hour := currentTime.Hour()
	minute := currentTime.Minute()

	var category string

	//Start Face Recognition
	fmt.Println("Facial Recognition System v0.01")

	reccc := filepath.Join(dataDir, "models")
	rec, err := face.NewRecognizer(reccc)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")

	faces, err := rec.RecognizeFile(PhotoDatabaseEmployee.Picture)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	// Pass samples to the recognizer.
	rec.SetSamples([]face.Descriptor{faces[0].Descriptor}, []int32{0})

	// Now let's try to classify some not yet known image.
	tonyStark, err := rec.RecognizeFile(PhotoUploadEmployee.(string))
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if tonyStark == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Face Not Found"})
		os.Remove(PhotoUploadEmployee.(string))
	}

	descriptor1 := faces[0].Descriptor[:]     // Foto wajah pertama (vektor fitur)
	descriptor2 := tonyStark[0].Descriptor[:] // Foto wajah kedua (vektor fitur)

	similarity := compareFaces(descriptor1, descriptor2)
	fmt.Printf("Similarity: %.2f\n", (100 - similarity*100))

	// os.Remove(PhotoUploadEmployee.(string)) ((Cadangan))
	//End face recoginition

	if hour < 8 {
		category = "On-Time"
	} else if hour >= 8 && minute > 00 && hour < 17 {
		category = "Terlambat"
	} else if hour == 8 && minute == 00 {
		category = "On-Time"
	}

	attendance := models.Attendance{
		Date:       time.Now().Format("2006-01-02"),
		CheckIn:    time.Now().Format("15:04"),
		Status:     StatusCheckIn,
		StatusNote: category,
		EmployeeId: id,
	}

	if (100 - similarity*100) < 80 {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Please Re-Take Photo"})
		os.Remove(PhotoUploadEmployee.(string))
	} else if (100 - similarity*100) > 80 {
		_, ErrData := h.AttendanceRepository.CreateAttendanceCheckIn(attendance, id, StatusCheckIn)
		if ErrData != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: ErrData.Error()})
		} else {
			c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: attendance})
		}
		os.Remove(PhotoUploadEmployee.(string))
	}
}

func (h *handlerAttendance) CreateAttendanceCheckOut(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	StatusCheckOut := "CheckOut"
	PhotoDatabaseEmployee, _ := h.AttendanceRepository.GetByID(id)
	PhotoUploadEmployee, _ := c.Get("UploadedCheckerFile")

	currentTime := time.Now()
	hour := currentTime.Hour()
	minute := currentTime.Minute()

	// fmt.Println("Jam :", hour, minute)

	var category string

	//Start Face Recognition
	fmt.Println("Facial Recognition System v0.01")

	reccc := filepath.Join(dataDir, "models")
	rec, err := face.NewRecognizer(reccc)
	if err != nil {
		fmt.Println("Cannot initialize recognizer")
	}
	defer rec.Close()

	fmt.Println("Recognizer Initialized")

	faces, err := rec.RecognizeFile(PhotoDatabaseEmployee.Picture)
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}

	// Pass samples to the recognizer.
	rec.SetSamples([]face.Descriptor{faces[0].Descriptor}, []int32{0})

	// Now let's try to classify some not yet known image.
	tonyStark, err := rec.RecognizeFile(PhotoUploadEmployee.(string))
	if err != nil {
		log.Fatalf("Can't recognize: %v", err)
	}
	if tonyStark == nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Face Not Found"})
		os.Remove(PhotoUploadEmployee.(string))
	}

	descriptor1 := faces[0].Descriptor[:]     // Foto wajah pertama (vektor fitur)
	descriptor2 := tonyStark[0].Descriptor[:] // Foto wajah kedua (vektor fitur)

	similarity := compareFaces(descriptor1, descriptor2)
	fmt.Printf("Similarity: %.2f\n", (100 - similarity*100))

	// os.Remove(PhotoUploadEmployee.(string)) ((Cadangan))
	//End face recoginition

	if hour > 17 {
		category = "On-Time"
	} else if hour >= 8 && minute > 00 && hour < 17 {
		category = "Pulang Cepat"
	} else if hour == 17 && minute == 00 {
		category = "On-Time"
	}

	attendance := models.Attendance{
		Date:       time.Now().Format("2006-01-02"),
		CheckOut:   time.Now().Format("15:04"),
		Status:     StatusCheckOut,
		StatusNote: category,
		EmployeeId: id,
	}

	if (100 - similarity*100) < 80 {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Please Re-Take Photo"})
		os.Remove(PhotoUploadEmployee.(string))
	} else if (100 - similarity*100) > 80 {
		_, ErrData := h.AttendanceRepository.CreateAttendanceCheckOut(attendance, id, StatusCheckOut)
		if ErrData != nil {
			c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: ErrData.Error()})
		} else {
			c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: attendance})
		}
		os.Remove(PhotoUploadEmployee.(string))
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: attendance})
}

func compareFaces(descriptor1 []float32, descriptor2 []float32) float64 {
	if len(descriptor1) != len(descriptor2) {
		return -1.0 // Panjang vektor fitur tidak cocok
	}

	var sum float64
	for i := 0; i < len(descriptor1); i++ {
		diff := float64(descriptor1[i]) - float64(descriptor2[i])
		sum += diff * diff
	}

	return sum
}
