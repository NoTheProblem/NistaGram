package main

import (
	"auth-service/handler"
	"auth-service/model"
	"auth-service/repository"
	"auth-service/service"
	"fmt"
	"github.com/gorilla/mux"
	cors "github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"auth-service","test","auth-db","5432")
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
	router.HandleFunc("/authorize", handler.Authorize).Methods("GET")

	c := SetupCors()

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}

}

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}

