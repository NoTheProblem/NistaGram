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
	"post-service/handler"
	"post-service/model"
	"post-service/repository"
	"post-service/service"
)

func initDB() *gorm.DB {

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		"post-service","test","post-db","5432")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	database.AutoMigrate(&model.Post{})
	database.AutoMigrate(&model.Comment{})
	database.AutoMigrate(&model.Tag{})
	database.AutoMigrate(&model.Location{})
	return database
}

func initPostRepo(database *gorm.DB) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initCommentRepo(database *gorm.DB) *repository.CommentRepository {
	return &repository.CommentRepository{Database: database}
}

func initTagRepo(database *gorm.DB) *repository.TagRepository {
	return &repository.TagRepository{Database: database}
}

func initLocationRepo(database *gorm.DB) *repository.LocationRepository {
	return &repository.LocationRepository{Database: database}
}

func initServices(locationRepo *repository.LocationRepository, commentRepo *repository.CommentRepository,
	postRepo *repository.PostRepository, tagRepo *repository.TagRepository) *service.PostService {
	return &service.PostService{LocationRepository: locationRepo, CommentRepository: commentRepo, PostRepository: postRepo,
		TagRepository: tagRepo}
}

func initHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
}
func handleFunc(handler *handler.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)


	router.HandleFunc("/uploadPost/{username}", handler.CreateNewPost).Methods("POST")

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
	locationRepo := initLocationRepo(database)
	commentRepo := initCommentRepo(database)
	postRepo := initPostRepo(database)
	tagRepo := initTagRepo(database)
	service := initServices(locationRepo,commentRepo,postRepo,tagRepo)
	handler := initHandler(service)

	handleFunc(handler)
}
