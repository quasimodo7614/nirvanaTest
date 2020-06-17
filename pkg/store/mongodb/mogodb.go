package mongodb

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/caicloud/nirvana/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/quasimodo7614/nirvanatest/pkg/store"
	"github.com/quasimodo7614/nirvanatest/pkg/types"
)

const (
	// DefaultDBName is default db name to store vmapp data and so on
	DefaultDBName         = "test"
	MessageCollectionName = "test"
)

// MS is global mongo driver instance
var MS *MongoStore

func init() {
	MS = NewMongoStore()
}

// MongoStore is a mongo client
type MongoStore struct {
	Client            *mongo.Client
	MessageCollection *mongo.Collection
}

// NewMongoStore news a mongo client
func NewMongoStore() *MongoStore {
	host := os.Getenv("MONGODB_URL")
	if host != "" && !strings.Contains(host, "mongodb://") {
		host = fmt.Sprintf("mongodb://%s", host)
	}
	if host == "" {
		host = "mongodb://localhost:27917"
	}
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("MongoDB connected: %s", host)

	db := client.Database(DefaultDBName)
	return &MongoStore{
		Client:            client,
		MessageCollection: db.Collection(MessageCollectionName),
	}
}

func (s *MongoStore) Add(ctx context.Context, obj types.Message) (error) {
	_, err := s.MessageCollection.InsertOne(ctx, obj)
	return err
}

func (s *MongoStore) Del(ctx context.Context, id int) error {
	_, err := s.MessageCollection.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func (s *MongoStore) Upd(ctx context.Context, obj types.Message) (error) {
	filter := bson.D{{"id", obj.ID}}
	r, err := s.MessageCollection.UpdateOne(ctx, filter, bson.D{{"$set", obj}})
	if r.MatchedCount < 1 {
		return errors.New("not found")
	}
	return err
}

func (s *MongoStore) Get(ctx context.Context, id int) (types.Message, error) {
	re := types.Message{}
	filter := bson.M{"id": id}
	err := s.MessageCollection.FindOne(ctx, filter).Decode(&re)
	return re, err
}

func (s *MongoStore) List(ctx context.Context, count int) ([]types.Message, error) {
	var ret []types.Message
	cur, err := s.MessageCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if count == 0 {
		count = store.DefaultCountNum
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) && count > 0 {
		var a types.Message
		err := cur.Decode(&a)
		if err != nil {
			return nil, err
		}
		ret = append(ret, a)
		count --
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
