package database

import (
	"fmt"
	"selling/models"
	"selling/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		// &models.Project{},
		// &models.OnGoing{},
		// &models.Appsi{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}
	fmt.Println("Migration Success")
}
