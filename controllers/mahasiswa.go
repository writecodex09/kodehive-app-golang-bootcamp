package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/writecodex09/kodehive-app-golang-bootcamp/models"
	"net/http"
)

type MahasiswaInput struct {
	gorm.Model
	Nim		string 	`json:"nim"`
	Nama 	string	`json:"nama"`
}

func MahasiswaTampil(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//array dari model mahasiswa
	var mhs []models.Mahasiswa
	db.Find(&mhs)
	c.JSON(http.StatusOK, gin.H{"Data":mhs})
}

func MahasiswaTambah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi inputan
	var dataInput MahasiswaInput
	//validasi agar supaya data yand diberikan dalam format json
	if err := c.ShouldBindJSON(&dataInput);
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//tambah validasi jika data nya tidak lengkap
	if dataInput.Nim == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Data NIM harus diisi"})
		return
	}

	//tambah validasi jika data nya tidak lengkap
	if dataInput.Nama == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Data Nama harus diisi"})
		return
	}

	//proses input data
	mhs := models.Mahasiswa{
		Nim:  dataInput.Nim,
		Nama: dataInput.Nama,
	}

	//proses membuat data
	db.Create(&mhs)
	//menampilkan data
	c.JSON(http.StatusOK, gin.H{"Data":mhs})
}