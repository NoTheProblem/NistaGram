package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"post-service/model"
)

type PostRepository struct {
	Database *mongo.Database
}

func (repository *PostRepository) AddPost(post *model.Post) (string, error) {
	posts := repository.Database.Collection("posts")
	res, err := posts.InsertOne(context.TODO(), &post)
	if err != nil {
		return "", fmt.Errorf("post is not created")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}


func (repository *PostRepository) GetAll() []bson.D{
	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		log.Fatal(err)
	}
	return postsFiltered
}


