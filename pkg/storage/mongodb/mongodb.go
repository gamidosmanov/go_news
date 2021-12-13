package mongodb

import (
	"context"
	"go_news/pkg/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "local"
	collectionName = "posts"
)

type Storage struct {
	db *mongo.Client
}

func New(uri string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return &Storage{db: client}, nil
}

func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	for cur.Next(context.Background()) {
		var post storage.Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (s *Storage) AddPost(post storage.Post) error {
	maxID, err := s.maxID()
	if err != nil {
		return err
	}
	collection := s.db.Database(databaseName).Collection(collectionName)
	post.CreatedAt = time.Now().Unix()
	post.ID = maxID + 1
	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdatePost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{{"id", post.ID}}
	update := bson.D{{"$set", bson.D{
		{"title", post.Title},
		{"content", post.Content},
		{"authorid", post.AuthorID},
		{"publishedat", post.PublishedAt},
	}}}
	_, err := collection.UpdateOne(context.Background(), filter, update, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeletePost(post storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{{"id", post.ID}}
	_, err := collection.DeleteOne(context.Background(), filter, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) maxID() (int, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$name"},
			{"maxID", bson.D{
				{"$max", "$id"},
			}},
		}},
	}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)
	cursor, err := collection.Aggregate(
		context.Background(),
		mongo.Pipeline{groupStage},
		opts)
	if err != nil {
		return -1, err
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		return -1, err
	}
	res := results[0]["maxID"].(int32)
	return int(res), nil
}
