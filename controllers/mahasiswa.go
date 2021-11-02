package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/writecodex09/kodehive-app-golang-bootcamp/models"
	"net/http"
	"time"
)

type MahasiswaInput struct {
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

	//tambahkan validasi jika jumlah karakter kurang dari 5
	if len(dataInput.Nama) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Nama harus lebih dari 5 karakter"})
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

func MahasiswaUbah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek apakah data tersebut ada atau tidak didalam database
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error;
	err  != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Data tidak ditemukan"})
		return
	}

	//validasi inputan
	var dataInput MahasiswaInput
	//validasi agar supaya data yand diberikan dalam format json
	if err := c.ShouldBindJSON(&dataInput);
		err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	//proses ubah data
	db.Model(&mhs).Update(dataInput)
	//menampilkan data
	c.JSON(http.StatusOK, gin.H{"Data":mhs, "message":"berhasil mengubah data", "time":time.Now()})
}

func MahasiswaHapus(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//cek apakah data tersebut ada atau tidak didalam database
	var mhs models.Mahasiswa
	if err := db.Where("nim = ?", c.Param("nim")).First(&mhs).Error;
		err  != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Data tidak ditemukan"})
		return
	}

	//proses hapus data
	db.Delete(&mhs)
	//menampilkan data
	c.JSON(http.StatusOK, gin.H{"message":"berhasil menghapus data"})
}

