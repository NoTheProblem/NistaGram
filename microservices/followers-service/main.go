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

func initFollowRepository(databaseSession *neo4j.Session) *repository.FollowRepository {
	return &repository.FollowRepository{DatabaseSession: databaseSession}
}

func initFollowService(repository *repository.FollowRepository) *service.FollowService {
	return &service.FollowService{FollowRepository: repository}
}

func initFollowHandler(service *service.FollowService) *handler.FollowHandler {
	return &handler.FollowHandler{FollowService: service}
}


func handleFunc(handler *handler.FollowHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/follow", handler.Follow).Methods("POST")
	router.HandleFunc("/unfollow/{following}", handler.RemoveFollower).Methods("PUT")
	router.HandleFunc("/block/{following}", handler.Block).Methods("POST")
	router.HandleFunc("/unblock/{following}", handler.Unblock).Methods("PUT")
	router.HandleFunc("/acceptRequest/{follower}", handler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/following", handler.FindAllFollowing).Methods("GET")
	router.HandleFunc("/followers", handler.FindAllFollowers).Methods("GET")
	router.HandleFunc("/turnNotification/{username}", handler.TurnNotificationsForUserOn).Methods("PUT")
	router.HandleFunc("/turnOffNotification/{username}", handler.TurnNotificationsForUserOff).Methods("PUT")
	router.HandleFunc("/followersNot", handler.FindAllFollowersWithNotificationTurnOn).Methods("GET")


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}

func initDatabase() (neo4j.Session, error) {
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	if driver, err = neo4j.NewDriver("neo4j://neo4j:7687", neo4j.BasicAuth("neo4j", "12345", "")); err != nil {
		return nil, err
	}

	if session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}); err != nil {
		return nil, err
	}

	_, err = session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run("match (u) return u;", map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		if result.Next() {
			return result.Record().Values[0], err
		}
		return nil, result.Err()
	})

	if err != nil {
		return nil, err
	}
	return session, nil

}

func main() {
	session, err := initDatabase()
	if err != nil {
		fmt.Println(err)
		return
	}

	authenticationRepository := initFollowRepository(&session)
	authenticationService := initFollowService(authenticationRepository)
	authenticationHandler := initFollowHandler(authenticationService)

	handleFunc(authenticationHandler)
}