package repository

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type FollowRepository struct {
	DatabaseDriver neo4j.Driver
}

func (u *FollowRepository) RegisterUser(usernameER string, usernameING string) (err error) {
	session := u.DatabaseDriver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	defer func() {
		err = session.Close()
	}()
	if _, err := session.
		WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
			return u.persistUser(tx, usernameER, usernameING)
		}); err != nil {
		return err
	}
	return nil
}

func (u *FollowRepository) persistUser(tx neo4j.Transaction, usernameER string, usernameING string) (interface{}, error) {
	query := "CREATE (:Relationship {UsernameFollower: $UsernameFollower, UsernameFollowing: $UsernameFollowing})"

	parameters := map[string]interface{}{
		"UsernameFollower":   usernameER,
		"UsernameFollowing":  usernameING,
	}
	_, err := tx.Run(query, parameters)
	return nil, err
}


