package main

import (
	"auth-service/handler"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"postgres","1234567","auth-service","5432")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.User{})
	return database
}

func initRepo(database *gorm.DB) *repository.AuthRepository {
	return &repository.AuthRepository{Database: database}
}

func initServices(repository *repository.AuthRepository) *service.AuthService {
	return &service.AuthService{AuthRepository: repository}
}
func initHandler(service *service.AuthService) *handler.AuthHandler {
	return &handler.AuthHandler{AuthService: service}
}
func handleFunc(handler *handler.AuthHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handler.Hello).Methods("GET")
	router.HandleFunc("/register", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/update", handler.UpdateUser).Methods("POST")
	router.HandleFunc("/passwordChange", handler.PasswordChange).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))

}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}
