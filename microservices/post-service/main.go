package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"post-service/handler"
	"post-service/model"
	"post-service/repository"
	"post-service/service"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"postgres","root","post-service","5432")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	database.AutoMigrate(&model.Post{},&model.Tag{},&model.Comment{}, &model.Location{})
	return database
}

func initRepo(database *gorm.DB) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initServices(repository *repository.PostRepository) *service.PostService {
	return &service.PostService{PostRepository: repository}
}

func initHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
}
func handleFunc(handler *handler.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)


	router.HandleFunc("/uploadPost/{username}", handler.CreateNewPost).Methods("POST")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8081"), router))

}

func main() {
	database := initDB()
	repo := initRepo(database)
	service := initServices(repo)
	handler := initHandler(service)
	handleFunc(handler)
}