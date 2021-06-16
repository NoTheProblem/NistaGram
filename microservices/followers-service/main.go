package main

import (
	"fmt"
	"followers-service/handler"
	"followers-service/repository"
	"followers-service/service"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"net/http"
	"os"
)

func initFollowRepository(databaseDriver neo4j.Driver) *repository.FollowRepository {
	return &repository.FollowRepository{DatabaseDriver: databaseDriver}
}

func initFollowService(repository *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepository: repository}
}

func initFollowHandler(service *service.FollowService) *handler.FollowHandler {
	return &handler.FollowHandler{FollowService: service}
}


func handleFunc(handler *handler.FollowHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/follow/{follower}/{following}", handler.FollowRequest).Methods("PUT")
	router.HandleFunc("/hello", handler.Hello).Methods("GET")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func initDatabase() *neo4j.Driver {
	var (
		driver neo4j.Driver
		err    error
	)
	for {
		driver, err = neo4j.NewDriver("bolt://"+os.Getenv("NEO4J_DBNAME")+":"+os.Getenv("NEO4J_PORT")+"/neo4j", neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASS"), "Neo4j"))

		if err != nil {
			fmt.Println("Cannot connect to database!")
		} else {
			fmt.Println(" Successfully connected to the database!")
			break
		}
	}
	return &driver
}

func main() {
	database := initDatabase()

	authenticationRepository := initFollowRepository(*database)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}