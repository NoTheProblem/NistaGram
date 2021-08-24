package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	cors "github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"post-service/handler"
	"post-service/repository"
	"post-service/service"
)

func initDB() *mongo.Database {

	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database("post-database")
	return database
}

func initPostRepo(database *mongo.Database) *repository.PostRepository {
	return &repository.PostRepository{Database: database}
}

func initServices(postRepo *repository.PostRepository) *service.PostService {
	return &service.PostService{PostRepository: postRepo}
}

func initHandler(service *service.PostService) *handler.PostHandler {
	return &handler.PostHandler{PostService: service}
}
func handleFunc(handler *handler.PostHandler) {
	router := mux.NewRouter().StrictSlash(true)


	router.HandleFunc("/uploadPost", handler.CreateNewPost).Methods("POST")
	router.HandleFunc("/homeFeed", handler.GetHomeFeed).Methods("GET")
	router.HandleFunc("/{username}", handler.GetPostsByUsername).Methods("GET")
	router.HandleFunc("/delete", handler.Delete).Methods("POST")

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
	postRepo := initPostRepo(database)
	service := initServices(postRepo)
	handler := initHandler(service)

	handleFunc(handler)
}
