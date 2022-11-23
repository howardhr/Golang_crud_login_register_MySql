package commons

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gomysql/user/models"
	"log"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@/test")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Executing Migration")
	db.AutoMigrate(&models.Usuario{})
}
