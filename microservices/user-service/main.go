package main

import (
	"fmt"
	"github.com/gorilla/mux"
	cors "github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"user-service/handler"
	"user-service/model"
	"user-service/repository"
	"user-service/service"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"user-service","test","user-db","5432")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.User{})
	return database
}

func initRepo(database *gorm.DB) *repository.UserRepository {
	return &repository.UserRepository{Database: database}
}

func initServices(repository *repository.UserRepository) *service.UserService {
	return &service.UserService{UserRepository: repository}
}
func initHandler(service *service.UserService) *handler.UserHandler {
	return &handler.UserHandler{UserService: service}
}
func handleFunc(handler *handler.UserHandler) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/registerUser", handler.RegisterUser).Methods("POST")
	router.HandleFunc("/updateProfileInfo", handler.UpdateProfileInfo).Methods("POST")
	router.HandleFunc("/updateNotificationSettings", handler.UpdateNotificationSettings).Methods("POST")
	router.HandleFunc("/updatePrivacySettings", handler.UpdatePrivacySettings).Methods("POST")
	router.HandleFunc("/loadMyProfile", handler.LoadMyProfile).Methods("GET")

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


