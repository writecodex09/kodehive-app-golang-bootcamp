package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	//connect to database
	db, err := gorm.Open("mysql", "root:root@/kodehivedb?charset=utf8&parseTime=True&loc=Local")
	//check database connection
	if err != nil {
		panic("cannot connect to database kodehive")
	}
	//migrate Mahasiswa Tables
	db.AutoMigrate(&Mahasiswa{})

	return db
}
