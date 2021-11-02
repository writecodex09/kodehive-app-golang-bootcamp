package main

import (
	"github.com/gin-gonic/gin"
	"github.com/writecodex09/kodehive-app-golang-bootcamp/controllers"
	"github.com/writecodex09/kodehive-app-golang-bootcamp/models"
	"net/http"
)

func main() {

	r := gin.Default()

	//call function SetupModels
	db := models.SetupModels()

	//use database in models folder
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	//first url [ / ]
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Message":"Restful API Kodehive"})
	})
	//get all mahasiswa
	r.GET("/mahasiswa", controllers.MahasiswaTampil)
	//membuat data baru mahasiswa
	r.POST("/mahasiswa", controllers.MahasiswaTambah)
	//mengupdate data mahasiswa
	r.PUT("/mahasiswa/:nim", controllers.MahasiswaUbah)
	//menghapus data mahasiswa
	r.DELETE("/mahasiswa/:nim", controllers.MahasiswaHapus)
	//run server default in port 8080
	r.Run()
}
