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

	//grouping router
	v1 := r.Group("/api/v1/")
	{
		//first url [ / ]
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"Message":"Restful API Kodehive"})
		})
		//get all mahasiswa
		v1.GET("/mahasiswa", controllers.MahasiswaTampil)
		//membuat data baru mahasiswa
		v1.POST("/mahasiswa", controllers.MahasiswaTambah)
		//mengupdate data mahasiswa
		v1.PUT("/mahasiswa/:nim", controllers.MahasiswaUbah)
		//menghapus data mahasiswa
		v1.DELETE("/mahasiswa/:nim", controllers.MahasiswaHapus)
		//run server default in port 8080
	}

	r.Run()
}
