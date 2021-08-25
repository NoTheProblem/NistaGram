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
	"verification-service/handler"
	"verification-service/model"
	"verification-service/repository"
	"verification-service/service"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"verification-service","test","verification-db","5432")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.VerificationRequest{})
	return database
}

func initRepo(database *gorm.DB) *repository.VerificationRepository {
	return &repository.VerificationRepository{Database: database}
}
func initServices(postRepo *repository.VerificationRepository) *service.VerificationService {
	return &service.VerificationService{VerificationRepository: postRepo}
}

func initHandler(service *service.VerificationService) *handler.VerificationHandler {
	return &handler.VerificationHandler{VerificationService: service}
}
func handleFunc(handler *handler.VerificationHandler) {
	router := mux.NewRouter().StrictSlash(true)


	router.HandleFunc("/user", handler.CreateNewUserRequest).Methods("POST")
	router.HandleFunc("/answer", handler.AnswerRequest).Methods("PUT")
	router.HandleFunc("/getUnAnswered", handler.GetAllUnAnswered).Methods("GET")

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
	verificationRepo := initRepo(database)
	service := initServices(verificationRepo)
	handler := initHandler(service)

	handleFunc(handler)
}
