package main

import (
	"fmt"
	"followers-service/handler"
	"followers-service/repository"
	"followers-service/service"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/rs/cors"
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

	router.HandleFunc("/addUser", handler.AddUser).Methods("POST")
	router.HandleFunc("/updateUser", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/deleteUser/{username}", handler.DeleteUser).Methods("DELETE")
	router.HandleFunc("/follow/{username}", handler.Follow).Methods("POST")
	router.HandleFunc("/unfollow/{username}", handler.RemoveFollower).Methods("PUT")
	router.HandleFunc("/isFollowing/{username}", handler.IsFollowing).Methods("GET")
	router.HandleFunc("/block/{user}", handler.Block).Methods("POST")
	router.HandleFunc("/unblock/{user}", handler.Unblock).Methods("PUT")
	router.HandleFunc("/acceptRequest/{follower}", handler.AcceptRequest).Methods("PUT")
	router.HandleFunc("/following", handler.FindAllFollowing).Methods("GET")
	router.HandleFunc("/followers", handler.FindAllFollowers).Methods("GET")
	router.HandleFunc("/turnOnNotification/{username}", handler.TurnNotificationsForUserOn).Methods("PUT")
	router.HandleFunc("/turnOffNotification/{username}", handler.TurnNotificationsForUserOff).Methods("PUT")
	router.HandleFunc("/followersWithNotification", handler.FindAllFollowersWithNotificationTurnOn).Methods("GET")
	router.HandleFunc("/recommendedProfiles", handler.GetRecommendedProfiles).Methods("GET")

	c := SetupCors()

	http.Handle("/", c.Handler(router))
	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), c.Handler(router))
	if err != nil {
		log.Println(err)
	}
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

func SetupCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins, for now
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	})
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
