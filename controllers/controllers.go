package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Arturo0911/go-android/connection"
	"github.com/Arturo0911/go-android/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JsonResponse struct {
	Success bool
	Error   string
}

type StructRepo struct {
	Db *gorm.DB
}

type MeasurementsResponse struct {
	IdLevels       int64     `json:"id_niveles"`
	IdReaders      int64     `json:"id_lecturaN"`
	DateWrite      time.Time `json:"fecha_lectura"`
	OxigenMin      float64   `json:"oxigen_min_value"`
	OxigenMax      float64   `json:"oxigen_max_value"`
	TemperatureMin float64   `json:"temperature_min_value"`
	TemperatureMax float64   `json:"temperature_max_value"`
	HumidityMin    float64   `json:"humidity_min_value"`
	HumidityMax    float64   `json:"humidity_max_value"`
	DioxideMin     float64   `json:"dioxide_min_value"`
	DioxideMax     float64   `json:"dioxide_max_value"`
}

func New() *StructRepo {
	db := connection.NewInstance()
	db.AutoMigrate(&models.UsersValues{})
	return &StructRepo{
		Db: db,
	}
}

func GetMeasurements(c *gin.Context) {
	c.JSON(200, MeasurementsResponse{
		OxigenMin: 25.22,
	})
}

func (repository *StructRepo) PostUsersControllers(c *gin.Context) {
	var reading models.UsersValues
	c.ShouldBindJSON(&reading)
	//reading.DateReading = time.Now().Truncate()
	err := models.CreateUsers(repository.Db, &reading)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	fmt.Println(&reading)
	c.JSON(http.StatusOK, reading)
}

func (repository *StructRepo) GetUsersControllers(c *gin.Context) {
	var reading []models.UsersValues
	err := models.GetUsers(repository.Db, &reading)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error": err,
			})
		return
	}
	c.JSON(http.StatusOK, reading)
}
