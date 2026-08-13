package main

import (
	"fmt"

	"school-management/config"
	"school-management/models"
	"school-management/seed"
)

func main() {
	fmt.Println("Menjalankan Backend LMS...")

	config.ConnectDatabase()

	fmt.Println("Menjalankan AutoMigrate...")

	err := config.DB.AutoMigrate(
		&models.User{},
		&models.EducationLevel{},
		&models.Class{},
		&models.Subject{},
	)

	if err != nil {
		fmt.Println("Gagal melakukan AutoMigrate:", err)
		return
	}

	fmt.Println("AutoMigrate berhasil!")

	fmt.Println("Menjalankan Seeder...")

	seed.SeedEducationLevels()
	seed.SeedClasses()

	fmt.Println("Seeder berhasil dijalankan!")
	fmt.Println("Backend LMS berhasil dijalankan!")
}