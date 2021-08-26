package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func (repository *PostRepository) GetHomeFeedPublic() []bson.D {
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

func (repository *PostRepository) GetHomeFeedUsername(username string) []bson.D {
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

func (repository *PostRepository) GetProfilePosts(username string) []bson.D {

	postsCollection := repository.Database.Collection("posts")
	filterCursor, err := postsCollection.Find(context.TODO(), bson.M{"owner": username})
	if err != nil {
		fmt.Println("other error")
		log.Fatal(err)
	}

	var postsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &postsFiltered); err != nil {
		fmt.Println("some error")
		log.Fatal(err)
	}
	return postsFiltered
}


func (repository *PostRepository) GetPost(id uuid.UUID) bson.D {
	postsCollection := repository.Database.Collection("posts")
	var post bson.D
	err := postsCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	return post
}

func (repository *PostRepository) AddComment(post *model.Post)  error {
	posts := repository.Database.Collection("posts")
	result, err := posts.UpdateOne(
		context.TODO(),
		bson.M{"id": post.Id},
		bson.D{
			{"$set", bson.D{{"postComments", post.PostComments}}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return  nil
}

func (repository *PostRepository) UpdateLikes(post *model.Post)  error {
	posts := repository.Database.Collection("posts")
	result, err := posts.UpdateOne(
		context.TODO(),
		bson.M{"id": post.Id},
		bson.D{
			{"$set", bson.D{{"numberOfLikes", post.NumberOfLikes}}},
			{"$set", bson.D{{"usersLiked", post.UsersLiked}}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return  nil
}

func (repository *PostRepository) UpdateDisLikes(post *model.Post)  error {
	posts := repository.Database.Collection("posts")
	result, err := posts.UpdateOne(
		context.TODO(),
		bson.M{"id": post.Id},
		bson.D{
			{"$set", bson.D{{"numberOfDislikes", post.NumberOfDislikes}}},
			{"$set", bson.D{{"usersDisliked", post.UsersDisliked}}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return  nil
}

func (repository *PostRepository) AddReport(report *model.Report) (string, error) {
	reports := repository.Database.Collection("reports")
	res, err := reports.InsertOne(context.TODO(), &report)
	if err != nil {
		return "", fmt.Errorf("report is not created")
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repository *PostRepository) GetReportById(id uuid.UUID) bson.D {
	reportsCollection := repository.Database.Collection("reports")
	var report bson.D
	err := reportsCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&report)
	if err != nil {
		log.Fatal(err)
	}
	return report
}

func (repository *PostRepository) AnswerReport(id uuid.UUID,  penalty int)  error {
	reportsCollection := repository.Database.Collection("reports")
	result, err := reportsCollection.UpdateOne(
		context.TODO(),
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"isanswered", true}}},
			{"$set", bson.D{{"penalty", penalty}}},
		},
	)
	if err != nil {
		fmt.Println("failed to update")
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return  nil
}

func (repository *PostRepository) GetUnAnsweredReports() []bson.D {

	reportsCollection := repository.Database.Collection("reports")
	filterCursor, err := reportsCollection.Find(context.TODO(), bson.M{"isanswered": false})
	fmt.Println(filterCursor)
	if err != nil {
		log.Fatal(err)
	}
	var reportsFiltered []bson.D
	if err = filterCursor.All(context.TODO(), &reportsFiltered); err != nil {
		log.Fatal(err)
	}
	return reportsFiltered
}

func (repository *PostRepository) DeletePost(id uuid.UUID) error{

	posts := repository.Database.Collection("posts")
	_, err := posts.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil
}


func (repository *PostRepository) DeleteUserPosts(username string) error{

	posts := repository.Database.Collection("posts")
	_, err := posts.DeleteMany(context.TODO(), bson.M{"owner": username})
	if err != nil {
		return err
	}
	return nil
}
